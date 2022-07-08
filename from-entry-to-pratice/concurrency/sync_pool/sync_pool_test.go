package syncpool_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("Create a new object.")
			return "new"
		},
	}

	obj := "update"
	v := pool.Get().(string)
	fmt.Println(v)
	pool.Put(obj)
	runtime.GC()                      // GC 会清除 sync.pool 中缓存的对象
	time.Sleep(50 * time.Millisecond) // GC 运行需要一些时间
	v1, _ := pool.Get().(string)
	fmt.Println(v1)
}

func TestSyncPoolMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("Create a new object.")
			return "new"
		},
	}
	obj := "add"
	pool.Put(obj)
	pool.Put(obj)
	pool.Put(obj)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
		wg.Wait()
	}
}
