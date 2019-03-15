package main

import "fmt"

// Vertex - The vertex Of Rectangle
type Vertex struct {
	X int
	Y int
}

func basic() {
	fmt.Println(Vertex{1, 3})
}

func medium() {
	v := Vertex{1, 5}
	v.X = 4
	fmt.Println(v.X)
}

func main() {
	basic()
	medium()
}
