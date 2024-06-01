package main

import (
	"fmt"
)

func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		val := int(s[i]) - int(s[i+1])
		// fmt.Printf("%d - %d = %d\n", int(s[i]), int(s[i+1]), val)
		if val < 0 {
			val *= -1
		}

		sum += val
	}

	return sum
}

func main() {
	fmt.Println("Potato")

	// Input: s = "hello"
	// Output: 13
	s := "hello"
	fmt.Println(scoreOfString(s))

	// Input: s = "zaz"
	// Output: 50
	s = "zaz"
	fmt.Println(scoreOfString(s))
}
