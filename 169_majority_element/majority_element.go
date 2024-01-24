package main

import (
	"fmt"
)

func majorityElement(nums []int) int {

	countMap := make(map[int]int)

	for _, val := range nums {
		_, keyExist := countMap[val]
		if keyExist {
			countMap[val]++
			// fmt.Printf("%d: %d vs %d\n", val, countMap[val], len(nums)/2)
			if countMap[val] > len(nums)/2 {
				return val
			}
		} else {
			countMap[val] = 1
		}
	}

	return nums[0]
}

func main() {
	fmt.Println("potato")

	// Input: nums = [3,2,3]
	// Output: 3
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))

	// Input: nums = [2,2,1,1,1,2,2]
	// Output: 2
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))

	// Input: nums = [6,5,5]
	// Output: 5
	nums = []int{6, 5, 5}
	fmt.Println(majorityElement(nums))
}
