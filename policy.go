package ristretto

import (
	"sync"

	"github.com/dgraph-io/ristretto/v2/z"
)

const (
	lfuSample = 5
)

func newPolicy[V any](numCounters, maxCost int64) *defaultPolicy[V] {
	_ = "STUB: not implemented"
	return nil
}

type defaultPolicy[V any] struct {
	sync.Mutex
	admit    *tinyLFU
	evict    *sampledLFU
	itemsCh  chan []uint64
	stop     chan struct{}
	done     chan struct{}
	isClosed bool
	metrics  *Metrics
}

func newDefaultPolicy[V any](numCounters, maxCost int64) *defaultPolicy[V] {
	_ = "STUB: not implemented"
	return nil
}

func (p *defaultPolicy[V]) CollectMetrics(metrics *Metrics) { _ = "STUB: not implemented"; return }

type policyPair struct {
	key  uint64
	cost int64
}

func (p *defaultPolicy[V]) processItems() { _ = "STUB: not implemented"; return }

func (p *defaultPolicy[V]) Push(keys []uint64) bool { _ = "STUB: not implemented"; return false }

func (p *defaultPolicy[V]) Add(key uint64, cost int64) ([]*Item[V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *defaultPolicy[V]) Has(key uint64) bool { _ = "STUB: not implemented"; return false }

func (p *defaultPolicy[V]) Del(key uint64) { _ = "STUB: not implemented"; return }

func (p *defaultPolicy[V]) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (p *defaultPolicy[V]) Update(key uint64, cost int64) { _ = "STUB: not implemented"; return }

func (p *defaultPolicy[V]) Cost(key uint64) int64 { _ = "STUB: not implemented"; return 0 }

func (p *defaultPolicy[V]) Clear() { _ = "STUB: not implemented"; return }

func (p *defaultPolicy[V]) Close() { _ = "STUB: not implemented"; return }

func (p *defaultPolicy[V]) MaxCost() int64 { _ = "STUB: not implemented"; return 0 }

func (p *defaultPolicy[V]) UpdateMaxCost(maxCost int64) { _ = "STUB: not implemented"; return }

type sampledLFU struct {
	maxCost  int64
	used     int64
	metrics  *Metrics
	keyCosts map[uint64]int64
}

func newSampledLFU(maxCost int64) *sampledLFU { _ = "STUB: not implemented"; return nil }

func (p *sampledLFU) getMaxCost() int64 { _ = "STUB: not implemented"; return 0 }

func (p *sampledLFU) updateMaxCost(maxCost int64) { _ = "STUB: not implemented"; return }

func (p *sampledLFU) roomLeft(cost int64) int64 { _ = "STUB: not implemented"; return 0 }

func (p *sampledLFU) fillSample(in []*policyPair) []*policyPair {
	_ = "STUB: not implemented"
	return nil
}

func (p *sampledLFU) del(key uint64) { _ = "STUB: not implemented"; return }

func (p *sampledLFU) add(key uint64, cost int64) { _ = "STUB: not implemented"; return }

func (p *sampledLFU) updateIfHas(key uint64, cost int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *sampledLFU) clear() { _ = "STUB: not implemented"; return }

type tinyLFU struct {
	freq    *cmSketch
	door    *z.Bloom
	incrs   int64
	resetAt int64
}

func newTinyLFU(numCounters int64) *tinyLFU { _ = "STUB: not implemented"; return nil }

func (p *tinyLFU) Push(keys []uint64) { _ = "STUB: not implemented"; return }

func (p *tinyLFU) Estimate(key uint64) int64 { _ = "STUB: not implemented"; return 0 }

func (p *tinyLFU) Increment(key uint64) { _ = "STUB: not implemented"; return }

func (p *tinyLFU) reset() { _ = "STUB: not implemented"; return }

func (p *tinyLFU) clear() { _ = "STUB: not implemented"; return }
