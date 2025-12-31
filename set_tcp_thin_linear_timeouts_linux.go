//go:build linux
// +build linux

package sockopt

import (
	"golang.org/x/sys/unix"
)

func SetTCPThinLinearTimeouts(fd int, enable bool) error {
	val := 0
	if enable {
		val = 1
	}
	return unix.SetsockoptInt(fd, unix.IPPROTO_TCP, unix.TCP_THIN_LINEAR_TIMEOUTS, val)
}