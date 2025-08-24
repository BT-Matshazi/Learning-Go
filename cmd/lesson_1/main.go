package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(printMe("Bekithemba", "Matshazi"))

	var answer, err = cal(20, 2)
	fmt.Printf("answer is %v", answer)
	fmt.Printf("error is %v", err)
}

func printMe(firstName string, lastName string) string{
	var result string = firstName + " " + lastName

	return result
}

func cal( numOne int, numTwo int) (int, error) {

	var err error
	if numTwo == 0 {
		err = errors.New("Cant divide by zero")
		return 0, err
	}

	var result int = numOne / numTwo

	return result, err
}