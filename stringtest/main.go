package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	s := "abbca"
//	t := "bbcaa"
//	s1, s2 := []byte(s), []byte(t)
//	fmt.Println(s1)
//	fmt.Println(s2)
//	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
//	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
//	fmt.Println(string(s1) == string(s2))
//	ss := make([]int, 0)
//	fmt.Println(ss == nil)
//	intersection()
//
//	m := make(map[string]int)
//	m["hello"] = 1
//	fmt.Println(m["right"])
//	if v, ok := m["key"]; !ok {
//		fmt.Println(v, ok)
//	}
//}
//
//func intersection() (res []int) {
//	m := []int{6, 5, 4, 3, 2, 1}
//	for _, num := range m {
//		fmt.Println(num)
//	}
//	k := make([]int, 0)
//	k[0] = 1
//	fmt.Println(len(res))
//	fmt.Println(res == nil)
//	return append(res, 1)
//}

//func main() {
//	//s := "  a good   example  "
//	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	//strs := strings.Split(s, " ")
//	fmt.Println(reflect.TypeOf(s))
//	fmt.Println(reflect.TypeOf(s[5:7]))
//
//	var k int
//	var str string
//	fmt.Scanln(&k)
//	fmt.Scanln(&str)
//	var res string
//	res += str[len(str)-k:]
//	res += str[0 : len(str)-k]
//	fmt.Println(res)
//}

//func main() {
//	var k int
//	var str string
//
//	fmt.Scanln(&str)
//	fmt.Scanln(&k)
//	fmt.Println(str)
//	fmt.Println(k)
//	words := []byte(str)
//	for _, word := range words {
//		fmt.Println(word, ' ' == word)
//	}
//	//reverse(words)
//	//reverse(words[0:k])
//	//reverse(words[k:])
//	//fmt.Println(string(words))
//}
//
//func reverse(word []byte) {
//	left, right := 0, len(word)-1
//	for left < right {
//		word[left], word[right] = word[right], word[left]
//		left++
//		right--
//	}
//}

//	func main() {
//		//var str string
//		//fmt.Println("请输入一个字符串:")
//		//fmt.Scanln(&str)
//		//fmt.Println("您输入的字符串是:", str)
//		reader := bufio.NewReader(os.Stdin)
//		line, _, _ := reader.ReadLine()
//
//		fmt.Println(line)
//		fmt.Println(string(line))
//	}

//func main() {
//	haystack := "sadbutsad"
//	needle := "sad"
//	fmt.Println(haystack[0:4] == needle)
//}

func main() {
	s := "hello"
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			fmt.Println(reflect.TypeOf(' '))
			stack = append(stack, s[i])
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}
}
