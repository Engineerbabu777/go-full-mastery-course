package main;

// Import necessary packages
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"` 
	Age   int    `json:"age"` 
	Email string `json:"email"` 
}

func main() {
	// Create an instance of the Person struct with some sample data
	person := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}

	// Use json.Marshal to convert the Person struct into a JSON-encoded byte slice
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(jsonData) // OUTPUTS: SLICE CONTAINING ASCII REPRESENTATION

	// Print the JSON-encoded data as a string
	fmt.Println(string(jsonData))
}
