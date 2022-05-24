package main

import (
	"fmt"
	"log"
	"time"
)

// TickerDemo 用于演示 Ticker 的基础用法
func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

// TickerLaunch 用于演示 Ticker 聚合任务用法
func TickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 30 // 没车最大装载人数
	passengers := make([]string, 0, maxPassenger)

	for {
		passenger := GetNewPassenger() // 获取一个新乘客
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}

		select {
		case <-ticker.C: // 时间到，发车
			Launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger { // 时间没到，车已坐满，发车
				Launch(passengers)
				passengers = []string{}
			}
		}
	}
}

func GetNewPassenger() string {
	return "1"
}

func Launch(passengers []string) {
	fmt.Println(passengers)
}

func main() {
	TickerDemo()
}
