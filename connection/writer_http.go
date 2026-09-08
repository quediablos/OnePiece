package connection

import "net"

func WriteResponseHttp(conn net.Conn, response string) {

	conn.Write([]byte(response))
	conn.Close()
}
