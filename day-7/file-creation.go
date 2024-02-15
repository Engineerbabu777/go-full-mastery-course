package main

import (
	"fmt"
	"os"
)

func main() {

	// CREATING NEW FILE USING OS PACKAGE!

	// CHECKING IF FILE ALREADY EXISTS!
	_, err := os.Stat("file.txt");
	// ERROR: SPECIFIC FILE NOT EXISTS IN THE DIRECTORY!
	if !os.IsNotExist(err) {
		fmt.Println("File Already Exists!")
		return;
	}

	// CREATING NEW FILE!
	_,err = os.Create("file.txt");
	if err!=nil {
		fmt.Println("File Creation Error!")
	}else{
		fmt.Println("File has been created!")
	}

}