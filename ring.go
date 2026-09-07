package ristretto

import (
	"sync"
)

type ringConsumer interface {
	Push([]uint64) bool
}

type ringStripe struct {
	cons ringConsumer
	data []uint64
	capa int
}

func newRingStripe(cons ringConsumer, capa int64) *ringStripe {
	_ = "STUB: not implemented"
	return nil
}

func (s *ringStripe) Push(item uint64) { _ = "STUB: not implemented"; return }

type ringBuffer struct {
	pool *sync.Pool
}

func newRingBuffer(cons ringConsumer, capa int64) *ringBuffer {
	_ = "STUB: not implemented"
	return nil
}

func (b *ringBuffer) Push(item uint64) { _ = "STUB: not implemented"; return }
