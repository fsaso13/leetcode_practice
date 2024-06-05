package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	countMap := map[rune]int{}
	pairs := 0

	for _, val := range s {
		_, ok := countMap[val]
		if ok {
			pairs++
			delete(countMap, val)
		} else {
			countMap[val] = 1
		}
	}

	res := 2 * pairs
	if len(countMap) > 0 {
		res++
	}

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: s = "abccccdd"
	// Output: 7
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))

	// Input: s = "a"
	// Output: 1
	s = "a"
	fmt.Println(longestPalindrome(s))
}
