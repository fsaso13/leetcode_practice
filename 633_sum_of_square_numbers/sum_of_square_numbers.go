package main

import (
	"fmt"
)

func judgeSquareSum(c int) bool {
	for i := 2; i*i <= c; i++ {
		count := 0
		if c%i == 0 {
			for {
				if c%i != 0 {
					break
				}

				count++
				c /= i
			}
			if i%4 == 3 && count%2 != 0 {
				return false
			}
		}
	}
	return c%4 != 3
}

func main() {
	fmt.Println("Potato")
}
