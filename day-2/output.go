package main

import "fmt"

func main() {
	// 1. Println - Print a line with a newline at the end
	fmt.Println("Hello, Go!")

	// 2. Printf - Format and print values
	name := "John"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 3. Print - Print values without formatting
	fmt.Print("This is ")
	fmt.Print("a single line\n")

	// 4. Sprintf - Format and return a string
	formattedString := fmt.Sprintf("The answer is %d\n", 42)
	fmt.Print(formattedString)

	// 5. Fprintln - Format and write to a specified output stream (e.g., a file or network)
	file, _ := // Assume a file is opened for writing
		fmt.Fprintln(file, "Writing to a file")

	// 6. Errorf - Format an error message
	err := fmt.Errorf("An error occurred: %s", "something went wrong")
	fmt.Println(err)
}
