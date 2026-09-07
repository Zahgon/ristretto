//go:build !jemalloc || !cgo
// +build !jemalloc !cgo

package z

func Calloc(n int, tag string) []byte { _ = "STUB: not implemented"; return nil }

func CallocNoRef(n int, tag string) []byte { _ = "STUB: not implemented"; return nil }

func Free(b []byte) { _ = "STUB: not implemented"; return }

func Leaks() string { _ = "STUB: not implemented"; return "" }
func StatsPrint()   { _ = "STUB: not implemented"; return }

func ReadMemStats(_ *MemStats) { _ = "STUB: not implemented"; return }
