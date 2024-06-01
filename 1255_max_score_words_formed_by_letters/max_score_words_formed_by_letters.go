package main

import (
	"fmt"
)

func maxScoreWords(words []string, letters []byte, score []int) int {

	alph := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	scoreMap := map[string]int{}

	for ind, val := range score {
		scoreMap[alph[ind]] = val
	}

	fmt.Println(scoreMap)

	lettersMap := map[string]int{}
	for _, val := range letters {
		lettersMap[string(val)] += 1
	}

	fmt.Println(lettersMap)
	maxScore := 0

	for _, word := range words {
		wordScore := 0
		for _, chr := range word {
			wordScore += scoreMap[string(chr)]
		}

		if wordScore > maxScore {
			maxScore = wordScore
		}
	}

	return maxScore
}

func main() {
	fmt.Println("Potato")

	// Input: words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
	// Output: 23
	words := []string{"dog", "cat", "dad", "good"}
	letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(maxScoreWords(words, letters, score))

}
