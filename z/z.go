package z

import (
	"context"
	"sync"
)

type Key interface {
	~uint64 | ~string | ~[]byte | ~byte | ~int | ~uint | ~int32 | ~uint32 | ~int64
}

func KeyToHash[K Key](key K) (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

var (
	dummyCloserChan <-chan struct{}
	tmpDir          string
)

type Closer struct {
	waiting sync.WaitGroup

	ctx    context.Context
	cancel context.CancelFunc
}

func SetTmpDir(dir string) { _ = "STUB: not implemented"; return }

func NewCloser(initial int) *Closer { _ = "STUB: not implemented"; return nil }

func (lc *Closer) AddRunning(delta int) { _ = "STUB: not implemented"; return }

func (lc *Closer) Ctx() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (lc *Closer) Signal() { _ = "STUB: not implemented"; return }

func (lc *Closer) HasBeenClosed() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (lc *Closer) Done() { _ = "STUB: not implemented"; return }

func (lc *Closer) Wait() { _ = "STUB: not implemented"; return }

func (lc *Closer) SignalAndWait() { _ = "STUB: not implemented"; return }

func ZeroOut(dst []byte, start, end int) { _ = "STUB: not implemented"; return }
