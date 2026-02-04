package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var a int = 10
	b := 20
	var c = a + b

	fmt.Println("Sum =", c)

	var n = 45
	fmt.Printf("Decimal: %d\nBinary: %b\nHexa: %X\nGo-format: %#v\n", n, n, n, n)

	s := "Hello"
	fmt.Printf("Value = %v, Type = %T\n", s, s)

	x := 1
	p := &x
	fmt.Printf("Address: %p, Value: %d\n", p, *p)
}

// Single line comment

/*
Multi
Line
Comment
*/
