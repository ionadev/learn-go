package main

import "fmt"

// Vertex - The vertex Of Rectangle
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 3}  //Normal
	v2 = Vertex{X: 4}  // Y is Zero
	v3 = Vertex{}      // X & Y Is Zero
	p  = &Vertex{1, 2} // Type *Vertex
)

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

func advpoint() {
	fmt.Println(v1, p, v2, v3)
}

func main() {
	basic()
	medium()
	pointer()
	advpoint()
}
