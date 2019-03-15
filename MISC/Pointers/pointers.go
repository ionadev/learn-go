package main

import "fmt"

func main() {
	//Declare The variables
	i, j := 10, 256
	//Point p variable to address of i
	p := &i
	//Read The Value Of Pointer
	fmt.Println(*p)
	//Assign a Different value to p
	*p = 21
	//Print i's Value To check  If It Has Changed
	fmt.Println(i)

	//Point p to j now
	p = &j
	//Read value of j
	fmt.Println(*p)
	//Change Its Value
	*p = 568
	//Send It To Console
	fmt.Println(j)

}
