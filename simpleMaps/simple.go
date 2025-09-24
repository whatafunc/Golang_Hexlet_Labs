package main

import (
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func CountWords(text string) map[string]int {
	res := make(map[string]int)
	text = strings.ToLower(text)
	if text == "" {
		return res
	}

	//trimmed := strings.Trim(text, "!?,.;:")
	//words := strings.Split(trimmed, " ")

	//fmt.Println("words: ", words)

	words := strings.Fields(text)
	for _, word := range words {
		//words[i] = strings.Trim(word, "!?,.;:")
		word = strings.Trim(word, "!?,.;:")
		res[word]++
	}

	// for _, word := range words {

	// 	res[word]++
	// }

	return res
}
