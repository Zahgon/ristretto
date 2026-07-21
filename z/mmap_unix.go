//go:build !windows && !darwin && !plan9 && !linux && !wasip1 && !js
// +build !windows,!darwin,!plan9,!linux,!wasip1,!js

package z

import (
	"os"
)

func mmap(fd *os.File, writable bool, size int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func munmap(b []byte) error { _ = "STUB: not implemented"; return nil }

func madvise(b []byte, readahead bool) error { _ = "STUB: not implemented"; return nil }

func msync(b []byte) error { _ = "STUB: not implemented"; return nil }
