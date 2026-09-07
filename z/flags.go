package z

import (
	"time"
)

type SuperFlagHelp struct {
	head     string
	defaults *SuperFlag
	flags    map[string]string
}

func NewSuperFlagHelp(defaults string) *SuperFlagHelp { _ = "STUB: not implemented"; return nil }

func (h *SuperFlagHelp) Head(head string) *SuperFlagHelp { _ = "STUB: not implemented"; return nil }

func (h *SuperFlagHelp) Flag(name, description string) *SuperFlagHelp {
	_ = "STUB: not implemented"
	return nil
}

func (h *SuperFlagHelp) String() string { _ = "STUB: not implemented"; return "" }

func parseFlag(flag string) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

type SuperFlag struct {
	m map[string]string
}

func NewSuperFlag(flag string) *SuperFlag { _ = "STUB: not implemented"; return nil }

func newSuperFlagImpl(flag string) (*SuperFlag, error) { _ = "STUB: not implemented"; return nil, nil }

func (sf *SuperFlag) String() string { _ = "STUB: not implemented"; return "" }

func (sf *SuperFlag) MergeAndCheckDefault(flag string) *SuperFlag {
	_ = "STUB: not implemented"
	return nil
}

func (sf *SuperFlag) mergeAndCheckDefaultImpl(flag string) (*SuperFlag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *SuperFlag) Has(opt string) bool { _ = "STUB: not implemented"; return false }

func (sf *SuperFlag) GetDuration(opt string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (sf *SuperFlag) GetBool(opt string) bool { _ = "STUB: not implemented"; return false }

func (sf *SuperFlag) GetFloat64(opt string) float64 { _ = "STUB: not implemented"; return 0 }

func (sf *SuperFlag) GetInt64(opt string) int64 { _ = "STUB: not implemented"; return 0 }

func (sf *SuperFlag) GetUint64(opt string) uint64 { _ = "STUB: not implemented"; return 0 }

func (sf *SuperFlag) GetUint32(opt string) uint32 { _ = "STUB: not implemented"; return 0 }

func (sf *SuperFlag) GetString(opt string) string { _ = "STUB: not implemented"; return "" }

func (sf *SuperFlag) GetPath(opt string) string { _ = "STUB: not implemented"; return "" }

func expandPath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }
