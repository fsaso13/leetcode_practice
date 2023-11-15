package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	byte_str := []rune(s)
	sum := 0
	start := 0
	end := len(s) - 1
	for {
		if start > end {
			break
		}

		lhs := string(byte_str[start])
		if start == end {
			sum += romanMap[lhs]
			start++
			continue
		}
		rhs := string(byte_str[start+1])

		if lhs == rhs {
			sum += 2 * romanMap[rhs]
			start += 2
		} else {
			lhsNum := romanMap[lhs]
			rhsNum := romanMap[rhs]

			if lhsNum < rhsNum {
				sum += (rhsNum - lhsNum)
				start += 2
			} else {
				sum += lhsNum
				start++
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("IV"))
}
