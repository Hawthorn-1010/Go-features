package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 10
	value, ok := x.(string)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(value)
	fmt.Println(reflect.TypeOf(value))

	if ok {
		fmt.Println("x is an int:", value)
	} else {
		fmt.Println("x is not an int")
	}

}
