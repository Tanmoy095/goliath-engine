package memory

import (
	"testing"
)

// A dummy 64-byte struct to simulate our Phase 2 TaskHeader
type TaskHeader struct {
	ID        uint64   // 8 bytes
	State     uint64   // 8 bytes
	Timestamp uint64   // 8 bytes
	Padding   [40]byte // 40 bytes (Total: 64 bytes)
}

// Benchmark the standard Go Garbage Collected Heap
func BenchmarkStandardAllocation(b *testing.B) {
	// b.ReportAllocs() tells Go to track how many times it goes to the heap
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// This forces Go to allocate on the heap and triggers the GC
		task := new(TaskHeader)
		task.ID = uint64(i)
	}
}

// Benchmark our Custom Slab Allocator
func BenchmarkArenaAllocation(b *testing.B) {
	b.ReportAllocs()
	// Initialize a 10MB Arena for the test
	allocator := NewSlabAllocator(10 * 1024 * 1024)

	b.ResetTimer() // Don't count the setup time

	for i := 0; i < b.N; i++ {
		// 1. Allocate O(1)
		ptr := allocator.Allocate(64)

		// 2. Cast the raw memory to our Go struct (Zero-Allocation)
		task := (*TaskHeader)(ptr)
		task.ID = uint64(i)

		// 3. Free O(1)
		allocator.Free(ptr, 64)
	}
}
