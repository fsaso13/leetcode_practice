package main

import (
	"fmt"
)

// func subarraysDivByK(nums []int, k int) int {
// 	sums := []int{}
// 	modSums := []int{}
// 	mapSums := map[int]int{}
// 	sumVal := 0
// 	res := 0
// 	for _, val := range nums {
// 		sumVal += val
// 		sums = append(sums, sumVal%k)
// 		modSums = append(modSums, sumVal)
// 		if sumVal%k == 0 {
// 			res++
// 		}
// 	}

// 	fmt.Println(modSums)
// 	fmt.Println(sums)
// 	mapFreq := map[int]int{}
// 	for ind, val := range sums {
// 		_, ok := mapSums[val]
// 		if ok {
// 			res++
// 			mapFreq[val]++
// 		} else {
// 			mapSums[val] = ind
// 			mapFreq[val] = 1
// 		}
// 	}

// 	for _, v := range mapFreq {
// 		if v > 1 {
// 			res++
// 		}
// 	}

// 	return res
// }

func subarraysDivByK(nums []int, k int) int {
	prefixMod := 0
	res := 0

	modGroups := make([]int, k)
	modGroups[0] = 1

	for _, num := range nums {
		prefixMod = (prefixMod + num%k + k) % k

		res += modGroups[prefixMod]
		modGroups[prefixMod]++
	}

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: nums = [4,5,0,-2,-3,1], k = 5
	// Output: 7
	nums := []int{4, 5, 0, -2, -3, 1}
	k := 5
	fmt.Println(subarraysDivByK(nums, k))

	// Input: nums = [5], k = 9
	// Output: 0
	nums = []int{5}
	k = 9
	fmt.Println(subarraysDivByK(nums, k))
}
