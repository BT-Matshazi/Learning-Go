package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("my fav number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println("Value of Pi is: ", math.Pi)

	fmt.Println(add(43, 20))

	fmt.Println(addTwo(77, 55))
	
	fmt.Println(swap("s", "b"))

	fmt.Println(split(12))
}

// type comes after variable
func add( x int, y int) int {
	return x + y
}

// is arguments that follow each other share a type, you cant omit
func addTwo(x , y int) int {
	return x + y
}

// function that return 2 value
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int){
	x = sum * 4 / 2
	y = sum - x
	return
}