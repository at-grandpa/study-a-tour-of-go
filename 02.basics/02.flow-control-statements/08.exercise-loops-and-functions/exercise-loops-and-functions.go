package main

import (
	"fmt"
	"math"
)

// Sqrt func(x float64) float64
func Sqrt(x float64) float64 {
	zn := 1.0
	zn1 := 0.0
	diff := 10.0

	for i := 0; diff > 1e-15; i++ {
		zn1 = zn - (zn*zn-x)/(2*zn)
		diff = math.Abs(zn1 - zn)
		zn = zn1
		fmt.Println(i)
	}
	return zn1
}

func main() {
	fmt.Println(Sqrt(2))
}
