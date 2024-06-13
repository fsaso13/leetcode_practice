package main

import (
	"fmt"
)

// func minMovesToSeat(seats []int, students []int) int {
// 	seatMap := map[int]int{}

// 	for _, val := range seats {
// 		seatMap[val] += 1
// 	}

// 	fmt.Println(seatMap)
// 	res := 0

// 	for _, val := range students {
// 		_, ok := seatMap[val]

// 		fmt.Printf("Current Student: %d\n", val)

// 		if ok {
// 			// res++
// 			seatMap[val]--
// 			if seatMap[val] == 0 {
// 				delete(seatMap, val)
// 			}
// 		} else {
// 			index := 1
// 			for {
// 				res++
// 				fmt.Printf("# movements: %d\n", res)
// 				fmt.Printf("%d vs %d\n", val+index, val-index)

// 				_, ok := seatMap[val-index]
// 				if ok {
// 					// res++
// 					seatMap[val-index]--
// 					if seatMap[val-index] == 0 {
// 						delete(seatMap, val-index)
// 					}
// 					break
// 				}

//					_, ok = seatMap[val+index]
//					if ok {
//						// res++
//						seatMap[val+index]--
//						if seatMap[val+index] == 0 {
//							delete(seatMap, val+index)
//						}
//						break
//					}
//					index++
//					// res++
//				}
//			}
//		}
//		return res
//	}
func findMax(arr []int) int {
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}

func minMovesToSeat(seats []int, students []int) int {
	m1 := findMax(seats)
	m2 := findMax(students)

	max := m1
	if m2 > m1 {
		max = m2
	}

	diffs := make([]int, max)

	for _, val := range seats {
		diffs[val-1]++
	}

	for _, val := range students {
		diffs[val-1]--
	}

	moves := 0
	unmatched := 0

	for _, val := range diffs {
		absUnmatched := unmatched
		if absUnmatched < 0 {
			absUnmatched *= -1
		}

		moves += absUnmatched
		unmatched += val
	}

	return moves
}

func main() {
	fmt.Println("Potato")

	// Input: seats = [3,1,5], students = [2,7,4]
	// Output: 4
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	fmt.Println(minMovesToSeat(seats, students))

	// Input: seats = [4,1,5,9], students = [1,3,2,6]
	// Output: 7
	seats = []int{4, 1, 5, 9}
	students = []int{1, 3, 2, 6}
	fmt.Println(minMovesToSeat(seats, students))

	// Input: seats = [2,2,6,6], students = [1,3,2,6]
	// Output: 4
	seats = []int{2, 2, 6, 6}
	students = []int{1, 3, 2, 6}
	fmt.Println(minMovesToSeat(seats, students))
}
