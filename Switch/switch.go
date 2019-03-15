package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	basic()
	weekday()
	nocond()
}

func basic() {

	fmt.Println("GO Works On ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("OS is %s\n", os)
	}
}
func weekday() {
	fmt.Println("When Is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Day After Tomorrow")
	default:
		fmt.Println("Far Away")
	}
}

/* No Condition Switch case */
func nocond() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
