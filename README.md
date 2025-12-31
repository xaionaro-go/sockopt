# sockopt

A Go package for setting various socket options on file descriptors.

## Supported Options

### Socket Level Options

- `SetLinger(fd int, onoff int32, linger int32) error` - Set the SO_LINGER option.
- `SetReadBuffer(fd int, bufferSize int) error` - Set the SO_RCVBUF option (tries SO_RCVBUFFORCE first).
- `SetWriteBuffer(fd int, bufferSize int) error` - Set the SO_SNDBUF option (tries SO_SNDBUFFORCE first).

### TCP Level Options

- `SetTCPNodelay(fd int, enable bool) error` - Enable or disable TCP_NODELAY (Nagle's algorithm).
- `SetTCPQuickack(fd int, enable bool) error` - Enable or disable TCP_QUICKACK.
- `SetTCPThinDupack(fd int, enable bool) error` - Enable or disable TCP_THIN_DUPACK.
- `SetTCPThinLinearTimeouts(fd int, enable bool) error` - Enable or disable TCP_THIN_LINEAR_TIMEOUTS.

## Platform Support

- Linux: All options are supported.
- Other platforms: Only socket level options are supported; TCP options return `ErrNotImplemented{}`.

## Usage

To use with a raw file descriptor:

```go
// Disable Nagle's algorithm
if err := sockopt.SetTCPNodelay(fd, true); err != nil {
    // handle error
}
```

To get the file descriptor from a `net.TCPConn`:

```go
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
```

## Error Handling

Functions return an error if the operation fails. On unsupported platforms, TCP options return `ErrNotImplemented{}`.