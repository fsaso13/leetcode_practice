package main

import (
	"fmt"
)

func getMax(bloomDay []int) int {
	maxVal := 0

	for _, val := range bloomDay {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func minDays(bloomDay []int, m int, k int) int {

	if len(bloomDay) < k*m {
		return -1
	}

	n := len(bloomDay)

	bloomedFlowers := make([]int, n)

}

func main() {
	fmt.Println("Potato")

	// Input: bloomDay = [1,10,3,10,2], m = 3, k = 1
	// Output: 3
	bloomDay := []int{1, 10, 3, 10, 2}
	m := 3
	k := 1
	fmt.Println(minDays(bloomDay, m, k))
}
