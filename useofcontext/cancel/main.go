package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	// 假设该操作因某种原因而失败
	// 下面模拟业务执行一定时间
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// 该方法要么正常执行完成
	// 要么取消，不再继续执行
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	// 创建上下文
	ctx := context.Background()

	// 基于上下文创建需求上下文
	ctx, cancel := context.WithCancel(ctx)

	// 在不同协程中执行两个操作
	go func() {
		err := operation1(ctx)
		// 如果该方法返回错误，则取消该上下文中的后续操作
		if err != nil {
			cancel()
		}
	}()

	// 实用相同上下文执行操作2
	operation2(ctx)
}
