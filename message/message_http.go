package message

import (
	"fmt"
	"strconv"
)

func GenerateAcquireLockHttpResponse(resourceId string) string {

	body := "Acquired lock for resourceId: " + resourceId
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateReleaseLockHttpResponse(resourceId string) string {

	body := "Released lock for resourceId: " + resourceId
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateReserveStockSuccessfulHttpResponse(reserveId string, resourceId string) string {
	body := "Reserved stock for resourceId: " + resourceId + " reserveId:" + reserveId
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateReserveStockFailedHttpResponse(resourceId string) string {
	body := "Depleted stock for resourceId: " + resourceId
	response := "HTTP/1.1 204 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateCreateStockHttpResponse(resourceId string, quantity string) string {

	body := "Created stock for resourceId: " + resourceId + " quantity:" + quantity
	response := "HTTP/1.1 204 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response
}

func GenerateGenericErrorHttpResponse() string {
	body := "Error"
	response := "HTTP/1.1 204 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response

}

func GenerateErrorHttpResponse(err string) string {
	body := err
	response := "HTTP/1.1 204 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(body))) +
		"\r\n" +
		body

	return response

}
