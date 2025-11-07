package example

import (
	"fmt"
)

func Example_1_3() {
	var number_a, number_b int
	fmt.Print("Enter number_a: ")
	fmt.Scanln(&number_a)
	fmt.Print("Enter number_b: ")
	fmt.Scanln(&number_b)
	var sum, subtract, multiply, divide int
	sum = number_a + number_b
	subtract = number_a - number_b
	multiply = number_a * number_b
	divide = number_a / number_b
	fmt.Println("sum: ", sum)
	fmt.Println("subtract: ", subtract)
	fmt.Println("multiply: ", multiply)
	fmt.Println("divide: ", divide)
}
