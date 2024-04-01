package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":            4.99,
		"meatpie":             6.99,
		"ayigbe toffee": 3.35,
	}
	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		44556677: "hasky",
		66554433: "akosua",
		77889900: "chrome",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[44556677])

	phonebook[44556677] = "bowser"
	fmt.Println(phonebook)

}
