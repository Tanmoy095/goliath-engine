package runtime

import "sync/atomic"

// task.go
type TaskHeader struct {
	ID           uint64   // 8 bytes
	Timestamp    int64    // 8 bytes
	PayloadPtr   uintptr  // 8 bytes (Pointer to Phase 1 Arena memory)
	DependencyID uint64   // 8 bytes
	State        uint32   // 4 bytes  <-- THIS CONNECTS TO state.go
	Type         uint16   // 2 bytes
	Flags        uint8    // 1 byte   <-- THIS CONNECTS TO state.go
	Priority     uint8    // 1 byte
	_padding     [24]byte // 32 bytes (Padding to make the struct exactly 64 bytes) to avoid false sharing and ensure cache line alignment
}

//TaskHeader, it takes up exactly 64 bytes of physical RAM. The State field lives exactly 32 bytes from the start of the struct.
// meticulously mapped the geography of this object so the CPU can read it blindly.

func (t *TaskHeader) Claim() bool {
	return atomic.CompareAndSwapUint32(&t.State, StatePending, StateRunning)
}
func (t *TaskHeader) MarkDone() {
	atomic.StoreUint32(&t.State, StateCompleted)
}
func (t *TaskHeader) MarkFailed() {
	atomic.StoreUint32(&t.State, StateFailed)
}
