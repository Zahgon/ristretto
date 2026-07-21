package z

var mask = []uint8{1, 2, 4, 8, 16, 32, 64, 128}

func getSize(ui64 uint64) (size uint64, exponent uint64) { _ = "STUB: not implemented"; return 0, 0 }

func calcSizeByWrongPositives(numEntries, wrongs float64) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func NewBloomFilter(params ...float64) (bloomfilter *Bloom) { _ = "STUB: not implemented"; return nil }

type Bloom struct {
	bitset  []uint64
	ElemNum uint64
	sizeExp uint64
	size    uint64
	setLocs uint64
	shift   uint64
}

func (bl *Bloom) Add(hash uint64) { _ = "STUB: not implemented"; return }

func (bl Bloom) Has(hash uint64) bool { _ = "STUB: not implemented"; return false }

func (bl *Bloom) AddIfNotHas(hash uint64) bool { _ = "STUB: not implemented"; return false }

func (bl *Bloom) TotalSize() int { _ = "STUB: not implemented"; return 0 }

func (bl *Bloom) Size(sz uint64) { _ = "STUB: not implemented"; return }

func (bl *Bloom) Clear() { _ = "STUB: not implemented"; return }

func (bl *Bloom) Set(idx uint64) { _ = "STUB: not implemented"; return }

func (bl *Bloom) IsSet(idx uint64) bool { _ = "STUB: not implemented"; return false }

type bloomJSONImExport struct {
	FilterSet []byte
	SetLocs   uint64
}

func newWithBoolset(bs *[]byte, locs uint64) *Bloom { _ = "STUB: not implemented"; return nil }

func JSONUnmarshal(dbData []byte) (*Bloom, error) { _ = "STUB: not implemented"; return nil, nil }

func (bl Bloom) JSONMarshal() []byte { _ = "STUB: not implemented"; return nil }
