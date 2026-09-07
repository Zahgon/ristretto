package main

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

type S struct {
	key  uint64
	val  []byte
	next *S
	inGo bool
}

var (
	ssz      = int(unsafe.Sizeof(S{}))
	lo, hi   = int64(1 << 30), int64(16 << 30)
	increase = true
	stop     int32
	fill     []byte
	maxMB    = 32

	cycles int64 = 16
)
var numbytes int64
var counter int64

func newS(sz int) *S { _ = "STUB: not implemented"; return nil }

func freeS(s *S) { _ = "STUB: not implemented"; return }

func (s *S) allocateNext(sz int) { _ = "STUB: not implemented"; return }

func (s *S) deallocNext() { _ = "STUB: not implemented"; return }

func memory() { _ = "STUB: not implemented"; return }

func viaLL() { _ = "STUB: not implemented"; return }

func main() {
	check()
	fill = make([]byte, maxMB<<20)
	_, _ = rand.Read(fill)

	c := make(chan os.Signal, 10)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Stopping")
		atomic.StoreInt32(&stop, 1)
	}()
	go func() {
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	viaLL()
	if left := NumAllocBytes(); left != 0 {
		log.Fatalf("Unable to deallocate all memory: %v\n", left)
	}
	runtime.GC()
	fmt.Println("Done. Reduced to zero memory usage.")
	time.Sleep(5 * time.Second)
}
