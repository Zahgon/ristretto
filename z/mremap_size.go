//go:build linux && !arm64 && !arm && !s390x && !js
// +build linux,!arm64,!arm,!s390x,!js

package z

func mremap(data []byte, size int) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint:lll
	return nil, nil
}
