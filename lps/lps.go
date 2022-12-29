package main

import (
	"fmt"
)

func isPairPalindrome(s string, start int, end int) bool {
	return s[start] == s[end]
}

func expandPalindrome(s string, start int, end int) string {
	// fmt.Printf("Values from call: [%d,%d]\n", start, end)

	st := start - 1
	en := end + 1

	// fmt.Printf("Starting values: [%d,%d]\n", st, en)

	if st < 0 || en == len(s) {
		return s[st+1 : en]
	}

	for {
		// fmt.Printf("Expanding on [%d,%d]: %s\n", st, en, s[st:en+1])
		if isPairPalindrome(s, st, en) {
			// fmt.Printf("Palindrome on [%d,%d]: %s\n", st, en, s[st:en+1])
			st -= 1
			en += 1
			if st < 0 || en == len(s) {
				return s[st+1 : en]
			}
		} else {
			return s[st+1 : en]
		}
	}
	// expansionRightside := true
	// reachedRight := false
	// reachedLeft := false

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

	start := 0
	end := 1

	odd := false

	res := s[0:1]

	for {
		// fmt.Println("longest")
		// fmt.Println(end)
		if end == len(s) {
			return res
		}

		if isPairPalindrome(s, start, end) {
			candidate := expandPalindrome(s, start, end)
			// fmt.Printf("Candidate: %s\n", candidate)
			if len(candidate) > len(res) {
				res = candidate
			}

		}

		if odd {
			start++
			odd = !odd
		} else {
			end++
			odd = !odd
		}

	}

}

func main() {
	in1 := "babad"
	fmt.Println(longestPalindrome(in1))

	in2 := "cbbd"
	fmt.Println(longestPalindrome(in2))

	in3 := "ac"
	fmt.Println(longestPalindrome(in3))

	in4 := "ccc"
	fmt.Println(longestPalindrome(in4))

	in5 := "aaaa"
	fmt.Println(longestPalindrome(in5))

	in6 := "ccd"
	fmt.Println(longestPalindrome(in6))

	in7 := "abcba"
	fmt.Println(longestPalindrome(in7))

}
