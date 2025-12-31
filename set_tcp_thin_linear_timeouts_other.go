//go:build !linux
// +build !linux

package sockopt

func SetTCPThinLinearTimeouts(fd int, enable bool) error {
	return ErrNotImplemented{}
}