package example

import "fmt"

func ExampleScan() {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)
}
