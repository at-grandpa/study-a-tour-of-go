package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	zn := 1.0
	zn1 := 0.0
	diff := 10.0
	for i := 0; diff > 1e-15; i++ {
		zn1 = zn - (zn*zn-x)/(2*zn)
		diff = math.Abs(zn1 - zn)
		zn = zn1
	}
	return zn1, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
