package memory

import "unsafe"

//IN allocator.go slabAllocator  will combine arena and freelist into a working system

type SlabAllocator struct {
	arena    *Arena
	class64  *FreeList //holds exactly 64 bytes of chunks. has its own freeList
	class128 *FreeList //holds exactly 128 bytes of chunks. has its own freeList

}

func NewSlabAllocator(arenaSize int) *SlabAllocator {
	return &SlabAllocator{
		arena:    NewArena(arenaSize),
		class64:  NewFreeList(),
		class128: NewFreeList(),
	}
}

func (s *SlabAllocator) Allocate(size uintptr) unsafe.Pointer {
	if size <= 64 {
		ptr := s.class64.Pop()
		if ptr != nil {
			return ptr
		}
		return s.arena.Alloc(64, 64)

	}
	if size <= 128 {
		ptr := s.class128.Pop()
		if ptr != nil {
			return ptr
		}
		return s.arena.Alloc(128, 64)

	}
	panic("allocation size too large for slab classes")
}

func (s *SlabAllocator) Free(ptr unsafe.Pointer, size uintptr) {
	if size <= 64 {
		s.class64.Push(ptr)
		return
	}

	if size <= 128 {
		s.class128.Push(ptr)
		return
	}
	panic("Invalid Size for free")

}
