package main

import (
	"fmt"
	"time"
)

// how to jump out of for select loop

func main() {
	chExit := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			chExit <- true
			time.Sleep(2 * time.Second)
		}
		close(chExit)
	}()
EXIT:
	for {
		select {
		case v, ok := <-chExit:
			if !ok {
				fmt.Println("close channel 2", v)
				break EXIT //goto EXIT2
			}

			fmt.Println("ch2 val =", v)
		}
	}

	//EXIT2:
	fmt.Println("exit testSelectFor2")
}
