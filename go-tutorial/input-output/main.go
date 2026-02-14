package main

import "fmt"

func main() {
	fmt.Println("Enter name:")
	var name string
	fmt.Scan(&name) // Taking string as input

	// Scan() argument- Pointer to variable storing the input

	fmt.Printf("Your name is %s\n", name) // String variable output

	var age int
	var height float32

	fmt.Println("Enter age and height(cm):")

	fmt.Scan(&age) // Integer input

	fmt.Scan(&height) // Float input

	// In case of multiple inputs in sequence, values can be entered in a single line
	// with spaces or can be entered in multiple lines

	fmt.Printf("Your age is %d, and height is %.2f cm\n", age, height) // int & float output

	fmt.Println(name, age, height)

	fmt.Printf("Value: %v, Type: %T\n", name, name) // value & type output

	var isT = 9 > 7

	fmt.Printf("Statement is %t\n", isT) // bool output
}
