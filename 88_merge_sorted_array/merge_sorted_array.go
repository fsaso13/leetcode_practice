package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	lhs := 0
	rhs := 0
	resultIndex := 0
	for {
		if lhs < m && rhs < n {
			insertIndex := (resultIndex + m) % (m + n)
			if nums1[lhs] <= nums2[rhs] {
				nums1[insertIndex] = nums1[lhs]
				lhs++
				resultIndex++
			} else {
				nums1[insertIndex] = nums2[rhs]
				rhs++
				resultIndex++
			}
		} else if lhs < m && rhs == n {
			insertIndex := (resultIndex + m) % (m + n)
			nums1[insertIndex] = nums1[lhs]
			lhs++
			resultIndex++
		} else if lhs == m && rhs < n {
			insertIndex := (resultIndex + m) % (m + n)
			nums1[insertIndex] = nums2[rhs]
			rhs++
			resultIndex++
		} else {
			break
		}
	}

	// fmt.Println(nums1)
	aux := nums1[:m]
	// fmt.Println(aux)
	aux2 := nums1[m:]
	// fmt.Println(aux2)
	aux3 := append(aux2, aux...)
	// fmt.Println(aux3)

	copy(nums1, aux3)
}

func main() {
	fmt.Println("potato")

	// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	// Output: [1,2,2,3,5,6]
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
	// Input: nums1 = [1], m = 1, nums2 = [], n = 0
	// Output: [1]
	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)
	// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
	// Output: [1]
	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
	// [4,0,0,0,0,0]
	// 1
	// [1,2,3,5,6]
	// 5
	nums1 = []int{4, 0, 0, 0, 0, 0}
	nums2 = []int{1, 2, 3, 5, 6}
	merge(nums1, 1, nums2, 5)
	fmt.Println(nums1)
}
