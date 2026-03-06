package processor

import (
	"bufio"
	"os"
	"sync"
)

func ProcessFileConcurrentIO(path string, workers int) (Counts, error) {
	file, err := os.Open(path)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	lines := make(chan string, 1000)
	results := make(chan Counts, workers)

	var wg sync.WaitGroup

	// worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			local := Counts{}

			for line := range lines {
				switch line {
				case "INFO":
					local.Info++
				case "WARN":
					local.Warn++
				case "ERROR":
					local.Error++
				}
			}

			results <- local
		}()
	}

	// reader
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	wg.Wait()
	close(results)

	final := Counts{}

	for r := range results {
		final.Info += r.Info
		final.Warn += r.Warn
		final.Error += r.Error
	}

	return final, nil
}
