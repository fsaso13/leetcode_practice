package main

import (
	"fmt"
)

func rotate(nums []int, k int) {

	if len(nums) == 1 {
		return
	}

	if len(nums) == 2 {
		if k%2 == 0 {
			return
		} else {
			nums[0], nums[1] = nums[1], nums[0]
			return
		}
	}

	if k > len(nums) {
		k = k % len(nums)
	}
	cut := len(nums) - k

	lhs := nums[:cut]
	rhs := nums[cut:]

	rhs = append(rhs, lhs...)
	copy(nums, rhs)
}

func main() {
	fmt.Println("Potato")

	// Input: nums = [1,2,3,4,5,6,7], k = 3
	// Output: [5,6,7,1,2,3,4]
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	// Input: nums = [-1,-100,3,99], k = 2
	// Output: [3,99,-1,-100]
	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)

	// Input: nums = [1,2], k = 1
	// Output: [2,1]
	nums = []int{1, 2}
	rotate(nums, 1)
	fmt.Println(nums)
}
