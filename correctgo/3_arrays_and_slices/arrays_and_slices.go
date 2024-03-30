package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{33, 33, 44}
	var ages = [3]int{20, 5, 44}
	names := [4]string{"repu", "conti", "kat", "ridd"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{33, 55, 77}
	scores[2] = 222
	scores = append(scores, 42)

	fmt.Println(scores, len(scores))

	// //slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	rangeOne = append(rangeOne, "ama")

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
