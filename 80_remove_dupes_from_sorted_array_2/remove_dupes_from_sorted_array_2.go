package main

import (
	"fmt"
)

func removeDuplicatesV2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	start := 0
	end := 1
	count := 0
	currentLast := len(nums) - 1

	if len(nums) == 2 {
		if nums[start] != nums[end] {
			return 2
		} else {
			return 1
		}
	}

	for {
		fmt.Printf("Start: %d\n", start)
		fmt.Printf("LHS: %d\n", nums[start])
		fmt.Printf("End: %d\n", end)
		fmt.Printf("RHS: %d\n", nums[end])
		// fmt.Println(nums)
		if start == len(nums) || end == len(nums) || nums[start] > nums[end] {
			// count++
			break
		}

		if nums[start] == nums[end] {
			end++
		} else {
			if end-start > 1 {
				tail := nums[end:]
				// fmt.Printf("Tail: ")
				// fmt.Println(tail)

				head := nums[:start+2]
				// fmt.Printf("Head: ")
				// fmt.Println(head)

				reps := nums[start+2 : end]
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				currentLast -= len(reps)

				aux := append(tail, reps...)
				// fmt.Println(aux)
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				aux = append(head, aux...)
				// fmt.Println(aux)
				copy(nums, aux)
				fmt.Println(nums)
				count += 2
				fmt.Println(count)
				if end > currentLast {
					break
				}
			} else {
				count++
			}
			start++
			end = start + 1
		}
	}

	return count
}

func removeDuplicatesV3(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	if len(nums) == 2 {
		return 2
	}

	start := 0
	end := 1
	count := 0
	currentLast := len(nums) - 1

	for {
		// fmt.Println("--------------------")
		// fmt.Printf("Start: %d\n", start)
		// fmt.Printf("LHS: %d\n", nums[start])
		// fmt.Printf("End: %d\n", end)
		// fmt.Printf("RHS: %d\n", nums[end])
		// fmt.Printf("Current Last: %d\n", currentLast)
		// fmt.Println(nums)
		// fmt.Println("--------------------")
		// if start == len(nums) || end == len(nums) || nums[start] > nums[end] {
		if end > currentLast {
			if end-start == 1 {
				count++
			} else {
				tail := nums[end:]
				// fmt.Printf("Tail: ")
				// fmt.Println(tail)

				head := nums[:start+2]
				// fmt.Printf("Head: ")
				// fmt.Println(head)

				reps := nums[start+2 : end]
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				// currentLast -= len(reps)

				aux := append(tail, reps...)
				// fmt.Println(aux)
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				aux = append(head, aux...)
				// fmt.Println(aux)
				copy(nums, aux)
				// fmt.Println(nums)
				count += 2
				// fmt.Println(count)
			}
			break
		}

		if nums[start] == nums[end] {
			end++
		} else {
			if end-start > 1 {
				tail := nums[end:]
				// fmt.Printf("Tail: ")
				// fmt.Println(tail)

				head := nums[:start+2]
				// fmt.Printf("Head: ")
				// fmt.Println(head)

				reps := nums[start+2 : end]
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				currentLast -= len(reps)

				aux := append(tail, reps...)
				// fmt.Println(aux)
				// fmt.Printf("Reps: ")
				// fmt.Println(reps)
				aux = append(head, aux...)
				// fmt.Println(aux)
				copy(nums, aux)
				// fmt.Println(nums)
				count += 2
				// fmt.Println(count)
				start += 2
			} else {
				start++
				count++
			}
			// start = end
			end = start + 1
		}
	}

	// fmt.Println(currentLast)
	// fmt.Println(count)

	// fmt.Println(end)
	return count
	// return end
	// return currentLast + 1
}

func main() {
	fmt.Println("potato")

	// Input: nums = [1,1,1,2,2,3]
	// Output: 5, nums = [1,1,2,2,3,_]
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicatesV3(nums))

	// Input: nums = [0,0,1,1,1,1,2,3,3]
	// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicatesV3(nums))

	// Input: nums = [1,1,1]
	// Output: 2, nums = [1,1,_]
	nums = []int{1, 1, 1}
	fmt.Println(removeDuplicatesV3(nums))

	// Input: nums = [1,1,1,2,2,2,3,3]
	// Output: 6, nums = [1,1,2,2,3,3,_,_]
	nums = []int{1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Println(removeDuplicatesV3(nums))
}
