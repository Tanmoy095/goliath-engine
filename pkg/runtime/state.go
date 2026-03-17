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
	// Bit-packed flags (Strictly uint8 to match the struct!)
	FlagRetry    uint8 = 1 << 0 // 0000 0001
	FlagCritical uint8 = 1 << 1 // 0000 0010
	FlagGpu      uint8 = 1 << 2 // 0000 0100
)
