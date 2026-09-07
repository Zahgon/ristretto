//go:build jemalloc
// +build jemalloc

package main

import (
	"log"
)

func Calloc(size int) []byte { _ = "STUB: not implemented"; return nil }
func Free(bs []byte)         { _ = "STUB: not implemented"; return }
func NumAllocBytes() int64   { _ = "STUB: not implemented"; return 0 }

func check() { _ = "STUB: not implemented"; return }

func init() {
	log.Println("USING JEMALLOC")
}
