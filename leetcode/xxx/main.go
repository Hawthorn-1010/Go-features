package main

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	search(nums, target)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 分成连续的和不连续的
		// 数组左边连续
		if nums[left] < nums[mid] {
			if nums[left] < target && target < nums[mid] {
				right = mid
			} else {
				left = mid
			}
		} else {
			if nums[mid] < target && target < nums[right] {
				left = mid
			} else {
				right = mid
			}
		}
	}
	return -1
}
