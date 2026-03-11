//arena.go

package memory

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
