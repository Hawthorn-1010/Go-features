package main

import (
	"fmt"
	"sync"
)

/*
Golang循环中的闭包和go func变量取值问题
https://go.dev/wiki/CommonMistakes
*/

type Student struct {
	Name string
}

func main() {
	list := []*Student{{"lily"}, {"jack"}, {"alice"}}
	var wg sync.WaitGroup
	// 1. wrong case
	//for _, v := range list {
	//	wg.Add(1)
	//	go func() {
	//		fmt.Println(v.Name)
	//		wg.Done()
	//	}()
	//}

	// 2. 解决方式一
	for _, v := range list {
		wg.Add(1)
		go func(val string) {
			fmt.Println(val)
			wg.Done()
		}(v.Name)
	}
	// 3. 解决方式二
	//for _, v := range list {
	//	wg.Add(1)
	//	val := v
	//	go func() {
	//		fmt.Println(val.Name)
	//		wg.Done()
	//	}()
	//}

	wg.Wait()
}
