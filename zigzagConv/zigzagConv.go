package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	steps := numRows - 2
	totalMov := numRows + steps

	currentLevel := 0

	answer := ""

	col := 0
	currentLetter := currentLevel
	it := 0
	for {
		if it == len(s) {
			return answer
		}

		if currentLetter >= len(s) {
			col = 0
			currentLevel++
			currentLetter = currentLevel
		}

		if currentLevel == numRows {
			return answer
		}

		answer += s[currentLetter : currentLetter+1]

		if !(currentLevel == 0 || currentLevel == numRows-1) {
			currentLetter = (numRows-currentLevel-1)*2 + currentLetter
			if currentLetter >= len(s) {
				continue
			}

			answer += s[currentLetter : currentLetter+1]
		}

		col++
		currentLetter = currentLevel + col*totalMov
		it++
	}

}

func main() {

	in1 := "PAYPALISHIRING"
	r1 := 3
	fmt.Println("A: PAHNAPLSIIGYIR")
	fmt.Printf("R: %s\n", convert(in1, r1))

	in2 := "PAYPALISHIRING"
	r2 := 4
	fmt.Println("A: PINALSIGYAHRPI")
	fmt.Printf("R: %s\n", convert(in2, r2))
}
