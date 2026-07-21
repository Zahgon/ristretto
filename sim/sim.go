package sim

import (
	"errors"
	"io"
)

var (
	ErrDone = errors.New("no more values in the Simulator")

	ErrBadLine = errors.New("bad line for trace format")
)

type Simulator func() (uint64, error)

func NewZipfian(s, v float64, n uint64) Simulator {
	_ = "STUB: not implemented"
	return *new(Simulator)
}

func NewUniform(max uint64) Simulator { _ = "STUB: not implemented"; return *new(Simulator) }

type Parser func(string, error) ([]uint64, error)

func NewReader(parser Parser, file io.Reader) Simulator {
	_ = "STUB: not implemented"
	return *new(Simulator)
}

func ParseLIRS(line string, err error) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseARC(line string, err error) ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

func Collection(simulator Simulator, size uint64) []uint64 { _ = "STUB: not implemented"; return nil }

func StringCollection(simulator Simulator, size uint64) []string {
	_ = "STUB: not implemented"
	return nil
}
