package main

import "fmt"

// Intergers

func main() {
	fmt.Println("1 + 1 =", 1+1)

	// Floating-point numbers
	fmt.Println("1 + 1 =", 1.0+1.0)

	// Strings
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello," + " World")

	// Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	/* Multiplication of two numbers */

	fmt.Println("The multiplication of 32,132 X 42,452 =", 32132*42452)

	// Find the value of the expression
	fmt.Println((true && false) || (false && true) || !(false && false))

}
