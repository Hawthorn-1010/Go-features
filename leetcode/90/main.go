package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{
		1, 2, 2,
	}
	fmt.Println(subsetsWithDup(nums))
}

var (
	res    [][]int
	path   []int
	status []bool // 是否已经使用
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res = make([][]int, 0)
	path = make([]int, 0)
	status = make([]bool, len(nums))
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] && status[i-1] == false {
			i++
		}
		if i == len(nums) {
			break
		}
		path = append(path, nums[i])
		status[i] = true
		dfs(nums, i+1)
		path = path[:len(path)-1]
		status[i] = false
	}
}
