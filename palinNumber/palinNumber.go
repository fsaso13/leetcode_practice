package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x < 99 {
		if x%11 == 0 {
			return true
		} else {
			return false
		}
	} else {
		asString := strconv.FormatInt(int64(x), 10)
		last := len(asString) - 1
		byte_str := []rune(asString)
		for i := len(asString) - 1; i > -1; i-- {
			if string(byte_str[i]) != string(byte_str[last-i]) {
				return false
			}
		}
		return true
	}
}

func main() {

	in1 := 121
	fmt.Println(isPalindrome(in1))

	in2 := -121
	fmt.Println(isPalindrome(in2))

	in3 := 10
	fmt.Println(isPalindrome(in3))
}
