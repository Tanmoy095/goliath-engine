// state.go
// This file will hold our lock-free atomic states and bit-packed flags.
package runtime

const (
	StatePending   uint32 = 0 // uint32 because CompareAndSwap is highly optimized by Intel and AMD to work on 32-bit (4-byte) integers
	StateRunning   uint32 = 1
	StateCompleted uint32 = 2
	StateFailed    uint32 = 3
)

// The Flags
// If we used bool for IsRetry, IsCritical, and IsGPU, Go would waste 3 whole bytes (3 separate wall plates).
// By using bit-shifting (<<), we are assigning meaning to specific switches on a single 1-byte wall plate
const (
	FlagRetry    uint32 = 1 << 0 // 0001
	FlagCritical uint32 = 1 << 1 // 0010
	FlagGpu      uint32 = 1 << 2 // 0100
)
