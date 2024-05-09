package main

import (
	"fmt"
	"sort"
)

// func getMax(arr []int, excluded map[int]bool) (int, int) {
// 	max := -1
// 	ind := 0
// 	for i, val := range arr {
// 		_, ok := excluded[i]
// 		if (i == 0 || val > max) && !ok {
// 			max = val
// 			ind = i
// 		}
// 	}

// 	return ind, max
// }

// func reduceArr(arr *[]int, exc map[int]bool) {
// 	for i, val := range *arr {
// 		_, ok := exc[i]
// 		if ok {
// 			continue
// 		}

// 		if val > 0 {
// 			(*arr)[i]--
// 		}
// 	}
// }

// func getSum(arr []int, excluded map[int]bool) int64 {
// 	sum := int64(0)

// 	for key, _ := range excluded {
// 		sum += int64(arr[key])
// 	}

// 	return sum
// }

// func checkAllOnes(arr []int) bool {

// 	m := map[int]bool{}
// 	for _, val := range arr {
// 		_, ok := m[val]
// 		if !ok {
// 			m[val] = true
// 		}
// 	}

// 	return len(m) == 1 && arr[0] == 1
// }

// func maximumHappinessSum(happiness []int, k int) int64 {

// 	if checkAllOnes(happiness) {
// 		// fmt.Println("early")
// 		return 1
// 	}

// 	excl := map[int]bool{}

// 	for i := 0; i < k; i++ {
// 		ind, _ := getMax(happiness, excl)
// 		// fmt.Println(ind)
// 		excl[ind] = true
// 		// fmt.Println(happiness)
// 		reduceArr(&happiness, excl)
// 		// fmt.Println(excl)
// 		// fmt.Println(happiness)
// 	}

// 	return getSum(happiness, excl)
// }

func maximumHappinessSum(happiness []int, k int) int64 {

	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	sum := 0

	for i := 0; i < k; i++ {
		currHap := happiness[i] - i
		if currHap < 0 {
			currHap = 0
		}

		sum += currHap
	}

	return int64(sum)
}
func main() {
	fmt.Println("Potato")

	// Input: happiness = [1,2,3], k = 2
	// Output: 4
	happiness := []int{1, 2, 3}
	k := 2
	fmt.Println(maximumHappinessSum(happiness, k))

	// Input: happiness = [1,1,1,1], k = 2
	// Output: 1
	happiness = []int{1, 1, 1, 1}
	k = 2
	fmt.Println(maximumHappinessSum(happiness, k))

	// Input: happiness = [2,3,4,5], k = 1
	// Output: 5
	happiness = []int{2, 3, 4, 5}
	k = 1
	fmt.Println(maximumHappinessSum(happiness, k))

	// Input: happiness = [7,50,3], k = 3
	// Output: 57
	happiness = []int{7, 50, 3}
	k = 3
	fmt.Println(maximumHappinessSum(happiness, k))
}
