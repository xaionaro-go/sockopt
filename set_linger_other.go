//go:build !linux
// +build !linux

// Package sockopt provides functions to set socket options.
package sockopt

func SetLinger(fd int, onoff int32, linger int32) error {
	return ErrNotImplemented{}
}
