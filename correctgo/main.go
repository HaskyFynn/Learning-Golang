package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Good Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}

}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func main() {

	a1 := circleArea(10.5)
	a2 := circleArea(16)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f ", a1)
	fmt.Printf("\ncircle 2 is %0.3f ", a2)

	// sayGreeting("Hasky")
	// sayGreeting("gob3")
	// sayBye("tea")

	// cycleNames([]string{"cloud", "google", "facebook"}, sayGreeting)
	// cycleNames([]string{"cloud", "google", "facebook"}, sayBye)
	// // //  var ages[3] int = [3]int{33, 33, 44}
	// // var ages = [3]int{20, 5, 44}
	// names := [4]string{"yoshi", "tom", "tee", "ridd"}

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // slices
	// var scores = []int{33, 55, 77}
	// scores[2] = 222
	// scores = append(scores, 42)

	// fmt.Println(scores, len(scores))

	// //slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// rangeOne = append(rangeOne, "ama")

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// imports
	// greeting := "hello there friends"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "too bad"))

	// fmt.Println("original string value =", greeting)

	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))
	// ages := []int{34, 56, 53, 5}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 5)
	// fmt.Println(index)

	// names := []string{"kwabena", "kofi", "kobby", "theresa"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "kobby"))

	// loops
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is =", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {

	// 	fmt.Println("the value of i is =", i)
	// }
	// names := []string{"dan", "hasky", "emma", "queen"}
	// for i := 0; i < len(names); i++ {

	// 	fmt.Println(names[i])
	// // }
	// for index, value := range names {
	// 	fmt.Printf("The position at index %v is %v \n", index, value)
	// }

	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is greater than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("the value at pos %v is %v \n", index, value)

	// }

}
