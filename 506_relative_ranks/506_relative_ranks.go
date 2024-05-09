package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	ogPos := map[int]int{}
	for i, val := range score {
		ogPos[val] = i
	}

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	// fmt.Println(score)

	res := make([]string, len(score))
	for i, val := range score {
		if i == 0 {
			res[ogPos[val]] = "Gold Medal"
		} else if i == 1 {
			res[ogPos[val]] = "Silver Medal"
		} else if i == 2 {
			res[ogPos[val]] = "Bronze Medal"
		} else {
			res[ogPos[val]] = strconv.FormatInt(int64(i+1), 10)
		}
	}

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: score = [5,4,3,2,1]
	// Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
	score := []int{5, 4, 3, 2, 1}
	fmt.Println(findRelativeRanks(score))

	// Input: score = [10,3,8,9,4]
	// Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
	score = []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
}
