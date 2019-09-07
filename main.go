package main

import (
	"fmt"
	"math/cmplx"
)

var packageLevelVariable string = "packageLevelVariable"

const Pi = 3.1415

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int)(x, y int){
	x = sum * 4/9
	y = sum - x
	return
}

var (
	ToBe bool = false
	MaxInt int64 = 12
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Println("Hello World!")

	a, b := swap("hellow", "word")
	fmt.Println(a, b)

	x, y := split(17)
	fmt.Println(x, y)

	var functionLevelVariable string = "functionLevelVariable"
	fmt.Println(packageLevelVariable, functionLevelVariable)

	var c, python, java = true, false, "no!"	
	fmt.Println(c, python, java)

	k := 3
	fmt.Println(k)

	fmt.Printf("Type : %T Value : %v \n", ToBe, ToBe)
	fmt.Printf("Type : %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v \n", z, z)

}