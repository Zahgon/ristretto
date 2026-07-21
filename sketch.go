package ristretto

type cmSketch struct {
	rows [cmDepth]cmRow
	seed [cmDepth]uint64
	mask uint64
}

const (
	cmDepth = 4
)

func newCmSketch(numCounters int64) *cmSketch { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func (s *cmSketch) Increment(hashed uint64) { _ = "STUB: not implemented"; return }

func (s *cmSketch) Estimate(hashed uint64) int64 { _ = "STUB: not implemented"; return 0 }

func (s *cmSketch) Reset() { _ = "STUB: not implemented"; return }

func (s *cmSketch) Clear() { _ = "STUB: not implemented"; return }

type cmRow []byte

func newCmRow(numCounters int64) cmRow { _ = "STUB: not implemented"; return *new(cmRow) }

func (r cmRow) get(n uint64) byte { _ = "STUB: not implemented"; return 0 }

func (r cmRow) increment(n uint64) { _ = "STUB: not implemented"; return }

func (r cmRow) reset() { _ = "STUB: not implemented"; return }

func (r cmRow) clear() { _ = "STUB: not implemented"; return }

func (r cmRow) string() string { _ = "STUB: not implemented"; return "" }

func next2Power(x int64) int64 { _ = "STUB: not implemented"; return 0 }
