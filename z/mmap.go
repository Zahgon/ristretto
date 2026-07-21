package z

import (
	"os"
)

func Mmap(fd *os.File, writable bool, size int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Munmap(b []byte) error { _ = "STUB: not implemented"; return nil }

func Madvise(b []byte, readahead bool) error { _ = "STUB: not implemented"; return nil }

func Msync(b []byte) error { _ = "STUB: not implemented"; return nil }
