package memory

//free list implementation for memory management

import (
	"sync"
	"unsafe"
)

type FreeList struct {
	mu       sync.Mutex
	pointers []unsafe.Pointer
}

func NewFreeList() *FreeList {
	return &FreeList{
		pointers: make([]unsafe.Pointer, 0), // Initialize with an empty slice of pointers
	}
}

func (fl *FreeList) Push(ptr unsafe.Pointer) {
	fl.mu.Lock()
	defer fl.mu.Unlock()
	fl.pointers = append(fl.pointers, ptr) // Add the pointer to the end of the slice
}

func (f *FreeList) Pop() unsafe.Pointer {

	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.pointers) == 0 {
		return nil
	}
	index := len(f.pointers) - 1
	ptr := f.pointers[index]
	f.pointers = f.pointers[:index]
	return ptr
}
