//go:build (arm64 || arm || s390x) && linux && !js
// +build arm64 arm s390x
// +build linux
// +build !js

package z

func mremap(data []byte, size int) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint:lll
	return nil, nil
}
