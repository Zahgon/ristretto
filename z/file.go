package z

import (
	"errors"
	"io"
	"os"
)

type MmapFile struct {
	Data []byte
	Fd   *os.File
}

var NewFile = errors.New("Create a new file")

func OpenMmapFileUsing(fd *os.File, sz int, writable bool) (*MmapFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenMmapFile(filename string, flag int, maxSz int) (*MmapFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type mmapReader struct {
	Data   []byte
	offset int
}

func (mr *mmapReader) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *MmapFile) NewReader(offset int) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (m *MmapFile) Bytes(off, sz int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *MmapFile) Slice(offset int) []byte { _ = "STUB: not implemented"; return nil }

func (m *MmapFile) AllocateSlice(sz, offset int) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MmapFile) Sync() error { _ = "STUB: not implemented"; return nil }

func (m *MmapFile) Delete() error { _ = "STUB: not implemented"; return nil }

func (m *MmapFile) Close(maxSz int64) error { _ = "STUB: not implemented"; return nil }

func SyncDir(dir string) error { _ = "STUB: not implemented"; return nil }
