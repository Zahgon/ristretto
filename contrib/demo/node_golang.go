//go:build !jemalloc
// +build !jemalloc

package main

func newNode(val int) *node { _ = "STUB: not implemented"; return nil }

func freeNode(n *node) { _ = "STUB: not implemented"; return }
