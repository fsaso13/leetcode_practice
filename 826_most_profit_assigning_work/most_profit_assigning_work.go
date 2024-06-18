package main

import (
	"fmt"
	"sort"
)

// type Pair struct {
// 	difficulty int
// 	index      int
// }

func getMax(arr []int) int {
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}

func binarySearch(arr []int, target int) int {
	// fmt.Println("BINARY SEARCH")
	lhs := 0
	rhs := len(arr) - 1
	// fmt.Println(arr)
	for {
		pivot := (rhs + lhs) / 2

		// fmt.Printf("[%d -> %d -> %d]: val => %d\n", lhs, pivot, rhs, arr[pivot])

		if arr[pivot] == target || rhs-lhs <= 1 {
			if arr[pivot] == target {
				for arr[pivot] == target {
					pivot += 1
				}
				pivot -= 1
			}

			// fmt.Println("BINARY SEARCH")
			return pivot
		}

		if arr[pivot] < target {
			lhs = pivot
		} else {
			rhs = pivot
		}
	}
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

	hashMap := map[int][]int{}
	for ind, val := range difficulty {
		hashMap[val] = append(hashMap[val], ind)
	}

	// fmt.Println(difficulty)
	// fmt.Println(profit)

	sort.Slice(difficulty, func(i, j int) bool {
		return difficulty[i] < difficulty[j]
	})

	auxArr := make([]int, len(profit))
	for ind, val := range difficulty {
		ogIndex := hashMap[val][0]
		if len(hashMap[val]) > 1 {
			hashMap[val] = hashMap[val][1:]
		}
		auxArr[ind] = profit[ogIndex]
	}

	profit = auxArr
	// fmt.Println(difficulty)
	// fmt.Println(profit)

	res := 0
	for _, val := range worker {
		if val < difficulty[0] {
			continue
		}

		// fmt.Printf("val: %d\n", val)
		index := len(difficulty) - 1
		if val < difficulty[index] {
			index = binarySearch(difficulty, val)
		}
		// fmt.Printf("index: %d\n", index)
		maxProfit := getMax(profit[:index+1])
		// fmt.Printf("Profit: %d\n", maxProfit)
		res += maxProfit
	}

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
	// Output: 100
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}

	// fmt.Println(binarySearch(difficulty, 7))
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))

	// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
	// Output: 0
	difficulty = []int{85, 47, 57}
	profit = []int{24, 66, 99}
	worker = []int{40, 25, 25}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))

	// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
	// Output: 190
	difficulty = []int{13, 37, 58}
	profit = []int{4, 90, 96}
	worker = []int{34, 73, 45}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))

	// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
	// Output: 1392
	difficulty = []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63}
	profit = []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77}
	worker = []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))

	// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
	// Output: 553
	difficulty = []int{23, 30, 35, 35, 43, 46, 47, 81, 83, 98}
	profit = []int{8, 11, 11, 20, 33, 37, 60, 72, 87, 95}
	worker = []int{95, 46, 47, 97, 11, 35, 99, 56, 41, 92}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))

	// difficulty := []int{23, 30, 35, 35, 43, 46, 47, 81, 83, 98}
	// fmt.Println(binarySearch(difficulty, 35))
}
