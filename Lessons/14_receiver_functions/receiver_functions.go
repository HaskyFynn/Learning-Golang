package main

import "fmt"

func main() {

	mybill := newbill("hasky's bill")

	fmt.Println(mybill.format())
}
