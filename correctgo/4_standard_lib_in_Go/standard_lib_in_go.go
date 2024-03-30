package main

import (
	"fmt"
	"sort"
	"strings"
)

// imports
func main() {

	greeting := "hello there friends"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "too bad"))

	fmt.Println("original string value =", greeting)

	fmt.Println(strings.ToUpper(greeting))

	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
	ages := []int{34, 56, 53, 5}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 5)
	fmt.Println(index)

	names := []string{"kwabena", "kofi", "kobby", "theresa"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "kobby"))
}
