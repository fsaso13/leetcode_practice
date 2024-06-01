package main

import (
	"fmt"
)

func reversePrefix(word string, ch byte) string {

	pos := -1
	for ind, val := range []byte(word) {
		if val == ch {
			pos = ind
			break
		}
	}

	if pos == -1 {
		return word
	}

	// fmt.Printf("%d/2 : %d\n", pos, pos/2)

	w := []byte(word)
	nPos := pos / 2
	if pos%2 != 0 {
		nPos++
	}

	// fmt.Println(nPos)
	for i := 0; i < nPos; i++ {
		w[i], w[pos-i] = w[pos-i], w[i]
	}

	return string(w[:])
}

func main() {
	fmt.Println("Potato")

	// Input: word = "abcdefd", ch = "d"
	// Output: "dcbaefd"
	word := "abcdefd"
	ch := byte('d')
	fmt.Println(reversePrefix(word, ch))

	// Input: word = "xyxzxe", ch = "z"
	// Output: "zxyxxe"
	word = "xyxzxe"
	ch = byte('z')
	fmt.Println(reversePrefix(word, ch))

	// Input: word = "abcd", ch = "z"
	// Output: "abcd"
	word = "abcd"
	ch = byte('z')
	fmt.Println(reversePrefix(word, ch))
}
