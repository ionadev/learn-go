package main

import (
	"fmt"       //For Writing To Console
	"math"      //For math operations
	"math/rand" //For Random Numbers
)

/* Prints The Square Root Of Given Number*/
func sqrt(x float64) {
	fmt.Println(math.Sqrt(x))
}

/* Prints a Random Number | But Number is same everytime, you have to seed it. */
func random(x int) {
	fmt.Println("A Random Number is", rand.Intn(x))
}

/* Give The Value Of Pi */
func pi() {
	fmt.Println("The Value Of Pi is", math.Pi)
}

/* Adds 2 integers */
func add(x int, y int) {
	fmt.Printf("The Sum Of %b & %b is %b", x, y, x+y)
}

/* Swaps Two Strings */
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	return
}
