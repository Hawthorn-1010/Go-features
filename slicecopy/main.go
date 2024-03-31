package main

import "fmt"

/*
question describe: copy实际复制的元素个数是从两个slice中取最小值，即min(len(dst), len(src))，如果len(dst)=0则没有办法完成复制
*/

//func main() {
//	src := []int{1, 2, 3}
//	dst := make([]int, 0)
//	copy(dst, src)
//	fmt.Print(dst) // []
//}

//func main() {
//	src := []int{1, 2, 3}
//	dst := make([]int, len(src))
//	copy(dst, src)
//	fmt.Print(dst) //[1 2 3]
//}

/*
question describe: 使用range遍历切片时会先拷贝一份，然后再遍历拷贝数据
*/

type user struct {
	name string
	age  uint64
}

func main() {
	u := []user{
		{"Alice", 23},
		{"Bob", 19},
	}
	for _, v := range u {
		if v.age != 18 {
			v.age = 20
		}
	}
	fmt.Println(u)

	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10 // data 中原有元素是不会被修改的
	}
	fmt.Println("data: ", data) // data:  [1 2 3]
}
