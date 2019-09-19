package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int){
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

func singleSum(s []int) int {
	sum := 0
	for _, v:= range s {
		sum += v
	}
	return sum
}

func main(){
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	start := time.Now()
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	v := x+y
	end := time.Now()
	fmt.Println(v, end.Sub(start))

	start2 := time.Now()
	v2 := singleSum(s)
	end2 := time.Now()
	fmt.Println(v2, end2.Sub(start2))

}

