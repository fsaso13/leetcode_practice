package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	basePow := math.Pow(2, 31)
	min := int64(-basePow)
	max := int64(basePow - 1)

	abs := int64(x)
	negative := false
	if x < 0 {
		negative = true
		abs *= -1
	}

	asString := strconv.FormatInt(abs, 10)

	last := len(asString) - 1
	byte_str := []rune(asString)
	for i := 0; i < len(asString)/2; i++ {
		aux := last - i
		byte_str[i], byte_str[aux] = byte_str[aux], byte_str[i]
	}

	asString = string(byte_str)

	out, err := strconv.ParseInt(asString, 10, 64)

	if err != nil {
		panic("error parsing string to int")
	}

	if negative {
		out *= -1
	}

	if out <= max && out >= min {
		return int(out)
	} else {
		return 0
	}
}

func main() {
	in1 := 123
	fmt.Println(reverse(in1))

	in2 := -123
	fmt.Println(reverse(in2))

	in3 := 120
	fmt.Println(reverse(in3))
}
