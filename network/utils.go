package network

import (
	"net"
	"time"
)

// IsConnOpen Returns true if the connection is open.
func IsConnOpen(conn net.Conn) bool {
	one := make([]byte, 1)
	conn.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
	_, err := conn.Read(one)
	conn.SetReadDeadline(time.Time{}) // reset
	if err == nil {
		return true // data was available, conn is alive
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true // timed out = still open, just no data
	}
	return false // io.EOF, "connection reset by peer", etc.
}
