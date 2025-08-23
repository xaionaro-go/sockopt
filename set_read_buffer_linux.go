//go:build linux
// +build linux

package sockopt

import (
	"golang.org/x/sys/unix"
)

func SetReadBuffer(fd int, bufferSize int) error {
	err := unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_RCVBUFFORCE, bufferSize)
	if err == nil {
		return nil
	}
	return unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_RCVBUF, bufferSize)
}
