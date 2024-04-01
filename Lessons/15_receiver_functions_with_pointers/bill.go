package main

import "fmt"

func main() {

	mybill := newbill("hasky's bill")

	mybill.addItem("light soup", 4.50)
	mybill.addItem("egusi soup", 45.22)
	mybill.addItem("abenkwan", 77.21)
	mybill.addItem("coffee", 7.11)

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
