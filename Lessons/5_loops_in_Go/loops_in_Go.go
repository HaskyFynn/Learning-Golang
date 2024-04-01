package main

import "fmt"

func main() {

	// loops
	x := 0
	for x < 5 {
		fmt.Println("value of x is =", x)
		x++
	}

	for i := 0; i < 5; i++ {

		fmt.Println("the value of i is =", i)
	}
	names := []string{"dan", "hasky", "emma", "queen"}
	for i := 0; i < len(names); i++ {

		fmt.Println(names[i])
	}
	for index, value := range names {
		fmt.Printf("The position at index %v is %v \n", index, value)
	}
}
