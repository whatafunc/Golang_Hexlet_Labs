package main

import (
	"fmt"
	"slices"
)

var people []Person

func main() {
	fmt.Println("Hello, 世界")
	m := map[int]string{1: "go", 2: "javascript", 3: "python"}
	for key, value := range m {
		fmt.Println("key: ", key, "value: ", value)
	}

	//
	ages := map[string]int{"go": 1, "javascript": 44, "python": 25}
	for name, age := range ages {
		people = append(people, Person{Name: name, Age: age})
	}

	fmt.Println("people: ", people)

	//Теперь people можно отсортировать по возрасту:
	slices.SortFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})

	fmt.Println("people sorted: ", people)

	fmt.Println("Go Go go!", CountWords("Go Go go!"))
	// map[go:3]
	fmt.Println("Hello world hello Go go!", CountWords("Hello world hello Go go!"))
	fmt.Println("Hi, hi. Hi?", CountWords("Hi, hi. Hi?"))

}

//go run main.go simple.go
