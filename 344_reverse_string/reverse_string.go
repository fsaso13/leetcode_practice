package main

import (
	"fmt"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	fmt.Println("Potato")

	// Input: s = ["h","e","l","l","o"]
	// Output: ["o","l","l","e","h"]
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s[:]))

	// Input: s = ["H","a","n","n","a","h"]
	// Output: ["h","a","n","n","a","H"]
	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	fmt.Println(string(s[:]))
}
