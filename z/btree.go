package z

import (
	"math"
	"os"
)

var (
	pageSize = os.Getpagesize()
	maxKeys  = (pageSize / 16) - 1
	//nolint:unused
	oneThird = int(float64(maxKeys) / 3)
)

const (
	absoluteMax = uint64(math.MaxUint64 - 1)
	minSize     = 1 << 20
)

type Tree struct {
	buffer   *Buffer
	data     []byte
	nextPage uint64
	freePage uint64
	stats    TreeStats
}

func (t *Tree) initRootNode() { _ = "STUB: not implemented"; return }

func NewTree(tag string) *Tree { _ = "STUB: not implemented"; return nil }

func NewTreePersistent(path string) (*Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) reinit() { _ = "STUB: not implemented"; return }

func (t *Tree) Reset() { _ = "STUB: not implemented"; return }

func (t *Tree) Close() error { _ = "STUB: not implemented"; return nil }

type TreeStats struct {
	Allocated    int
	Bytes        int
	NumLeafKeys  int
	NumPages     int
	NumPagesFree int
	Occupancy    float64
	PageSize     int
}

func (t *Tree) Stats() TreeStats { _ = "STUB: not implemented"; return *new(TreeStats) }

func BytesToUint64Slice(b []byte) []uint64 { _ = "STUB: not implemented"; return nil }

func (t *Tree) newNode(bit uint64) node { _ = "STUB: not implemented"; return *new(node) }

func getNode(data []byte) node { _ = "STUB: not implemented"; return *new(node) }

func zeroOut(data []uint64) { _ = "STUB: not implemented"; return }

func (t *Tree) node(pid uint64) node { _ = "STUB: not implemented"; return *new(node) }

func (t *Tree) Set(k, v uint64) { _ = "STUB: not implemented"; return }

func (t *Tree) set(pid, k, v uint64) node { _ = "STUB: not implemented"; return *new(node) }

func (t *Tree) Get(k uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (t *Tree) get(n node, k uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (t *Tree) DeleteBelow(ts uint64) { _ = "STUB: not implemented"; return }

func (t *Tree) compact(n node, ts uint64) int { _ = "STUB: not implemented"; return 0 }

func (t *Tree) iterate(n node, fn func(node)) { _ = "STUB: not implemented"; return }

func (t *Tree) Iterate(fn func(node)) { _ = "STUB: not implemented"; return }

func (t *Tree) IterateKV(f func(key, val uint64) (newVal uint64)) {
	_ = "STUB: not implemented"
	return
}

func (t *Tree) print(n node, parentID uint64) { _ = "STUB: not implemented"; return }

func (t *Tree) Print() { _ = "STUB: not implemented"; return }

func (t *Tree) split(pid uint64) node { _ = "STUB: not implemented"; return *new(node) }

//nolint:unused
func (t *Tree) shareWithSiblingXXX(n node, idx int) bool { _ = "STUB: not implemented"; return false }

type node []uint64

func (n node) uint64(start int) uint64 { _ = "STUB: not implemented"; return 0 }

func keyOffset(i int) int          { _ = "STUB: not implemented"; return 0 }
func valOffset(i int) int          { _ = "STUB: not implemented"; return 0 }
func (n node) numKeys() int        { _ = "STUB: not implemented"; return 0 }
func (n node) pageID() uint64      { _ = "STUB: not implemented"; return 0 }
func (n node) key(i int) uint64    { _ = "STUB: not implemented"; return 0 }
func (n node) val(i int) uint64    { _ = "STUB: not implemented"; return 0 }
func (n node) data(i int) []uint64 { _ = "STUB: not implemented"; return nil }

func (n node) setAt(start int, k uint64) { _ = "STUB: not implemented"; return }

func (n node) setNumKeys(num int) { _ = "STUB: not implemented"; return }

func (n node) moveRight(lo int) { _ = "STUB: not implemented"; return }

const (
	bitLeaf = uint64(1 << 63)
)

func (n node) setBit(b uint64) { _ = "STUB: not implemented"; return }

func (n node) bits() uint64 { _ = "STUB: not implemented"; return 0 }

func (n node) isLeaf() bool { _ = "STUB: not implemented"; return false }

func (n node) isFull() bool { _ = "STUB: not implemented"; return false }

func (n node) search(k uint64) int { _ = "STUB: not implemented"; return 0 }

func (n node) maxKey() uint64 { _ = "STUB: not implemented"; return 0 }

func (n node) compact(lo uint64) int { _ = "STUB: not implemented"; return 0 }

func (n node) get(k uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (n node) set(k, v uint64) (numAdded int) { _ = "STUB: not implemented"; return 0 }

func (n node) iterate(fn func(node, int)) { _ = "STUB: not implemented"; return }

func (n node) print(parentID uint64) { _ = "STUB: not implemented"; return }
