package main

import (
	"fmt"
)

func appendCharacters(s string, t string) int {

	sIndex := 0
	tIndex := 0
	for {
		if sIndex == len(s) || tIndex == len(t) {
			break
		} else if s[sIndex] == t[tIndex] {
			sIndex++
			tIndex++
		} else {
			sIndex++
		}
	}
	return len(t) - tIndex
}

func main() {
	fmt.Println("Potato")

	// Input: s = "coaching", t = "coding"
	// Output: 4
	s := "coaching"
	t := "coding"
	fmt.Println(appendCharacters(s, t))

	// Input: s = "abcde", t = "a"
	// Output: 0
	s = "abcde"
	t = "a"
	fmt.Println(appendCharacters(s, t))

	// Input: s = "z", t = "abcde"
	// Output: 5
	s = "z"
	t = "abcde"
	fmt.Println(appendCharacters(s, t))
}
