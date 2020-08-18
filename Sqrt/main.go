package main

import "fmt"

func sqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)
		if z*z-x < 0.00001 {
			return z
		}
	}
}
func main() {
	fmt.Println(sqrt(2))
}
