package main

import "fmt"

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1)
	a := firstMissingPositive(nums)
	fmt.Println(a)
}

func firstMissingPositive(nums []int) int {
	// 常数额外空间，所以要利用原数组
	for i := 0; i < len(nums); {
		if nums[i] == i+1 || nums[i] == 0 {
			i++
			continue
		}
		if nums[i] > 0 && nums[i] <= len(nums) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return 0
}

func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])
	ans := 0
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans * ans
}
