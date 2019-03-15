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

func pointer() {
	v := Vertex{1, 3}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func main() {
	basic()
	medium()
	pointer()
}
