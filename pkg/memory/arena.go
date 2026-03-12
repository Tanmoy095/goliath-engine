//arena.go

package memory

import "unsafe"

type Arena struct {
	buffer []byte  //buffer is like giant arena, we can allocate memory from it
	offset uintptr //where the next memory block begins
	//Storing size duplicates state → dangerous.
}

func NewArena(size uintptr) *Arena {
	return &Arena{
		buffer: make([]byte, size),
		offset: 0,
	}
}

func (a *Arena) Alloc(size uintptr, alignment uintptr) unsafe.Pointer {
	basePtr := unsafe.Pointer(&a.buffer[0]) //get address of first byte in arena
	start := uintptr(basePtr) + a.offset
	aligned := Align(start, alignment) //must move pointer to next multiple of 64. if start 1000 it should be next multiple of 64 like near 1000: 1024
	newOffset := (aligned - uintptr(basePtr)) + size
	if newOffset > uintptr(len(a.buffer)) {
		panic("Arena out of memory ")
	}
	a.offset = newOffset
	return unsafe.Pointer(aligned)
}
