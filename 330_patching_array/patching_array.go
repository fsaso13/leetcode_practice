package main

import (
	"fmt"
)

func minPatches(nums []int, n int) int {
	preSum := []int{}
	mapFreq := map[int]int{}

	sum := 0
	for _, val := range nums {
		sum += val
		preSum = append(preSum, sum)
		mapFreq[val] = 1
	}

	newNums := []int{}
	newNums = append(newNums, nums...)
	newNums = append(newNums, n)

	fmt.Println(preSum)
	fmt.Println(newNums)
	fmt.Println(mapFreq)

	if n-preSum[len(preSum)-1] == 0 {
		return 0
	}

	rhsSums := len(preSum) - 1
	rhsNums := len(newNums) - 1
	res := 0

	for {
		currNum := newNums[rhsNums] - preSum[rhsSums]
		_, ok := mapFreq[currNum]
		if ok {
			return res
		} else {
			res++

			newNums = append(newNums[:len(newNums)-1], currNum)

		}

	}
	return 0
}

func minPatches2(nums []int, n int) int {
	preSum := []int{}
	mapFreq := map[int]int{}

	// Fill the presum array
	sum := 0
	for _, val := range nums {
		sum += val
		preSum = append(preSum, sum)
		mapFreq[val] = 1
	}

	// Copy of array with all previous numbers
	newNums := []int{}
	newNums = append(newNums, nums...)
	newNums = append(newNums, n)

}

func main() {
	fmt.Println("Potato")

	// Input: nums = [1,3], n = 6
	// Output: 1
	nums := []int{1, 3}
	n := 6
	fmt.Println(minPatches(nums, n))
}
