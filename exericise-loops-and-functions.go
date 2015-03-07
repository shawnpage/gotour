package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

func Sqrt(x float64) float64 {

	z    := float64(x)
	prev := 0.0


	for math.Abs(prev-z) > delta {
		prev = z
		z = z - (z*z - x)/(2*z)
	}
	
	return z
}

func main() {
	const x = 3
//	fmt.Println(Sqrt(x))
	mine, theirs := Sqrt(x), math.Sqrt(x)
    fmt.Println(mine, theirs, mine-theirs)
}
