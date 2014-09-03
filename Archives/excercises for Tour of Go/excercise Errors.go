package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(f float64) (float64, error) {

	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	var z, testval float64 = 2146218659, 0

	for (z-testval)*(z-testval) > 0.00000000000001 { //avoid using math.Abs()
		testval = z
		z = z - (z*z-f)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
