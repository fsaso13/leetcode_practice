package main

import (
	"fmt"
	"math"
)

func findMaxK(nums []int) int {
	exist := map[int]int{}
	max := -1
	for _, val := range nums {
		_, ok := exist[val]
		if !ok {
			exist[val] = 1
			_, ok2 := exist[-1*val]
			if ok2 {
				absV := math.Abs(float64(val))
				if int(absV) > max {
					max = int(absV)
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println("Potato")

	// Input: nums = [-1,2,-3,3]
	// Output: 3
	nums := []int{-1, 2, -3, 3}
	fmt.Println(findMaxK(nums))

	// Input: nums = [-1,10,6,7,-7,1]
	// Output: 7
	nums = []int{-1, 10, 6, 7, -7, 1}
	fmt.Println(findMaxK(nums))

	// Input: nums = [-10,8,6,7,-2,-3]
	// Output: -1
	nums = []int{-10, 8, 6, 7, -2, -3}
	fmt.Println(findMaxK(nums))

	// Input: nums = [-9,-43,24,-23,-16,-30,-38,-30]
	// Output: -1
	nums = []int{-9, -43, 24, -23, -16, -30, -38, -30}
	fmt.Println(findMaxK(nums))
}
