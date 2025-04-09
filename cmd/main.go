package main

import (
	"fmt"
	"github.com/whatafunc/Golang_Hexlet_Labs/cmd/for"
	"github.com/whatafunc/Golang_Hexlet_Labs/cmd/testVars"
	"os"
)
func main() {
	//pkgloop.UniqueSortedUserIDs([]int64{5, 3, 4, 7, 8, 9})
	pkgloop.UniqueSortedUserIDs([]int64{5, 5})
	//testvars.PrintCopiedSlice()
	os.Exit(0)
	pkgloop.MapSimple([]string{"While","until","do example"})

	//result := []string{} declined by linter
	//calling a Map function with a slice and a transformation.
	result := pkgloop.Map([]string{"While","until","do"}, func(s string) string{
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	})
	fmt.Println(result)

	testvars.TestVars()
}