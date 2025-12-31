package sockopt_test

import (
	"net"

	"github.com/xaionaro-go/sockopt"
)

func Example() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	tcpConn := conn.(*net.TCPConn)
	sc, err := tcpConn.SyscallConn()
	if err != nil {
		panic(err)
	}

	var setErr error
	err = sc.Control(func(fd uintptr) {
		if err := sockopt.SetTCPNodelay(int(fd), true); err != nil {
			setErr = err
			return
		}
		if err := sockopt.SetTCPQuickack(int(fd), true); err != nil {
			setErr = err
			return
		}
		if err := sockopt.SetTCPThinDupack(int(fd), true); err != nil {
			setErr = err
			return
		}
		if err := sockopt.SetTCPThinLinearTimeouts(int(fd), true); err != nil {
			setErr = err
			return
		}
	})
	if err != nil {
		panic(err)
	}
	if setErr != nil {
		panic(setErr)
	}
}
