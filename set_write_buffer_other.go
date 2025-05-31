//go:build !linux
// +build !linux

package sockopt

func SetWriteBuffer(fd int, bufferSize int) error {
	return ErrNotImplemented{}
}
