package test

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

const serverURL = "http://127.0.0.1:3003"

var sum = 0

func TestConcurrentLockUnlock(t *testing.T) {
	var wg sync.WaitGroup

	sum = 0

	for id := 0; id < 10; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			resource := "x"

			// Step 1: Acquire lock
			lockResp, err := http.Get(fmt.Sprintf("%s/lock/%s", serverURL, resource))
			if err != nil {
				t.Errorf("[goroutine %d] lock request failed: %v", id, err)
				return
			}
			body, _ := io.ReadAll(lockResp.Body)
			lockResp.Body.Close()
			fmt.Printf("[%s] [goroutine %d] Lock acquired for resource %s — response: %s\n", time.Now().Format(time.RFC3339), id, resource, string(body))

			// Step 2: Hold the lock for 3 seconds
			fmt.Printf("[%s] [goroutine %d] Sleeping for 3 seconds while holding lock on resource %s\n", time.Now().Format(time.RFC3339), id, resource)
			//time.Sleep(3 * time.Second)

			// Increment the sum.
			sum++

			// Step 3: Release lock
			unlockResp, err := http.Get(fmt.Sprintf("%s/unlock/%s", serverURL, resource))
			if err != nil {
				t.Errorf("[goroutine %d] unlock request failed: %v", id, err)
				return
			}
			body, _ = io.ReadAll(unlockResp.Body)
			unlockResp.Body.Close()
			fmt.Printf("[%s] [goroutine %d] Lock released for resource %s — response: %s\n", time.Now().Format(time.RFC3339), id, resource, string(body))
		}(id)
	}

	wg.Wait()
	fmt.Printf("[%s] All goroutines finished.\n", time.Now().Format(time.RFC3339))

	if sum != 10 {
		t.Errorf("expected sum to be 10, got %d", sum)
	}
}

func TestLockExpiryMechanism(t *testing.T) {
	var wg sync.WaitGroup
	resource := "expiry-test"

	sum = 0

	for id := 0; id < 3; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Step 1: Request the lock (blocks until acquired).
			fmt.Printf("[%s] [goroutine %d] Sending lock request for resource %s\n", time.Now().Format(time.RFC3339), id, resource)
			lockResp, err := http.Get(fmt.Sprintf("%s/lock/%s", serverURL, resource))
			if err != nil {
				t.Errorf("[goroutine %d] lock request failed: %v", id, err)
				return
			}
			body, _ := io.ReadAll(lockResp.Body)
			lockResp.Body.Close()
			fmt.Printf("[%s] [goroutine %d] Lock acquired for resource %s — response: %s\n", time.Now().Format(time.RFC3339), id, resource, string(body))

			sum++

			// Step 2: Do NOT call unlock — wait for the server's expiry mechanism to release the lock.
			fmt.Printf("[%s] [goroutine %d] Holding lock without unlocking; waiting for server-side expiry\n", time.Now().Format(time.RFC3339), id)
		}(id)
	}

	wg.Wait()
	fmt.Printf("[%s] All goroutines finished — all locks released via expiry.\n", time.Now().Format(time.RFC3339))

	if sum != 3 {
		t.Errorf("expected sum to be 3, got %d", sum)
	}
}
