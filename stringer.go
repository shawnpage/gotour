package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v --- %v years\n", p.Name, p.Age)
}

func main() {
	a := Person{"Author Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	b := Person{"Tricia McMillan", 36}
	fmt.Println(a,z,b)
}
