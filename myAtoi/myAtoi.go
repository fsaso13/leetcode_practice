package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	dumbMap := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	byte_str := []rune(s)

	negative := false
	signRead := false
	digitsRead := false
	digits := ""
	for i := 0; i < len(s); i++ {
		char := string(byte_str[i])
		if char == " " && !digitsRead {
			if signRead {
				return 0
			}
			continue
		}

		if char == "+" && !digitsRead {
			if signRead {
				return 0
			}

			signRead = true
			continue
		}

		if char == "-" && !digitsRead {
			if signRead {
				return 0
			}

			negative = true
			signRead = true
			continue
		}

		if unicode.IsDigit(byte_str[i]) {
			digits += char
			if !digitsRead {
				digitsRead = true
			}
		} else {
			if len(digits) == 0 {
				return 0
			}
			break
		}
	}

	// fmt.Println(digits)

	val := 0.
	last := len(digits) - 1
	// fmt.Println(last)
	for s, char := range digits {
		// fmt.Printf("%s -> %f\n", string(char), float64(dumbMap[string(char)]))
		val += math.Pow(10, float64(last-s)) * float64(dumbMap[string(char)])
	}

	if negative {
		val *= -1
	}

	basePow := math.Pow(2, 31)
	min := -basePow
	max := basePow - 1

	if val < min {
		return int(min)
	} else if val > max {
		return int(max)
	} else {
		return int(val)
	}
}

func main() {

	in1 := "42"
	fmt.Println(myAtoi(in1))

	in2 := "-42"
	fmt.Println(myAtoi(in2))

	in3 := "   -42"
	fmt.Println(myAtoi(in3))

	in4 := "4193 with words"
	fmt.Println(myAtoi(in4))

	in5 := "+-12"
	fmt.Println(myAtoi(in5))

	in6 := "+1"
	fmt.Println(myAtoi(in6))

	in7 := "00000-42a1234"
	fmt.Println(myAtoi(in7))
}
