package message

import (
	"fmt"
	"strconv"
)

func GenerateAcquireLockResponse(resourceId string) string {

	body := "Acquired lock for resourceId: " + resourceId
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateReleaseLockResponse(resourceId string) string {

	body := "Released lock for resourceId: " + resourceId
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}
