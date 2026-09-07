//go:build windows
// +build windows

package z

import (
	"os"
)

func mmap(fd *os.File, write bool, size int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func munmap(b []byte) error { _ = "STUB: not implemented"; return nil }

func madvise(b []byte, readahead bool) error { _ = "STUB: not implemented"; return nil }

func msync(b []byte) error { _ = "STUB: not implemented"; return nil }
