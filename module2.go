package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return (sqrt(-x) + "i")
	}

	return fmt.Sprint(math.Sqrt(x))
}


func main(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	here := 1

	for here < 1000 {
		here += here
	}
	fmt.Println(here)

	fmt.Println(sqrt(2), sqrt(-4))
}