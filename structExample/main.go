package main

import (
	"fmt"
	"sort"
)
type WordFreq struct {
    Word string
    Freq int
}
func updateFreq(wordsSlices []WordFreq, word string) []WordFreq{
	wordFound := false
	for i := range wordsSlices {
        if wordsSlices[i].Word == word {
            wordsSlices[i].Freq ++
			wordFound = true
            break
        }
    }
	if !wordFound {
		wordsSlices = append(wordsSlices, WordFreq{word, 1} )
	}
	return wordsSlices
}
func main(){
	fmt.Println("\033[H\033[2J", "start")
	wordFreqs := []WordFreq{}
	wordFreqs = append(wordFreqs, WordFreq{"apple", 3})
	fmt.Println("wordFreqs = ", wordFreqs)
	wordFreqs = updateFreq(wordFreqs, "banana" )
	wordFreqs = updateFreq(wordFreqs, "huana" )
	wordFreqs = updateFreq(wordFreqs, "huana" )
	wordFreqs = updateFreq(wordFreqs, "huana" )
	fmt.Println("wordFreqs updated = ", wordFreqs)
	sort.Slice(wordFreqs, func(i, j int) bool {
		if wordFreqs[i].Freq == wordFreqs[j].Freq {
	        return wordFreqs[i].Word[i] < wordFreqs[i].Word[j] // Alphabetical if freqs are equal
	    }
		return wordFreqs[i].Freq > wordFreqs[j].Freq
	})
	result := []string{}
	for i := range wordFreqs {
		result = append(result, wordFreqs[i].Word)
    }
	fmt.Println("words  = ", result)

}

