package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	str_len := len(s)

	first_char := false
	rhs := str_len - 1
	// lhs := 0
	for i := str_len - 1; i > -1; i-- {
		if !first_char {
			if s[i] == ' ' {
				continue
			} else {
				first_char = true
				rhs = i
			}
		} else {
			if s[i] == ' ' {
				return rhs - i
			} else {
				continue
			}
		}
	}

	return rhs + 1
}

func main() {
	fmt.Println("Potato")

	// Input: s = "Hello World"
	// Output: 5
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))

	// Input: s = "   fly me   to   the moon  "
	// Output: 4
	s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))

	// Input: s = "luffy is still joyboy"
	// Output: 6
	s = "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))

	// Input: s = "a "
	// Output: 1
	s = "a "
	fmt.Println(lengthOfLastWord(s))
}
