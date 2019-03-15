package main

import (
	"fmt"
	"runtime"
)

func main() {
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
