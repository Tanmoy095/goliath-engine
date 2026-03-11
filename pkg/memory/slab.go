// slab.go
package memory

func Align(ptr uintptr, alignment uintptr) uintptr {
	//To avoid false sharing between worker goroutines, I ensured every allocation starts at a 64-byte boundary.
	return (ptr + alignment - 1) & ^(alignment - 1) // The expression (ptr + alignment - 1) ensures that if ptr is already aligned, it remains unchanged. If ptr is not aligned, it rounds up to the next multiple of alignment.

}
