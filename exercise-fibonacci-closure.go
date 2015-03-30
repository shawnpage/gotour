package main

import "fmt"

// fibonacci is a function that returns a function that returns and int.
func fibonacci() func() int {
	prev := 0
	next := 1

	return func() int {
		next, prev = prev, next + prev
		return next
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
