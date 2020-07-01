package main

import (
	"fmt"
	"math"
	
)

func main() {
	fmt.Println("hello, world")
	fmt.Println("value of pi " , math.Pi)
	fmt.Println( add( 10 , 50 ))
	x,y,total := swap (10 ,20 )
	fmt.Println (x , y , total)
	
}

func add( x int, y int) int {
	var z int
	z = x + y 
	return z
}

func swap( x , y int)  (int, int ,int) {
	var a,b,sum int 
	a = x + 1
	b = y + 1
	sum = a + b 
	count :=10

	for i := 0; i < count; i++ {
		fmt.Println(" value of i" , i )
	}


	return a, b, sum
}


