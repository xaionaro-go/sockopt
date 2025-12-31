//go:build !linux
// +build !linux

package sockopt

func SetTCPQuickack(fd int, enable bool) error {
	return ErrNotImplemented{}
}