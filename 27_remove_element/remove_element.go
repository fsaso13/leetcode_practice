package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	lastElement := len(nums) - 1
	counter := 0
	if nums[lastElement] == val {
		lastElement--
		counter++
	}

	for i := 0; i <= lastElement; i++ {
		if nums[i] == val {
			for {
				if nums[lastElement] == val {
					lastElement--
					counter++
				} else {
					break
				}

				if lastElement < 0 || lastElement < i {
					return len(nums) - counter
				}
			}

			nums[i], nums[lastElement] = nums[lastElement], nums[i]
			lastElement--
			counter++
		}
		// fmt.Println(nums)
	}
	return len(nums) - counter
}

func main() {
	fmt.Println("potato")
	// Input: nums = [3,2,2,3], val = 3
	// Output: 2, nums = [2,2,_,_]
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))

	// Input: nums = [0,1,2,2,3,0,4,2], val = 2
	// Output: 5, nums = [0,1,4,0,3,_,_,_]
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))

	// Input: nums = [1], val = 1
	// Output: 0, nums = [_]
	nums = []int{1}
	fmt.Println(removeElement(nums, 1))

	// Input: nums = [3,3], val = 3
	// Output: 0, nums = [_]
	nums = []int{3, 3}
	fmt.Println(removeElement(nums, 3))

	// Input: nums = [4,5], val = 5
	// Output: 1, nums = [4]
	nums = []int{4, 5}
	fmt.Println(removeElement(nums, 5))

	// Input: nums = [], val = 0
	// Output: 0, nums = []
	nums = []int{}
	fmt.Println(removeElement(nums, 0))

	// Input: nums = [2,2,2], val = 2
	// Output: 0, nums = []
	nums = []int{2, 2, 2}
	fmt.Println(removeElement(nums, 2))

	// Input: nums = [2,2,3], val = 2
	// Output: 0, nums = []
	nums = []int{2, 2, 3}
	fmt.Println(removeElement(nums, 2))
}
