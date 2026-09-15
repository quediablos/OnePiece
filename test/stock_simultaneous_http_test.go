package test

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestStockReserve(t *testing.T) {
	const stockServerURL = "http://127.0.0.1:3003"
	const resource = "stock-x"
	const totalGoroutines = 15
	const stockQuantity = 10

	// Step 1: Create stock with quantity 10.
	fmt.Printf("[%s] Creating stock %s with quantity %d\n", time.Now().Format(time.RFC3339), resource, stockQuantity)
	createResp, err := http.Get(fmt.Sprintf("%s/create_stock/%s/%d", stockServerURL, resource, stockQuantity))
	if err != nil {
		t.Fatalf("Failed to create stock: %v", err)
	}
	io.ReadAll(createResp.Body)
	createResp.Body.Close()
	fmt.Printf("[%s] Stock created\n", time.Now().Format(time.RFC3339))

	// Step 2: Launch 15 goroutines to reserve stock concurrently.
	var wg sync.WaitGroup
	var mu sync.Mutex
	successCount := 0
	failureCount := 0

	for id := 0; id < totalGoroutines; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("[%s] [goroutine %d] Sending reserve request for resource %s\n", time.Now().Format(time.RFC3339), id, resource)
			resp, err := http.Get(fmt.Sprintf("%s/reserve_stock/%s", stockServerURL, resource))
			if err != nil {
				t.Errorf("[goroutine %d] reserve request failed: %v", id, err)
				return
			}
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			mu.Lock()
			if resp.StatusCode == 200 {
				successCount++
				fmt.Printf("[%s] [goroutine %d] Reserve SUCCESSFUL — response: %s\n", time.Now().Format(time.RFC3339), id, string(body))
			} else {
				failureCount++
				fmt.Printf("[%s] [goroutine %d] Reserve FAILED (depleted) — response: %s\n", time.Now().Format(time.RFC3339), id, string(body))
			}
			mu.Unlock()
		}(id)
	}

	wg.Wait()
	fmt.Printf("[%s] All goroutines finished. Successes: %d, Failures: %d\n", time.Now().Format(time.RFC3339), successCount, failureCount)

	// Step 3: Verify exactly 10 succeeded and 5 failed.
	if successCount != stockQuantity {
		t.Errorf("expected %d successful reserves, got %d", stockQuantity, successCount)
	}
	if failureCount != totalGoroutines-stockQuantity {
		t.Errorf("expected %d failed reserves, got %d", totalGoroutines-stockQuantity, failureCount)
	}
}
