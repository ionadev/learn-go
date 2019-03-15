package main

import "fmt"

func main() {
	/* Whats Defer OwO
	*  When The Function Is Fully Executed, Then Only The defer call's arguments works
	*  F For Ma So Un-Explainatory English
	 */
	defer fmt.Println("The Function Returned")

	fmt.Println("The Function Exectued!")
}
