package main

import "fmt"

func singleNumber(nums []int) []int {
	reps := map[int]int{}

	for _, val := range nums {
		_, ok := reps[val]
		if ok {
			delete(reps, val)
		} else {
			reps[val] = 1
		}
	}

	res := []int{}
	for key, _ := range reps {
		res = append(res, key)
	}

	return res
}

func main() {
	fmt.Println("potato")

	// Input: nums = [1,2,1,3,2,5]
	// Output: [3,5]
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))

	// Input: nums = [-1,0]
	// Output: [-1,0]
	nums = []int{-1, 0}
	fmt.Println(singleNumber(nums))

	// Input: nums = [0,1]
	// Output: [1,0]
	nums = []int{0, 1}
	fmt.Println(singleNumber(nums))
}
