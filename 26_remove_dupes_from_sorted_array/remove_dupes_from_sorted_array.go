package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	start := 0
	end := 1
	count := 0

	if len(nums) == 2 {
		if nums[start] != nums[end] {
			return 2
		} else {
			return 1
		}
	}

	for {
		// fmt.Printf("Start: %d\n", start)
		// fmt.Printf("LHS: %d\n", nums[start])
		// fmt.Printf("End: %d\n", end)
		// fmt.Printf("RHS: %d\n", nums[end])
		// fmt.Println(nums)
		if start == len(nums) || end == len(nums) || nums[start] > nums[end] {
			count++
			break
		}

		if nums[start] == nums[end] {
			end++
		} else {
			if end-start > 1 {
				tail := nums[end:]
				// fmt.Printf("Tail: ")
				// fmt.Println(tail)

				head := nums[:start+1]
				// fmt.Printf("Head: ")
				// fmt.Println(head)

				reps := nums[start+1 : end]
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)

				aux := append(tail, reps...)
				// fmt.Println(aux)
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				aux = append(head, aux...)
				// fmt.Println(aux)
				copy(nums, aux)
				fmt.Println(nums)
			}
			count++
			start++
			end = start + 1
		}
	}

	return count
}

func main() {
	fmt.Println("potato")

	// Input: nums = [1,1,2]
	// Output: 2, nums = [1,2,_]
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

	// Input: nums = [0,0,1,1,1,2,2,3,3,4]
	// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))

	// Input: nums = [1,1,2]
	// Output: 2, nums = [1,2,_]
	nums = []int{1, 1, 1}
	fmt.Println(removeDuplicates(nums))

	// Input: nums = [1,2]
	// Output: 2, nums = [1,2,_]
	nums = []int{1, 2}
	fmt.Println(removeDuplicates(nums))

	// Input: nums = [1,2,2]
	// Output: 2, nums = [1,2,_]
	nums = []int{1, 2, 2}
	fmt.Println(removeDuplicates(nums))
}
