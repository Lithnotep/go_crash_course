package main

import "fmt"

func main() {
	name := "Max"
	age := 30
	isCool := true
	size := 1.2

	fmt.Println(name, age, isCool, size)

	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

}
