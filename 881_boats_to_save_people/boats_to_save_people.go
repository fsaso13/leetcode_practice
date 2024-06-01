package main

import (
	"fmt"
)

func numRescueBoats(people []int, limit int) int {
	valMap := map[int][]int{}

	for ind, val := range people {
		valMap[val] = append(valMap[val], ind)
	}

	// sort.Slice(people, func(i, j int) bool {
	// 	return people[i] > people[j]
	// })

	// fmt.Println(people)

	boats := 0
	// val := 0
	// lhs := 0
	for ind, val := range people {
		fmt.Println(people)
		if val == -1 {
			continue
		}
		people[ind] = -1
		_, ok := valMap[val]
		if ok {
			if len(valMap[val]) == 1 {
				delete(valMap, val)
			} else {
				valMap[val] = valMap[val][1:]
			}
		}

		dif := limit - val
		for dif > 0 {
			_, ok := valMap[dif]
			if ok {
				fmt.Println(valMap[dif])
				newInd := valMap[dif][len(valMap[dif])-1]
				fmt.Println(newInd)

				people[newInd] = -1
				fmt.Println(people)
				if len(valMap[dif]) == 1 {
					delete(valMap, dif)
				} else {
					valMap[dif] = valMap[dif][:len(valMap[dif])-1]
				}
				break
			}

			dif--
		}
		boats++
	}

	return boats
}

func main() {
	fmt.Println("Potato")

	// Input: people = [1,2], limit = 3
	// Output: 1
	people := []int{1, 2}
	limit := 3
	fmt.Println(numRescueBoats(people, limit))

	// Input: people = [3,2,2,1], limit = 3
	// Output: 3
	people = []int{3, 2, 2, 1}
	limit = 3
	fmt.Println(numRescueBoats(people, limit))

	// Input: people = [3,5,3,4], limit = 5
	// Output: 4
	people = []int{3, 5, 3, 4}
	limit = 5
	fmt.Println(numRescueBoats(people, limit))

	// Input: people = 	[2,2], limit = 6
	// Output: 1
	people = []int{2, 2}
	limit = 6
	fmt.Println(numRescueBoats(people, limit))

}
