//go:build jemalloc
// +build jemalloc

package z

/*
#cgo LDFLAGS: /usr/local/lib/libjemalloc.a -L/usr/local/lib -Wl,-rpath,/usr/local/lib -ljemalloc -lm -lstdc++ -pthread -ldl
#include <stdlib.h>
#include <jemalloc/jemalloc.h>
*/
import "C"
import (
	"sync"
	"unsafe"
)

//go:linkname throw runtime.throw
func throw(s string)

type dalloc struct {
	t  string
	sz int
}

var dallocsMu sync.Mutex
var dallocs map[unsafe.Pointer]*dalloc

func init() {

	dallocs = make(map[unsafe.Pointer]*dalloc)
}

func Calloc(n int, tag string) []byte { _ = "STUB: not implemented"; return nil }

func CallocNoRef(n int, tag string) []byte { _ = "STUB: not implemented"; return nil }

func Free(b []byte) { _ = "STUB: not implemented"; return }

func Leaks() string { _ = "STUB: not implemented"; return "" }

func ReadMemStats(stats *MemStats) { _ = "STUB: not implemented"; return }

func fetchStat(s string) uint64 { _ = "STUB: not implemented"; return 0 }

func StatsPrint() { _ = "STUB: not implemented"; return }
