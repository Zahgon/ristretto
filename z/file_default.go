//go:build !linux
// +build !linux

package z

func (m *MmapFile) Truncate(maxSz int64) error { _ = "STUB: not implemented"; return nil }
