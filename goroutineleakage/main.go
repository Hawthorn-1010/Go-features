package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
avoid goroutine leakage: use context
*/

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//
//	ch := func(ctx context.Context) <-chan int {
//		ch := make(chan int)
//		go func() {
//			for i := 0; ; i++ {
//				select {
//				case <-ctx.Done():
//					return
//				case ch <- i:
//				}
//			}
//		}()
//		return ch
//	}(ctx)
//
//	for v := range ch {
//		fmt.Println(v)
//		if v == 5 {
//			cancel()
//			break
//		}
//	}
//}

/*
cases of goroutine leakage:https://segmentfault.com/a/1190000040161853
*/

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		ch <- struct{}{}
	}()

	time.Sleep(time.Second)
}
