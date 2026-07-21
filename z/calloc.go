package z

var numBytes int64

func NumAllocBytes() int64 { _ = "STUB: not implemented"; return 0 }

type MemStats struct {
	Allocated uint64

	Active uint64

	Resident uint64

	Retained uint64
}
