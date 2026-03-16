package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/Tanmoy095/goliath-engine/pkg/memory"
)

//entry point for our engine.

func main() {
	//Diagnostic backdoor. it allows us to monitor the performance of our allocator in real time. By starting an HTTP server on port 6000, we can expose metrics or diagnostic information about our allocator's performance. This is useful for benchmarking and understanding how our custom memory management system behaves under load, without needing to stop the application or use external profiling tools. It provides a convenient way to access performance data while the application is running.
	go func() {
		fmt.Println("Starting goliath-engine on localhost:6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Printf("failed to start goliath-engine: %v\n", err)
		}
	}()
	fmt.Println("Initializing 50MB Slab Allocator...")
	allocator := memory.NewSlabAllocator(1024 * 1024 * 50) //50MB arena

	iterations := 50000000                                               //50 million
	fmt.Printf("Simulating %d rapid allocations/frees...\n", iterations) //
	start := time.Now()
	for i := 0; i < iterations; i++ {
		ptr := allocator.Allocate(64)
		allocator.Free(ptr, 64)

	}
	elapsed := time.Since(start) //elapsed is a duration type that represents the time taken for the loop to execute
	println(elapsed)             //print the elapsed time

	//Operations Per Second (Throughput) . It's a measure of how many operations (allocations and frees) our allocator can perform in one second. We calculate it by dividing the total number of iterations by the elapsed time in seconds. This gives us a sense of the efficiency and performance of our custom slab allocator compared to the standard Go heap allocation.
	opsPerSec := float64(iterations) / elapsed.Seconds()
	fmt.Printf("Done.\n")
	fmt.Printf("Total Time ; %v\n", elapsed)
	fmt.Printf("Throughput: %.2f ops/sec\n", opsPerSec)

}
