package main

import (
	"fmt"
)

func updateName(x string) {
	x = "eii"
}
func updateNamepointer(x *string) {
	*x = "eii"
}

func main() {

	name := "kwesi"
	updateName(name)
	fmt.Println("name = ", name)

	fmt.Println("memory address of name is", &name)

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("memory address of pointer m is :", &m)
	fmt.Println("value at memory address:", *m)

	fmt.Println(name)
	updateNamepointer(m)
	fmt.Println(name)

}

/*
[--name--][--m--]
[ 0x001  ][0x002]
[--kwesi--][p0x001]
*/
