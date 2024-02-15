package main

import (
	"fmt"
	"os"
)

func main() {

	// RENAMING TO FILE!
	err := os.Rename("file.txt", "b.txt")

	if err != nil {
		fmt.Println("File has been renamed!")
	} else {
		fmt.Println("File not exists!")
	}

	// REMOVING TO A FILE!
	err = os.Remove("b.txt")
	if err != nil {
		fmt.Println("File has been removed!")
	} else {
		fmt.Println(err)
	}
}
