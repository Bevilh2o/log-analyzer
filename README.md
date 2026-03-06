# Log Analyzer -- Concurrency Benchmark in Go

## Overview

This project explores the performance characteristics of sequential
versus concurrent log processing in Go.

The goal is to understand how workload type (CPU-bound vs IO-bound)
affects the benefits of concurrency and how performance scales when
using worker pools.

The project simulates a backend workload by processing large log files
and analyzing how concurrency impacts throughput and execution time.

------------------------------------------------------------------------

## Objectives

-   Compare sequential and concurrent log processing
-   Measure execution time and performance scaling
-   Understand CPU-bound vs IO-bound workload behavior
-   Analyze concurrency overhead and bottlenecks
-   Explore Go profiling tools for performance analysis

------------------------------------------------------------------------

## Project Structure

    log-analyzer
    │
    ├── cmd
    │   ├── baseline        CPU-bound sequential version
    │   ├── concurrent      CPU-bound concurrent worker pool
    │   ├── io-baseline     IO-bound sequential version
    │   └── io-concurrent   IO-bound concurrent version
    │
    ├── internal
    │   └── processor
    │       ├── processor.go
    │       ├── concurrent.go
    │       └── io_processor.go
    │
    ├── testdata
    │   └── logs.txt
    │
    └── docs
        └── scaling.png

------------------------------------------------------------------------

## Implementation Details

### Baseline Version

Processes the log file sequentially, line by line.

### Concurrent Version

Uses a worker pool with goroutines and channels to process log lines in
parallel.\
A mutex ensures thread-safe aggregation of results.

To simulate a CPU-bound workload, each log line is hashed multiple times
using SHA-256.

------------------------------------------------------------------------

## Benchmark Results

### CPU-bound workload

Sequential execution time:

\~5.26 seconds

Concurrent execution time (8 workers):

\~1.85 seconds

This demonstrates how concurrency improves performance when the workload
is CPU-bound.

------------------------------------------------------------------------

## IO-bound Benchmark

Sequential IO version:

\~135µs

Concurrent IO version:

\~184µs

Unlike CPU-bound workloads, concurrency does not necessarily improve
performance when the bottleneck is disk I/O.\
In this case, scheduling and synchronization overhead slightly reduce
performance.

------------------------------------------------------------------------

## Scaling Analysis

Performance improves significantly as the number of workers increases up
to the number of available CPU cores.

Beyond that point, gains diminish due to synchronization overhead and
scheduling costs.

See the scaling graph below:

![Scaling Graph](docs/scaling.png)

------------------------------------------------------------------------

## Key Learnings

-   Concurrency does not always improve performance
-   Understanding system bottlenecks is critical before parallelizing
-   CPU-bound workloads benefit significantly from worker pools
-   IO-bound workloads may not benefit from concurrency
-   Benchmarking and profiling are essential for performance
    optimization

------------------------------------------------------------------------

## How to Run

Run the CPU-bound sequential version:

    go run cmd/baseline/main.go

Run the CPU-bound concurrent version:

    go run cmd/concurrent/main.go

Run the IO-bound sequential version:

    go run cmd/io-baseline/main.go

Run the IO-bound concurrent version:

    go run cmd/io-concurrent/main.go

------------------------------------------------------------------------

## Profiling

CPU profiling can be performed using Go's built-in profiler.

Example:

    go tool pprof cpu.prof

This allows inspection of hotspots and performance bottlenecks.

------------------------------------------------------------------------

## Future Improvements

-   Pipeline architecture with multiple stages
-   Dynamic worker configuration
-   Throughput measurement (lines/sec)
-   Memory profiling

------------------------------------------------------------------------

## Author

Michel Bevilacqua\
Backend / Systems Engineering
