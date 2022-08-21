package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		delta := (z*z - x) / (2 * z)

		if math.Abs(z-delta-x) < 0.000001 {
			return z
		}

		z -= delta
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
