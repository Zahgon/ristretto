//go:build amd64 || arm64 || arm64be || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x || sparc64
// +build amd64 arm64 arm64be ppc64 ppc64le mips64 mips64le riscv64 s390x sparc64

package z

const (
	MaxArrayLen = 1<<50 - 1

	MaxBufferSize = 256 << 30
)
