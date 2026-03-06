package processor

import (
	"bufio"
	"os"
)

type Counts struct {
	Info  int
	Warn  int
	Error int
}

func ProcessFileIO(path string) (Counts, error) {
	file, err := os.Open(path)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counts := Counts{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "INFO" {
			counts.Info++
		} else if line == "WARN" {
			counts.Warn++
		} else if line == "ERROR" {
			counts.Error++
		}
	}

	return counts, scanner.Err()
}
