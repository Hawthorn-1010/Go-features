package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	res := make([]int, len(array))

	//pos := len(array) - 1
	respos := 0

	for i := 13; i >= 1; i-- {
		res[respos] = i
		if i == 1 {
			break
		}
		j := respos + 1
		for flag := 0; flag < 2; j++ {
			if j > 12 {
				j = 0
			}
			if res[j] == 0 {
				flag++
			}
		}
		respos = j - 1
	}

	fmt.Println(res)
}

//func main() {
//	a := 5
//	b := 6
//	fmt.Println(a, b)
//	a, b, c := 1, 2, 4
//	fmt.Println(a, b, c)
//}

//func main() {
//	magazine := "happy"
//	//for _, ch := range magazine {
//	//	fmt.Println(reflect.TypeOf(ch))
//	//	fmt.Println(reflect.TypeOf(ch - 'a'))
//	//}
//
//	// rune int32
//	container := make(map[rune]int)
//	for _, ch := range magazine {
//		container[ch-'a']++
//	}
//	fmt.Println(container)
//	s := "hello world"
//	fmt.Println("the length of string is:", len(s))
//
//	var a [6]int
//	for index := range a {
//		fmt.Scanln(&a[index])
//	}
//	for index, value := range a {
//		fmt.Println(index, ":", value)
//	}
//}

//func main() {
//	s := "a1b2c3"
//	var result string
//	for _, char := range s {
//		if unicode.IsDigit(char) {
//			result += "number"
//		} else {
//			result += string(char)
//		}
//	}
//	fmt.Println(result)
//
//	var builder strings.Builder
//	for i := 0; i < 10; i++ {
//		builder.WriteString("hello")
//	}
//	result = builder.String()
//	fmt.Println(result)
//}

//func main() {
//	//a, b, c := 1, 1, 1
//	//for i := 0; i < 100; i++ {
//	//	temp := 3*a + 2*b + c
//	//	a = b
//	//	b = c
//	//	c = temp
//	//	fmt.Println(temp)
//	//}
//	//s := "你好世界！"
//	//for index, value := range s {
//	//	fmt.Println(index, value)
//	//}
//	//for i := 0; i < len(s); i++ {
//	//	fmt.Println(i, s[i])
//	//}
//
//	str := "hhhh"
//	for _, s := range str {
//		fmt.Println(reflect.TypeOf(s))
//		ss := string(s)
//		fmt.Println(reflect.TypeOf(ss))
//	}
//}

//func main() {
//	nums := []int{1, 2, 3}
//	fmt.Println(permute(nums))
//}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	path := make([]int, 0)
	status := make([]bool, length)

	dfs(nums, &res, path, status, length)

	return res
}

func dfs(nums []int, res *[][]int, path []int, status []bool, length int) {
	if len(path) == length {
		tmp := make([]int, length)
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	for i := 0; i < length; i++ {
		if !status[i] {
			path = append(path, nums[i])
			status[i] = true
			dfs(nums, res, path, status, length)
			path = path[:len(path)-1]
			status[i] = false
		}
	}
}
