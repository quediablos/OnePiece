package test

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func TestCreateStock(t *testing.T) {

	conn, err := net.Dial("tcp", "localhost:3003")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 1|REQ|CREATE_STOCK|<resourceId>|<quantity>
	request := "1|REQ|CREATE_STOCK|ITEM_A|100"
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}

	// Read the raw response back from the socket
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
	}
}

func TestCreateStockDuplicate(t *testing.T) {

	conn, err := net.Dial("tcp", "localhost:3003")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 1|REQ|CREATE_STOCK|<resourceId>|<quantity>
	request := "1|REQ|CREATE_STOCK|ITEM_A|100"
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}

	// Read the raw response back from the socket
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
	}

	// --- Duplicate case: sending the same CREATE_STOCK should return FAILED ---
	conn2, err := net.Dial("tcp", "localhost:3003")
	if err != nil {
		fmt.Println("Error connecting (duplicate case):", err)
		os.Exit(1)
	}
	defer conn2.Close()

	_, err = conn2.Write([]byte(request))
	if err != nil {
		fmt.Println("Error writing data (duplicate case):", err)
		return
	}

	// Expected: 1|RES|FAILED|CREATE_STOCK|ITEM_A|STOCK_ALREADY_CREATED|Stock already created.
	const expectedDuplicateResponse = "1|RES|FAILED|CREATE_STOCK|ITEM_A|STOCK_ALREADY_CREATED|Stock already created."
	scanner2 := bufio.NewScanner(conn2)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
		if scanner2.Text() != expectedDuplicateResponse {
			t.Errorf("expected %q, got %q", expectedDuplicateResponse, scanner2.Text())
		}
	}

	if err := scanner2.Err(); err != nil {
		fmt.Println("Error reading response (duplicate case):", err)
	}
}

func TestStockReserveAndRelease(t *testing.T) {
	resourceId := "ITEM_C"

	// Step 1: Create stock with quantity of 1
	resp := sendTcpRequest(t, fmt.Sprintf("1|REQ|CREATE_STOCK|%s|1", resourceId))
	fmt.Println("Step 1 (create):", resp)

	// Step 2: Reserve stock — save the reserve ID from the response
	// Expected response: 1|RES|SUCCESSFUL|RESERVE_STOCK|<resourceId>|<reserveId>
	resp = sendTcpRequest(t, fmt.Sprintf("1|REQ|RESERVE_STOCK|%s", resourceId))
	fmt.Println("Step 2 (reserve):", resp)
	parts := strings.Split(resp, "|")
	if len(parts) < 6 || parts[2] != "SUCCESSFUL" {
		t.Fatalf("step 2: expected successful reserve, got: %q", resp)
	}
	reserveId := parts[5]

	// Step 3: Reserve stock again — stock is exhausted, expect FAILED
	resp = sendTcpRequest(t, fmt.Sprintf("1|REQ|RESERVE_STOCK|%s", resourceId))
	fmt.Println("Step 3 (reserve again — expect fail):", resp)
	expectedFailed := fmt.Sprintf("1|RES|FAILED|RESERVE_STOCK|%s", resourceId)
	if resp != expectedFailed {
		t.Errorf("step 3: expected %q, got %q", expectedFailed, resp)
	}

	// Step 4: Release stock using the reserve ID from step 2
	resp = sendTcpRequest(t, fmt.Sprintf("1|REQ|RELEASE_STOCK|%s|%s", resourceId, reserveId))
	fmt.Println("Step 4 (release):", resp)

	// Step 5: Reserve stock — stock is available again, expect SUCCESSFUL
	resp = sendTcpRequest(t, fmt.Sprintf("1|REQ|RESERVE_STOCK|%s", resourceId))
	fmt.Println("Step 5 (reserve after release):", resp)
	expectedPrefix := fmt.Sprintf("1|RES|SUCCESSFUL|RESERVE_STOCK|%s|", resourceId)
	if !strings.HasPrefix(resp, expectedPrefix) {
		t.Errorf("step 5: expected response starting with %q, got %q", expectedPrefix, resp)
	}
}

func sendTcpRequest(t *testing.T, request string) string {
	t.Helper()
	conn, err := net.Dial("tcp", "localhost:3003")
	if err != nil {
		t.Fatalf("error connecting: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(request))
	if err != nil {
		t.Fatalf("error writing data: %v", err)
	}

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("error reading response: %v", err)
	}
	return ""
}
