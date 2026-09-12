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

func GenerateReserveStockSuccessfulResponse(reserveId string, resourceId string) string {
	body := "Reserved stock for resourceId: " + resourceId + " reserveId:" + reserveId
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateReserveStockFailedResponse(resourceId string) string {
	body := "Depleted stock for resourceId: " + resourceId
	response := "HTTP/1.1 204 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}
