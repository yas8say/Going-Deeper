package main

import (
	"fmt"
)

func main() {
	calculate(10, 2, "add")
	calculate(10, 2, "divide")
	calculate(10, 0, "divide")
}

func calculate(a, b int, operation string) {
	defer fmt.Println("Operation Completed.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	case "subtract":
		fmt.Printf("%d + %d = %d\n", a, b, subtract(a, b))
	case "multiply":
		fmt.Printf("%d + %d = %d\n", a, b, multiply(a, b))
	case "divide":
		fmt.Printf("%d + %d = %d\n", a, b, divide(a, b))
	default:
		fmt.Printf("Unknown Operation")
	}
}

func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
