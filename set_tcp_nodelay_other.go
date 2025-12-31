//go:build !linux
// +build !linux

package sockopt

func SetTCPNodelay(fd int, enable bool) error {
	return ErrNotImplemented{}
}