package ristretto

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/dgraph-io/ristretto/v2/z"
)

var (
	setBufSize = 32 * 1024
)

const itemSize = int64(unsafe.Sizeof(storeItem[any]{}))

func zeroValue[T any]() T { _ = "STUB: not implemented"; return *new(T) }

type Key = z.Key

type Cache[K Key, V any] struct {
	storedItems store[V]

	cachePolicy *defaultPolicy[V]

	getBuf *ringBuffer

	setBuf chan *Item[V]

	onEvict func(*Item[V])

	onReject func(*Item[V])

	onExit (func(V))

	keyToHash func(K) (uint64, uint64)

	stop chan struct{}
	done chan struct{}

	isClosed atomic.Bool

	cost func(value V) int64

	ignoreInternalCost bool

	cleanupTicker *time.Ticker

	Metrics *Metrics
}

type Config[K Key, V any] struct {
	NumCounters int64

	MaxCost int64

	BufferItems int64

	Metrics bool

	OnEvict func(item *Item[V])

	OnReject func(item *Item[V])

	OnExit func(val V)

	ShouldUpdate func(cur, prev V) bool

	KeyToHash func(key K) (uint64, uint64)

	Cost func(value V) int64

	IgnoreInternalCost bool

	TtlTickerDurationInSec int64
}

type itemFlag byte

const (
	itemNew itemFlag = iota
	itemDelete
	itemUpdate
)

type Item[V any] struct {
	flag       itemFlag
	Key        uint64
	Conflict   uint64
	Value      V
	Cost       int64
	Expiration time.Time
	wait       chan struct{}
}

func NewCache[K Key, V any](config *Config[K, V]) (*Cache[K, V], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cache[K, V]) Wait() { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Cache[K, V]) Set(key K, value V, cost int64) bool { _ = "STUB: not implemented"; return false }

func (c *Cache[K, V]) SetWithTTL(key K, value V, cost int64, ttl time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Cache[K, V]) Del(key K) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) GetTTL(key K) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

func (c *Cache[K, V]) IterValues(cb func(v V) (stop bool)) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) Close() { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) MaxCost() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Cache[K, V]) UpdateMaxCost(maxCost int64) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) RemainingCost() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Cache[K, V]) processItems() { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) collectMetrics() { _ = "STUB: not implemented"; return }

type metricType int

const (
	hit = iota
	miss

	keyAdd
	keyUpdate
	keyEvict

	costAdd
	costEvict

	dropSets
	rejectSets

	dropGets
	keepGets

	doNotUse
)

func stringFor(t metricType) string { _ = "STUB: not implemented"; return "" }

type Metrics struct {
	all [doNotUse][]*uint64

	mu   sync.RWMutex
	life *z.HistogramData
}

func newMetrics() *Metrics { _ = "STUB: not implemented"; return nil }

func (p *Metrics) add(t metricType, hash, delta uint64) { _ = "STUB: not implemented"; return }

func (p *Metrics) get(t metricType) uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) Hits() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) Misses() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) KeysAdded() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) KeysUpdated() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) KeysEvicted() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) CostAdded() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) CostEvicted() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) SetsDropped() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) SetsRejected() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) GetsDropped() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) GetsKept() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) Ratio() float64 { _ = "STUB: not implemented"; return 0 }

func (p *Metrics) trackEviction(numSeconds int64) { _ = "STUB: not implemented"; return }

func (p *Metrics) LifeExpectancySeconds() *z.HistogramData { _ = "STUB: not implemented"; return nil }

func (p *Metrics) Clear() { _ = "STUB: not implemented"; return }

func (p *Metrics) String() string { _ = "STUB: not implemented"; return "" }
