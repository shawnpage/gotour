package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

// func Sqrt(x float64) (float64, error) {
// 	return 0, nil
// }
const delta = 1e-6

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(x)
	prev := 0.0

	for math.Abs(prev-z) > delta {
		prev = z
		z = z - (z*z-x)/(2*z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
