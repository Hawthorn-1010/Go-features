package main

import (
	"fmt"
	"reflect"
)

// 定义一个接口
type MyInterface interface {
	Method() string
}

type YourInterface interface {
	Hi() string
}

// 定义一个结构体
type MyStruct struct{}

// 实现接口的方法
func (m MyStruct) Method() string {
	return "Hello from MyStruct"
}

func main() {
	// 创建一个 MyStruct 实例
	myStruct := MyStruct{}

	// 将 myStruct 转换为 interface{} 类型，并使用类型断言将其转换为 MyInterface 类型
	// 判断 MyStruct 是否实现了 MyInterface 接口
	_, ok := interface{}(myStruct).(MyInterface)
	if ok {
		fmt.Println("MyStruct 实现了 MyInterface 接口")
	} else {
		fmt.Println("MyStruct 没有实现 MyInterface 接口")
	}

	// 获取接口的类型对象
	myInterfaceType := reflect.TypeOf((*MyInterface)(nil)).Elem()

	// 获取结构体的类型对象
	myStructType := reflect.TypeOf(MyStruct{})

	// 判断结构体是否实现了接口
	if myStructType.Implements(myInterfaceType) {
		fmt.Println("MyStruct 实现了 MyInterface 接口")
	} else {
		fmt.Println("MyStruct 没有实现 MyInterface 接口")
	}
}
