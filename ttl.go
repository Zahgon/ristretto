package ristretto

import (
	"sync"
	"time"
)

var (
	bucketDurationSecs = int64(5)
)

func storageBucket(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func cleanupBucket(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

type bucket map[uint64]uint64

type expirationMap[V any] struct {
	sync.RWMutex
	buckets              map[int64]bucket
	lastCleanedBucketNum int64
}

func newExpirationMap[V any]() *expirationMap[V] { _ = "STUB: not implemented"; return nil }

func (m *expirationMap[_]) add(key, conflict uint64, expiration time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *expirationMap[_]) update(key, conflict uint64, oldExpTime, newExpTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *expirationMap[_]) del(key uint64, expiration time.Time) { _ = "STUB: not implemented"; return }

func (m *expirationMap[V]) cleanup(store store[V], policy *defaultPolicy[V], onEvict func(item *Item[V])) int {
	_ = "STUB: not implemented"
	return 0
}

func (m *expirationMap[V]) clear() { _ = "STUB: not implemented"; return }
