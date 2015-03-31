package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Hyp() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	x := Vertex{5,6}
	v := &Vertex{3,4}
	fmt.Println(v.Hyp())
	fmt.Println(x.Hyp())
}
