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

}
