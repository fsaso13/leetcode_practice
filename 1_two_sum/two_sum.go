package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	difMap := make(map[int]int)

	for ind, val := range nums {
		difMap[val] = ind
	}

	result := []int{0, 1}
	for ind, val := range nums {
		lookup := target - val

		pos, ok := difMap[lookup]

		if ok {

			if pos == ind {
				continue
			} else {
				result = []int{ind, pos}
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Potato")

	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	// Input: nums = [3,2,4], target = 6
	// Output: [1,2]
	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

	// Input: nums = [3,3], target = 6
	// Output: [0,1]
	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))

	// Input: nums = [3,2,3], target = 6
	// Output: [0,1]
	nums = []int{3, 2, 3}
	fmt.Println(twoSum(nums, 6))

}
