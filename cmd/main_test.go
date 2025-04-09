package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/whatafunc/Golang_Hexlet_Labs/cmd/for"
	"strings"
	"testing"
)

func TestUniqueSortedUserIDs(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int64{}, pkgloop.UniqueSortedUserIDs([]int64{}))
	a.Equal([]int64{10}, pkgloop.UniqueSortedUserIDs([]int64{10}))
	a.Equal([]int64{55}, pkgloop.UniqueSortedUserIDs([]int64{55, 55}))
	a.Equal([]int64{22, 33, 55}, pkgloop.UniqueSortedUserIDs([]int64{55, 55, 33, 22}))
	a.Equal([]int64{2, 33, 55, 88, 103}, pkgloop.UniqueSortedUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}

func TestMapSimple(t *testing.T) {
	a := assert.New(t)
	a.Equal([]string([]string(nil)), pkgloop.MapSimple([]string{}))
}

func TestMap(t *testing.T) {
	a := assert.New(t)
	testMap(a, []string{"John", "Peter", "Fedor"}, []string{"john", "peter", "fedor"}, func(s string) string {
		return strings.Title(s) //nolint
	})
	testMap(a, []string{"hello", "world"}, []string{"HELLO", "WORLD"}, func(s string) string {
		return strings.ToLower(s)
	})
}

func testMap(a *assert.Assertions, expected, input []string, mapFunc func(s string) string) {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	a.Equal(expected, pkgloop.Map(input, mapFunc))
	a.Equal(inputCopy, input) // check that the initial slice hasn't been modified.
}
