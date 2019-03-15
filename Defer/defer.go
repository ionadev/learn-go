package main

/* Whats Defer OwO
*  When The Function Is Fully Executed, Then Only The defer call's arguments works
*  F For Ma So Un-Explainatory English
 */

import "fmt"

func main() {
	basic()
	forefer()
}

func basic() {
	defer fmt.Println("The Function Returned")

	fmt.Println("The Function Exectued!")
}

func forefer() {
	fmt.Println("Initialised")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Closed")
}
