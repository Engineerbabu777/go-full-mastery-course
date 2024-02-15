package main

import (
	"fmt"
	"strings"
)

func main() {

	// JOINING TWO STRINGS!
	str1 := "Hello"
	str2 := "World"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result) // Output: Hello World

	// SPLITTING A STRING!
	fmt.Println(strings.Split(result, "")) // Output: [Hello World]

	// CHECKING IF A STRINGS CONTAINS ANY SUBSTRING!
	fmt.Println(strings.Contains(result, "World")) // Output: true

	// REPLACING SUBSTRINGS!
	fmt.Println(strings.Replace(result, "World", "Everyone", 1))
	// Output: Hello Everyone
	// 1 -> REPLACE ONLY FIRST EXISTENCE!

	// GETTING THE LENGTH OF A STRING!
	fmt.Println(len(result)) // Output: 11

	// CHANGING CASES!
	fmt.Println(strings.ToLower(result)) // Output: hello world
	fmt.Println(strings.ToUpper(result)) // Output: HELLO WORLD

	
}
