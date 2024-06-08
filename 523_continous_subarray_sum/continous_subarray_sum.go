package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	prefixSum := []int{}
	prefixMod := []int{}
	mapMod := map[int]int{}
	prevSum := 0
	for _, val := range nums {
		prevSum += val
		prefixSum = append(prefixSum, prevSum)
		prefixMod = append(prefixMod, prevSum%k)
	}

	fmt.Println(prefixSum)
	fmt.Println(prefixMod)

	for ind, val := range prefixMod {
		if val == 0 && ind > 0 {
			return true
		}

		_, ok := mapMod[val]

		if ok {
			if ind-mapMod[val] > 1 {
				return true
			}
		} else {
			mapMod[val] = ind
		}
	}

	return false
}

func main() {
	fmt.Println("Potato")

	// Input: nums = [23,2,4,6,7], k = 6
	// Output: true
	nums := []int{23, 2, 4, 6, 7}
	k := 6
	fmt.Println(checkSubarraySum(nums, k))

	// Input: nums = [1,5,2,1,5,2,1,3], k = 5
	// Output: true
	nums = []int{1, 5, 2, 1, 5, 2, 1, 3}
	k = 5
	fmt.Println(checkSubarraySum(nums, k))

	// Input: nums = [23,2,6,4,7], k = 13
	// Output: false
	nums = []int{23, 2, 6, 4, 7}
	k = 13
	fmt.Println(checkSubarraySum(nums, k))

	// Input: nums = [23,2,4,6,6], k = 7
	// Output: true
	nums = []int{23, 2, 4, 6, 6}
	k = 7
	fmt.Println(checkSubarraySum(nums, k))

	// Input: nums = [2,4,3], k = 6
	// Output: true
	nums = []int{2, 4, 3}
	k = 6
	fmt.Println(checkSubarraySum(nums, k))

	// Input: nums = [1,0], k = 2
	// Output: false
	nums = []int{1, 0}
	k = 2
	fmt.Println(checkSubarraySum(nums, k))
}
