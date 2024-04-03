package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	stringsLength := len(s)
	auxMapS1 := map[string]int{}
	auxMapS2 := map[string]int{}
	auxMap := map[string]string{}

	for i := 0; i < stringsLength; i++ {
		st := string(s[i])
		_, ok := auxMapS1[st]
		if !ok {
			auxMapS1[st] = 1
		} else {
			auxMapS1[st]++
		}

		st = string(t[i])
		_, ok = auxMapS2[st]
		if !ok {
			auxMapS2[st] = 1
		} else {
			auxMapS2[st]++
		}

		c1 := string(s[i])
		c2 := string(t[i])
		val, ok := auxMap[c1]

		if ok {
			if val != c2 {
				return false
			}
		} else {
			auxMap[c1] = c2
		}
	}

	// for i := 0; i < stringsLength; i++ {
	// 	st := string(t[i])
	// 	_, ok := auxMapS2[st]
	// 	if !ok {
	// 		auxMapS2[st] = 1
	// 	} else {
	// 		auxMapS2[st]++
	// 	}
	// }

	// fmt.Println((auxMapS1))
	// fmt.Println((auxMapS2))

	// fmt.Println(len(auxMapS1))
	// fmt.Println(len(auxMapS2))
	return len(auxMapS1) == len(auxMapS2)
}

func main() {
	fmt.Println("Potato")

	// Input: s = "egg", t = "add"
	// Output: true
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))

	// Input: s = "foo", t = "bar"
	// Output: false
	s = "foo"
	t = "bar"
	fmt.Println(isIsomorphic(s, t))

	// Input: s = "paper", t = "title"
	// Output: true
	s = "paper"
	t = "title"
	fmt.Println(isIsomorphic(s, t))

	// Input: s = "bbbaaaba", t = "aaabbbba"
	// Output: false
	s = "bbbaaaba"
	t = "aaabbbba"
	fmt.Println(isIsomorphic(s, t))
}
