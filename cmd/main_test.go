package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/whatafunc/Golang_Hexlet_Labs/cmd/for"
	"strings"
	"testing"
)

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
