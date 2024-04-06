package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	addFunc := getSequence()
	fmt.Println(addFunc())
	fmt.Println(addFunc())

	addFunc2 := getSequence()
	fmt.Println(addFunc2())
}
