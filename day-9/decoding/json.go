package main

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
	// JSON-formatted string representing a person
	jsonStr := `{"name":"Jane Doe","age":25,"email":"jane@example.com"}`

	var person Person

	// Use json.Unmarshal to decode the JSON string into the person struct
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		
		fmt.Println("Error:", err)
		return
	}

	// Print the decoded Person struct
	fmt.Printf("Decoded Person: %+v\n", person)
}
