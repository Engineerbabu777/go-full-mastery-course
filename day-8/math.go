package main

import (
	"fmt"
	"math"
)

func main() {

	// FINDING SQUARE ROOT!
	num := 25.0
	squareRoot := math.Sqrt(num)
	// OUTPUTS: Square root of 25 is 5
	fmt.Println("Square root of", num, "is", squareRoot)


	// FINDING POWER!
	base := 2.0
    exponent := 3.0
    power := math.Pow(base, exponent)
	// OUTPUTS: 2 raised to the power of 3 is 8
    fmt.Println(base, "raised to the power of", exponent, "is", power)

	// ABSOLUTE VALUE!
	n := -42.5
    absoluteValue := math.Abs(n)
	// OUTPUTS: Absolute value of -42.5 is 42.5
    fmt.Println("Absolute value of", n, "is", absoluteValue)

	// FLOOR && CEIL METHODS!!
	num3 := 3.75
    roundedDown := math.Floor(num3)
    roundedUp := math.Ceil(num3)
    roundedNearest := math.Round(num3)
    fmt.Println("Original:", num3) // Original: 3.75
    fmt.Println("Rounded Down:", roundedDown) //Rounded Down: 3
    fmt.Println("Rounded Up:", roundedUp) // Rounded Up: 4
    fmt.Println("Rounded to Nearest:", roundedNearest)  // Rounded to Nearest: 4
}
