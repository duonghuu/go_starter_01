package main

import (
	"fmt"

	"github.com/duonghuu/go_starter_01/utils"
)

func main() {
	// num1 := 4
	// result := utils.Factorial(num1)
	// fmt.Printf("%d\n", result)

	// kiem tra so nguyen to
	num1 := 4
	result := utils.PrimeNumber(num1)
	fmt.Println(result)
}
