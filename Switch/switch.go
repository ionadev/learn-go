package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	basic()
	weekday()
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
