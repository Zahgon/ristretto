package z

import (
	"os"
)

const (
	defaultCapacity = 64
	defaultTag      = "buffer"
)

type Buffer struct {
	padding       uint64
	offset        uint64
	buf           []byte
	bufType       BufferType
	curSz         int
	maxSz         int
	mmapFile      *MmapFile
	autoMmapAfter int
	autoMmapDir   string
	persistent    bool
	tag           string
}

func NewBuffer(capacity int, tag string) *Buffer { _ = "STUB: not implemented"; return nil }

func NewBufferPersistent(path string, capacity int) (*Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBufferTmp(dir string, capacity int) (*Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newBufferFile(file *os.File, capacity int) (*Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBufferSlice(slice []byte) *Buffer { _ = "STUB: not implemented"; return nil }

func (b *Buffer) WithAutoMmap(threshold int, path string) *Buffer {
	_ = "STUB: not implemented"
	return nil
}

func (b *Buffer) WithMaxSize(size int) *Buffer { _ = "STUB: not implemented"; return nil }

func (b *Buffer) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (b *Buffer) LenWithPadding() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) LenNoPadding() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Grow(n int) { _ = "STUB: not implemented"; return }

func (b *Buffer) Allocate(n int) []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) AllocateOffset(n int) int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) writeLen(sz int) { _ = "STUB: not implemented"; return }

func (b *Buffer) SliceAllocate(sz int) []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) StartOffset() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) WriteSlice(slice []byte) { _ = "STUB: not implemented"; return }

func (b *Buffer) SliceIterate(f func(slice []byte) error) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	UseCalloc BufferType = iota
	UseMmap
	UseInvalid
)

type BufferType int

func (t BufferType) String() string { _ = "STUB: not implemented"; return "" }

type LessFunc func(a, b []byte) bool
type sortHelper struct {
	offsets []int
	b       *Buffer
	tmp     *Buffer
	less    LessFunc
	small   []int
}

func (s *sortHelper) sortSmall(start, end int) { _ = "STUB: not implemented"; return }

func assert(b bool) { _ = "STUB: not implemented"; return }

func check(err error) { _ = "STUB: not implemented"; return }

func check2(_ interface{}, err error) { _ = "STUB: not implemented"; return }

func (s *sortHelper) merge(left, right []byte, start, end int) { _ = "STUB: not implemented"; return }

func (s *sortHelper) sort(lo, hi int) []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) SortSlice(less func(left, right []byte) bool) { _ = "STUB: not implemented"; return }

func (b *Buffer) SortSliceBetween(start, end int, less LessFunc) { _ = "STUB: not implemented"; return }

func rawSlice(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Slice(offset int) ([]byte, int) { _ = "STUB: not implemented"; return nil, 0 }

func (b *Buffer) SliceOffsets() []int { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Data(offset int) []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (b *Buffer) Release() error { _ = "STUB: not implemented"; return nil }
