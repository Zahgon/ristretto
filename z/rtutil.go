package z

import (
	"unsafe"
)

//go:linkname NanoTime runtime.nanotime
func NanoTime() int64

//go:linkname CPUTicks runtime.cputicks
func CPUTicks() int64

type stringStruct struct {
	str unsafe.Pointer
	len int
}

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr

func MemHash(data []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func MemHashString(str string) uint64 { _ = "STUB: not implemented"; return 0 }

//go:linkname FastRand runtime.fastrand
func FastRand() uint32

//go:linkname memclrNoHeapPointers runtime.memclrNoHeapPointers
func memclrNoHeapPointers(p unsafe.Pointer, n uintptr)

func Memclr(b []byte) { _ = "STUB: not implemented"; return }
