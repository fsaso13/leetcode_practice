package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	expected := []int{}
	expected = append(expected, heights...)

	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})

	// fmt.Println(heights)
	// fmt.Println(expected)

	res := 0
	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("Potato")

	// Input: heights = [1,1,4,2,1,3]
	// Output: 3
	heights := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(heights))

	// Input: heights = [5,1,2,3,4]
	// Output: 5
	heights = []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(heights))

	// Input: heights = [1,2,3,4,5]
	// Output: 0
	heights = []int{1, 2, 3, 4, 5}
	fmt.Println(heightChecker(heights))
}
