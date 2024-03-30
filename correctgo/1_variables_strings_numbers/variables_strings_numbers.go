package main

import "fmt"

func main() {
	//strings
	var nameOne string = "Fynn"
	var nameTwo = "Clement"

	fmt.Println(nameOne)
	fmt.Println(nameTwo)

	//change
	nameOne = "Hasky"
	fmt.Println(nameOne)

	//shortcut
	nameFour := "Dawson"
	fmt.Println(nameFour)

	//int
	var ageOne int = 29
	var ageTwo = 45
	ageThree := 66

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits % memory
	var numOne int8 = 25
	var numTwo int8 = -112
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8746326386263
	scoreThree := 3.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
