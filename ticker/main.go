package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker!")

		}
	}

	stop := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second) // 等待 10 秒
		close(stop)                  // 关闭 stop 通道
	}()
}
