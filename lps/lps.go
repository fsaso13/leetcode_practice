package main

import (
	"fmt"
)

func isPairPalindrome(s string, start int, end int) bool {
	return s[start] == s[end]
}

func expandPalindrome(s string, start int, end int) string {
	st := start
	en := end + 1

	solSt := st
	solEn := en

	if en > len(s) {
		// fmt.Println("Early return")
		return s[start:en]
	}

	for {
		fmt.Printf("Expanding for [%d , %d]\n", st, en)
		if en >= len(s) {
			fmt.Printf("Returning (1) %s\n", s[solSt:solEn])
			return s[solSt:solEn]
		}

		if isPairPalindrome(s, st, en) {
			solEn = en
			solSt = st
			en++
		} else {
			fmt.Printf("Returning (2) %s\n", s[solSt:solEn])
			return s[solSt:solEn]
		}
	}
}

// Current strat: search the whole string from subtrings of len 2
// then expand from them.
// Potential Opt: if len(s) - windowEnd < currentLen, return, we won't find a longest substr
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if isPairPalindrome(s, 0, 1) {
			return s
		} else {
			return s[0:1]
		}
	}

	longSubstr := s[0:1]
	duo := true

	// Min len for string to be palindrome is 2.
	windowStart := 0
	windowEnd := 1
	for {
		if windowEnd >= len(s) {
			return longSubstr
		}

		// fmt.Printf("Checking %s ...\n", s[windowStart:windowEnd+1])

		if isPairPalindrome(s, windowStart, windowEnd) {
			// fmt.Printf("%s is palindrome, expanding...\n", s[windowStart:windowEnd+1])
			candidatePalindrome := expandPalindrome(s, windowStart, windowEnd)
			if len(longSubstr) < len(candidatePalindrome) {
				longSubstr = candidatePalindrome
			}
		}

		if duo {
			windowEnd++
			duo = !duo
		} else {
			windowStart++
			duo = !duo
		}
	}
}

func main() {
	in1 := "babad"
	in2 := "cbbd"
	in3 := "ac"
	in4 := "ccc"
	in5 := "aaaa"

	fmt.Println(longestPalindrome(in1))
	fmt.Println(longestPalindrome(in2))
	fmt.Println(longestPalindrome(in3))
	fmt.Println(longestPalindrome(in4))
	fmt.Println(longestPalindrome(in5))

}
