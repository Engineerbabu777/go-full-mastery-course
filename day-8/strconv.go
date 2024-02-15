package main

import (
	"fmt"
	"strconv"
)

func main() {

	// CONVERTING TO INTEGER!
	str := "123"
	fmt.Printf("Current Type :%T", str) // OUTPUTS: string
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Converted integer:%T", num) // OUTPUTS: int

	// CONVERTING TO FLOAT!
	str2 := "123.45"
	num2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Converted float:%T", num2) // OUTPUTS: float

	// INTEGER TO STRING!
	num3 := 123
	str3 := strconv.Itoa(num3)
	fmt.Printf("Converted string:%T", str3) // OUTPUTS: string
}
