package test

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
)

func Test(t *testing.T) {

	conn, err := net.Dial("tcp", "localhost:3003")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	request := "1|REQ|LOCK|X"
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}

	// 3. Read the raw response back from the socket
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
	}
}
