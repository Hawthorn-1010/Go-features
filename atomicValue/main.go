package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var v atomic.Value

	// 设置存储的值
	v.Store("Hello, World!")

	// 获取存储的值
	val := v.Load()
	fmt.Println("Stored value:", val)

	// 原子交换存储的值
	oldVal := v.Swap("Goodbye, World!")
	fmt.Println("Previous value:", oldVal)

	// 再次获取存储的值
	newVal := v.Load()
	fmt.Println("New value:", newVal)
}
