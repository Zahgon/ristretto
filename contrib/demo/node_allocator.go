//go:build jemalloc && allocator
// +build jemalloc,allocator

package main

import (
	"github.com/dgraph-io/ristretto/v2/z"
)

func init() {
	alloc = z.NewAllocator(10<<20, "demo")
}

func newNode(val int) *node { _ = "STUB: not implemented"; return nil }

func freeNode(n *node) { _ = "STUB: not implemented"; return }
