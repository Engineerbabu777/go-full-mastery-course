package main

import "fmt"

func main() {
	// Integer Constant
	const daysInWeek int = 7
	fmt.Println("Days in a week:", daysInWeek)

	// Floating-Point Constant
	const pi float64 = 3.14159
	fmt.Println("Value of Pi:", pi)

	// String Constant
	const greeting string = "Hello, Go!"
	fmt.Println("Greeting:", greeting)

	// Boolean Constant
	const isWorkingDay bool = true
	fmt.Println("Is today a working day?", isWorkingDay)

	// Untyped Constant
	const universalAnswer = 42
	fmt.Println("The Answer to the Ultimate 
	Question of Life, the Universe, and 
	Everything:", universalAnswer)
}
