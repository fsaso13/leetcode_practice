package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	dictMap := map[string]bool{}
	for _, word := range dictionary {
		dictMap[word] = true
	}

	words := strings.Fields(sentence)
	res := []string{}

	for _, word := range words {
		newWord := word
		for i := 0; i < len(word); i++ {
			_, ok := dictMap[word[0:i+1]]
			if ok {
				newWord = word[0 : i+1]
				break
			}
		}
		res = append(res, newWord)
	}

	resString := ""
	for _, word := range res {
		resString += word + " "
	}
	return resString[:len(resString)-1]
}

func main() {
	fmt.Println("Potato")

	// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
	// Output: "the cat was rat by the bat"
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))

	// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
	// Output: "a a b c"
	dictionary = []string{"a", "b", "c"}
	sentence = "aadsfasf absbs bbab cadsfafs"
	fmt.Println(replaceWords(dictionary, sentence))
}
