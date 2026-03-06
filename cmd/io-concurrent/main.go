package main

import (
	"fmt"
	"log"
	"time"

	"log-analyzer/internal/processor"
)

func main() {
	start := time.Now()

	counts, err := processor.ProcessFileConcurrentIO("testdata/logs.txt", 8)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	fmt.Println("IO Concurrent Results")
	fmt.Println("---------------------")
	fmt.Println("INFO:", counts.Info)
	fmt.Println("WARN:", counts.Warn)
	fmt.Println("ERROR:", counts.Error)
	fmt.Println("Time:", elapsed)
}
