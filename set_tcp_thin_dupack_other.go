//go:build !linux
// +build !linux

package sockopt

func SetTCPThinDupack(fd int, enable bool) error {
	return ErrNotImplemented{}
}