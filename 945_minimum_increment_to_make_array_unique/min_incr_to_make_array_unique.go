package main

import (
	"fmt"
)

// func minIncrementForUnique(nums []int) int {
// 	freqMap := map[int]bool{}
// 	res := 0

// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

// 	for _, val := range nums {
// 		_, ok := freqMap[val]

// 		if ok {
// 			newVal := val + 1
// 			for {
// 				res++
// 				_, ok := freqMap[newVal]
// 				if ok {
// 					newVal++
// 				} else {
// 					freqMap[newVal] = true
// 					break
// 				}
// 			}
// 		} else {
// 			freqMap[val] = true
// 		}
// 	}
// 	return res
// }

func minIncrementForUnique(nums []int) int {
	max := 0
	for _, val := range nums {
		if val > max {
			max = val
		}
	}

	freqArr := make([]int, len(nums)+max)

	for _, val := range nums {
		freqArr[val]++
	}

	res := 0
	for ind, val := range freqArr {
		if val > 1 {
			dupes := val - 1
			freqArr[ind+1] += dupes
			freqArr[ind] = 1
			res += dupes
		}
		fmt.Println(freqArr)
	}

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: nums = [1,2,2]
	// Output: 1
	nums := []int{1, 2, 2}
	fmt.Println(minIncrementForUnique(nums))

	// Input: nums = [1,2,2]
	// Output: 1
	nums = []int{3, 2, 1, 2, 1, 7}
	fmt.Println(minIncrementForUnique(nums))
}
