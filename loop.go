package main

import (
    "fmt"
	"math"
)

func Sqrt(x float64) (lastResult float64) {
    z := float64(1)
    lastResult = x
    for {
        
        z = z - ((math.Pow(z, 2) - x)/(2*z))
        fmt.Printf("z: %f\n", z)

 
        if lastResult - z < 0.00000001 { break }
        fmt.Printf("lastResult: %f\n", lastResult)
        lastResult = z
    } 
    return
}

func main() {
    var num = float64(120)
    fmt.Printf("The calculated square root of %+v is %+v\n", num, Sqrt(num))
    fmt.Printf("The library square root of %+v is %+v\n", num, math.Sqrt(num))

}
