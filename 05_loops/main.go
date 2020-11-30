package main

import "fmt"

func main() {
	// long version
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Fizz Buzz Test")
	// short version
	for i := 1; i <= 100; i++ {
		switch {
		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}

}
