//go:build linux
// +build linux

package sockopt

import (
	"golang.org/x/sys/unix"
)

func SetLinger(fd int, onoff int32, linger int32) error {
	return unix.SetsockoptLinger(fd, unix.SOL_SOCKET, unix.SO_LINGER, &unix.Linger{
		Onoff:  onoff,
		Linger: linger,
	})
}
