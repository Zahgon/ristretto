package ristretto

import (
	"sync"
	"time"
)

type updateFn[V any] func(cur, prev V) bool

type storeItem[V any] struct {
	key        uint64
	conflict   uint64
	value      V
	expiration time.Time
}

type store[V any] interface {
	Get(uint64, uint64) (V, bool)

	Expiration(uint64) time.Time

	Set(*Item[V])

	Del(uint64, uint64) (uint64, V)

	Update(*Item[V]) (V, bool)

	Cleanup(policy *defaultPolicy[V], onEvict func(item *Item[V]))

	Clear(onEvict func(item *Item[V]))
	SetShouldUpdateFn(f updateFn[V])

	IterValues(cb func(v V) (stop bool))
}

func newStore[V any]() store[V] { _ = "STUB: not implemented"; return nil }

const numShards uint64 = 256

type shardedMap[V any] struct {
	shards    []*lockedMap[V]
	expiryMap *expirationMap[V]
}

func newShardedMap[V any]() *shardedMap[V] { _ = "STUB: not implemented"; return nil }

func (m *shardedMap[V]) SetShouldUpdateFn(f updateFn[V]) { _ = "STUB: not implemented"; return }

func (sm *shardedMap[V]) IterValues(cb func(v V) (stop bool)) { _ = "STUB: not implemented"; return }

func (sm *shardedMap[V]) Get(key, conflict uint64) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (sm *shardedMap[V]) Expiration(key uint64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (sm *shardedMap[V]) Set(i *Item[V]) { _ = "STUB: not implemented"; return }

func (sm *shardedMap[V]) Del(key, conflict uint64) (uint64, V) {
	_ = "STUB: not implemented"
	return 0, *new(V)
}

func (sm *shardedMap[V]) Update(newItem *Item[V]) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (sm *shardedMap[V]) Cleanup(policy *defaultPolicy[V], onEvict func(item *Item[V])) {
	_ = "STUB: not implemented"
	return
}

func (sm *shardedMap[V]) Clear(onEvict func(item *Item[V])) { _ = "STUB: not implemented"; return }

type lockedMap[V any] struct {
	sync.RWMutex
	data         map[uint64]storeItem[V]
	em           *expirationMap[V]
	shouldUpdate updateFn[V]
}

func newLockedMap[V any](em *expirationMap[V]) *lockedMap[V] { _ = "STUB: not implemented"; return nil }

func (m *lockedMap[V]) setShouldUpdateFn(f updateFn[V]) { _ = "STUB: not implemented"; return }

func (m *lockedMap[V]) get(key, conflict uint64) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *lockedMap[V]) Expiration(key uint64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (m *lockedMap[V]) Set(i *Item[V]) { _ = "STUB: not implemented"; return }

func (m *lockedMap[V]) Del(key, conflict uint64) (uint64, V) {
	_ = "STUB: not implemented"
	return 0, *new(V)
}

func (m *lockedMap[V]) Update(newItem *Item[V]) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *lockedMap[V]) Clear(onEvict func(item *Item[V])) { _ = "STUB: not implemented"; return }
