package sharemem_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mu.Unlock()
			}()
			mu.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}

func TestCounterGroup(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mu.Unlock()
			}()
			mu.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("counter = %d\n", counter)
}
