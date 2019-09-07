package main

import "fmt"

func main(){
	i, j := 52, 2801

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i, *p)
	fmt.Println(&i, p)

	p = &j
	*p = *p/37
	fmt.Println(j)
}