package main

import "fmt"

type ErrNegativeSqrt float64

func Abs(n float64) float64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %.2f\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		z := x
		for Abs(x-z*z) > 0.001 {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	number := 131.0
	sqrt, err := Sqrt(number)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Square root of %.2f is %.2f\n", number, sqrt)
	}
}
