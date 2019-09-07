package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		if z*z - x == 0 {
			return z
		}

		z -= (z*z - x)/(2*z)
	}

	return z
}

func main(){
	fmt.Println(Sqrt(2))
	defer fmt.Println(Sqrt(1))
	defer fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))

}