package main

import (
	"fmt"
	"sync"
)

//func main() {
//	defer func() {
//		fmt.Println("recovered: ", recover())
//	}()
//	panic("not good")
//	println("ok")
//}

// 错误的 recover 调用示例
//func main() {
//	recover()         // 什么都不会捕捉
//	panic("not good") // 发生 panic，主程序退出
//	recover()         // 不会被执行
//	println("ok")
//}

// 嵌套崩溃
//func main() {
//	defer fmt.Println("in main")
//	defer func() {
//		defer func() {
//			panic("panic again and again")
//		}()
//		panic("panic again")
//	}()
//
//	panic("panic once")
//}

var wg sync.WaitGroup

func main() {
	defer func() {
		fmt.Println("recover in main:", recover())
	}()
	wg.Add(1)
	go subFunc()
	wg.Wait()
}

func subFunc() {
	defer func() {
		wg.Done()
		fmt.Println("recover in sub:", recover())
	}()
	panic("not good")
}
