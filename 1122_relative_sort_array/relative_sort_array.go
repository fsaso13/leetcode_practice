package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	scoreMap := map[int]int{}

	res := []int{}
	res = append(res, arr1...)

	for ind, val := range arr2 {
		scoreMap[val] = ind
	}

	// fmt.Println(arr2)
	// fmt.Println(scoreMap)
	// fmt.Println(res)

	sort.Slice(res, func(i, j int) bool {
		val1, ok1 := scoreMap[res[i]]
		if !ok1 {
			val1 = len(scoreMap)
		}

		val2, ok2 := scoreMap[res[j]]
		if !ok2 {
			val2 = len(scoreMap)
		}
		// fmt.Println("=============================================")
		// fmt.Printf("Map[%d] = %d -> %t\n", res[i], val1, ok1)
		// fmt.Printf("Map[%d] = %d -> %t\n", res[j], val2, ok2)
		// fmt.Println("=============================================")

		if val1 == len(scoreMap) && val2 == len(scoreMap) {
			return res[i] < res[j]
		} else {
			return val1 < val2
		}
	})

	// fmt.Println(res)

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
	// Output: [2,2,2,1,4,3,3,9,6,7,19]
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))

	// Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
	// Output: [22,28,8,6,17,44]
	arr1 = []int{28, 6, 22, 8, 44, 17}
	arr2 = []int{22, 28, 8, 6, 17, 44}
	fmt.Println(relativeSortArray(arr1, arr2))

}
