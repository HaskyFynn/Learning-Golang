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

	sayGreeting("Hasky")
	sayGreeting("gob3")
	sayBye("tea")

	cycleNames([]string{"cloud", "google", "facebook"}, sayGreeting)
	cycleNames([]string{"cloud", "google", "facebook"}, sayBye)
	var ages [3]int = [3]int{33, 33, 44}
	var ages2 = [3]int{20, 5, 44}
	names := [4]string{"yoshi", "tom", "tee", "ridd"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	fmt.Println(ages2, len(ages2))
}
