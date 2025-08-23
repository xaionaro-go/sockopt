//go:build !linux
// +build !linux

package sockopt

func SetReadBuffer(fd int, bufferSize int) error {
	return ErrNotImplemented{}
}
