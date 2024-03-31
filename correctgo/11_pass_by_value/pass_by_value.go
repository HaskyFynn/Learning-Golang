package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 3.55
}

func main() {

	//group A types --> strings, ints, bools, floats, arrays, structs
	name := "kwesi"

	name = updateName(name)
	fmt.Println(name)

	// group B types --> slices, maps, functions
	menu := map[string]float64{

		"pie":       5.33,
		"ice cream": 7.89,
	}
	updateMenu(menu)
	fmt.Println(menu)

}
