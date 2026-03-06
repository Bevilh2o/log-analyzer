package main

import (
	"fmt"
	"log"
	"time"

	"log-analyzer/internal/processor"
)

func main() {
	start := time.Now()

	counts, err := processor.ProcessFileIO("testdata/logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	fmt.Println("IO Baseline Results")
	fmt.Println("-------------------")
	fmt.Println("INFO:", counts.Info)
	fmt.Println("WARN:", counts.Warn)
	fmt.Println("ERROR:", counts.Error)
	fmt.Println("Time:", elapsed)
}
