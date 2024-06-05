package main

import (
	"fmt"
)

func commonChars(words []string) []string {
	globalMap := map[string]int{}
	wordMaps := make([]map[string]int, len(words))

	for ind, word := range words {
		wordMaps[ind] = make(map[string]int)
		for _, chr := range word {
			(wordMaps[ind])[string(chr)] += 1
			globalMap[string(chr)] += 1
		}
	}

	// fmt.Println(globalMap)
	// fmt.Println(wordMaps)

	res := []string{}

	for gkey, gval := range globalMap {
		inAll := true
		minFreq := gval
		for _, fmap := range wordMaps {
			_, ok := fmap[gkey]
			if !ok {
				inAll = false
				break
			} else {
				if fmap[gkey] < minFreq {
					minFreq = fmap[gkey]
				}
			}
		}

		if inAll {
			for i := 0; i < minFreq; i++ {
				res = append(res, gkey)
			}
		}
	}

	return res
}

func main() {
	fmt.Println("Potato")

	// Input: words = ["bella","label","roller"]
	// Output: ["e","l","l"]
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(words))

	// Input: words = ["cool","lock","cook"]
	// Output: ["c","o"]
	words = []string{"cool", "lock", "cook"}
	fmt.Println(commonChars(words))
}
