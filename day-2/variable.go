// package main: Declares that this Go file is part
// of the main package, which is a special package
// indicating it's an executable program.
package main

// The (fmt package) is commonly used for printing
// output to the console or other output streams.
// It provides functions like Println and Printf
// for formatted printing, allowing you to display
// values, variables, and text in a specified format.
import "fmt"

// func main(): Defines the main function, which is
// the entry point of the program. The code inside
// the main function will be executed when the
// program is run
func main() {
	// INT EXAMPLE
	var myInt int = 42
	fmt.Println(myInt) // OUTPUT: 42

	// FLOAT32 EXAMPLE
	var myFloat float32 = 3.14
	fmt.Println(myFloat) // OUTPUT: 3.14

	// STRING EXAMPLE
	var myString string = "Hello, Go!"
	fmt.Println(myString) // OUTPUT: Hello, Go!

	// BOOL EXAMPLE
	var myBool bool = true
	fmt.Println(myBool) // OUTPUT: true

	// SHORT DECLARATION AND ASSIGNMENT!
	// Note: It is not possible to declare a
	// variable using :=, without assigning a value to it.
	name := "John"
	fmt.Println(name) // OUTPUT: John

	// MULTIPLE VARIABLE DECLARATION
	var a, b int = 5, 7
	fmt.Println(a)
	fmt.Println(b)

	// ANOTHER FMT METHOD FOR PRINTING!
	// %v - DEFAULT FORMAT
	number := 42
	fmt.Printf("Default format: %v\n", number)

	// %f - FLOATING-POINT FORMAT
	floatNumber := 3.14
	fmt.Printf("Floating-point format: %f\n", floatNumber)

	// %s - STRING FORMAT
	text := "Hello, Go!"
	fmt.Printf("String format: %s\n", text)

	// %d - INTEGER FORMAT
	age := 25
	fmt.Printf("Integer format: %d\n", age)
}
