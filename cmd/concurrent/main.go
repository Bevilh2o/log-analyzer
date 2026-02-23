package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Stats struct {
	info  int
	warn  int
	error int
	mu    sync.Mutex
}

func worker(lines <-chan string, stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done()

	for line := range lines {
		for i := 0; i < 50; i++ {
			hash := sha256.Sum256([]byte(line))
			_ = hash
		}

		if strings.Contains(line, "INFO") {
			stats.mu.Lock()
			stats.info++
			stats.mu.Unlock()
		} else if strings.Contains(line, "WARN") {
			stats.mu.Lock()
			stats.warn++
			stats.mu.Unlock()
		} else if strings.Contains(line, "ERROR") {
			stats.mu.Lock()
			stats.error++
			stats.mu.Unlock()
		}
	}
}

func main() {
	start := time.Now()

	file, err := os.Open("testdata/app.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines := make(chan string, 1000)
	stats := &Stats{}
	var wg sync.WaitGroup

	numWorkers := 8

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(lines, stats, &wg)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines)
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	elapsed := time.Since(start)

	fmt.Println("Concurrent Results")
	fmt.Println("------------------")
	fmt.Println("INFO:", stats.info)
	fmt.Println("WARN:", stats.warn)
	fmt.Println("ERROR:", stats.error)
	fmt.Println("Time:", elapsed)
}
