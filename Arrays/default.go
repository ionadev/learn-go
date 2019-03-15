package main

import "fmt"

func main() {
	// Using var
	var p [2]string
	p[0] = "Hello"
	p[1] = "World"
	fmt.Println(p[0], p[1])

	// Using shorthand operators
	a := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(a)
}
