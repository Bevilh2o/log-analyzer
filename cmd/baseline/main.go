package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"crypto/sha256"
)

func main() {
	start := time.Now()

	file, err := os.Open("testdata/app.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var infoCount, warnCount, errorCount int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < 50; i++ {
			hash := sha256.Sum256([]byte(line))
			_ = hash
		}

		if strings.Contains(line, "INFO") {
			infoCount++
		} else if strings.Contains(line, "WARN") {
			warnCount++
		} else if strings.Contains(line, "ERROR") {
			errorCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	elapsed := time.Since(start)

	fmt.Println("Baseline Results")
	fmt.Println("----------------")
	fmt.Println("INFO:", infoCount)
	fmt.Println("WARN:", warnCount)
	fmt.Println("ERROR:", errorCount)
	fmt.Println("Time:", elapsed)
}
