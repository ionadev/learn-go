package main

import "fmt"

//While Loop in Go is used by For
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
