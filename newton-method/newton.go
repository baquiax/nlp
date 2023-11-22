package main

import (
	"fmt"
	"math"
)

func main() {

	var e float64 = 0.0000001

	f := func(x float64) float64 {
		return x - math.Cos(x)
	}

	fPrime := func(x float64) float64 {
		return 1 + math.Sin(x)
	}

	x := 1.0
	k := 0
	fmt.Printf("k=%d, x_k=%.8f\n", k, x)

	for {
		k++

		x1 := x - (f(x) / fPrime(x))

		if math.Abs(x1-x) < e {
			break
		}

		x = x1

		fmt.Printf("k=%d, x_k=%.8f\n", k, x)
	}

}
