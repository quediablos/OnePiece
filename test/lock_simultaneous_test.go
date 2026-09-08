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

func TestConcurrentLockUnlock(t *testing.T) {
	var wg sync.WaitGroup

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
			time.Sleep(3 * time.Second)

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
}
