package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int, 3)
	m["apple"] = 1
	m["banana"] = 2
	m["mango"] = 3

	s := make([]string, 0)

	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)

	for _, key := range s {
		fmt.Println(key, m[key])
	}

}
