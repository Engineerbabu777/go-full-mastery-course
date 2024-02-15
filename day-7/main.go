// OS PACKAGE WE USED TO COMMUNICATE WITH OPERATING
// SYSTEM AND DO SOME STUFF LIKE READING FILES| CREATING
// FILES|FOLDERS
package main

import (
	"fmt"
	"os"
)

func main() {
	// READING ARGUMENTS FROM THE COMMAND LINE!
	args := os.Args
	fmt.Println("Command-line arguments:", args) // RETURNS SLICE OF ARGUMENTS!

	// CREATING NEW FOLDER!!
	err := os.Mkdir("new_directory", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Directory created successfully.")

	// CREATING AND WRITHING TO FILE!
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Write to the file
	content := []byte("Hello, Go!")
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data written to file.")

	// REMOVING FILE!
	err = os.Remove("file_to_remove.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File removed successfully.")

	// FILE EXISTENCE!
	path := "file_or_directory.txt"
	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		fmt.Println(path, "does not exist.")
	} else {
		fmt.Println(path, "exists.")
	}

	// RENAMING FILES!
	err = os.Rename("old_file.txt", "new_file.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File renamed successfully.")
}
