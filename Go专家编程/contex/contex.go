package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

func cancelCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()

	// Just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

func timerCtx() {
	// 创建一个 5s 超时的 context，并将其传递给子协程，5s 后自动关闭 context。
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandelRequest(ctx)

	time.Sleep(15 * time.Second)
}

func main() {
	// cancelCtx()
	timerCtx()
}

func HandelRequestValueCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func valueCtx() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequestValueCtx(ctx)

	time.Sleep(10 * time.Second)
}