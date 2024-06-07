package main

import (
	"fmt"
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	// sort.Slice(hand, func(i, j int) bool {
	// 	return hand[i] < hand[j]
	// })

	// fmt.Println(hand)

	freqMap := map[int]int{}

	for _, val := range nums {
		freqMap[val] += 1
	}

	keys := []int{}
	for key := range freqMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for i := 0; i < len(nums)/k; i++ {
		firstKey := -1

		for _, val := range keys {
			_, ok := freqMap[val]
			if ok {
				firstKey = val
				break
			}
		}

		for j := 0; j < k; j++ {
			// fmt.Printf("1st Key: %d\n", firstKey)
			// fmt.Println(freqMap)
			_, ok := freqMap[firstKey+j]

			if ok {
				freqMap[firstKey+j] -= 1

				if freqMap[firstKey+j] == 0 {
					delete(freqMap, firstKey+j)
				}
			} else {
				return false
			}
		}
	}

	// fmt.Println(freqMap)

	return true
}

func isNStraightHand(hand []int, groupSize int) bool {

	if len(hand)%groupSize != 0 {
		return false
	}

	// sort.Slice(hand, func(i, j int) bool {
	// 	return hand[i] < hand[j]
	// })

	// fmt.Println(hand)

	freqMap := map[int]int{}

	for _, val := range hand {
		freqMap[val] += 1
	}

	keys := []int{}
	for key := range freqMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for i := 0; i < len(hand)/groupSize; i++ {
		firstKey := -1

		for _, val := range keys {
			_, ok := freqMap[val]
			if ok {
				firstKey = val
				break
			}
		}

		for j := 0; j < groupSize; j++ {
			// fmt.Printf("1st Key: %d\n", firstKey)
			// fmt.Println(freqMap)
			_, ok := freqMap[firstKey+j]

			if ok {
				freqMap[firstKey+j] -= 1

				if freqMap[firstKey+j] == 0 {
					delete(freqMap, firstKey+j)
				}
			} else {
				return false
			}
		}
	}

	// fmt.Println(freqMap)

	return true
}

func main() {
	fmt.Println("Potato")

	// Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
	// Output: true
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	groupSize := 3
	fmt.Println(isNStraightHand(hand, groupSize))

	// Input: hand = [1,2,3,4,5], groupSize = 4
	// Output: false
	hand = []int{1, 2, 3, 4, 5}
	groupSize = 4
	fmt.Println(isNStraightHand(hand, groupSize))

	// Input: hand = [1,2,3,4], groupSize = 4
	// Output: true
	hand = []int{1, 2, 3, 4}
	groupSize = 4
	fmt.Println(isNStraightHand(hand, groupSize))

	// Input: hand = [1], groupSize = 1
	// Output: true
	hand = []int{1}
	groupSize = 1
	fmt.Println(isNStraightHand(hand, groupSize))

	// Input: hand = [8,10,12], groupSize = 3
	// Output: false
	hand = []int{8, 10, 12}
	groupSize = 3
	fmt.Println(isNStraightHand(hand, groupSize))

	// Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
	// Output: true
	hand = []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	groupSize = 3
	fmt.Println(isNStraightHand(hand, groupSize))
}
