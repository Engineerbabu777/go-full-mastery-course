package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What Programming Language are you currently learning?")
	// READING THE STRING UNTIL NEW LINE!
	input, err := reader.ReadString('\n')
	// HANDLING ERROR!
	if err != nil {
		fmt.Println(err)
	}

	// PRINTING THE STRING FROM THE INPUT!
	fmt.Printf("Currently Learning: %v", input)

}
