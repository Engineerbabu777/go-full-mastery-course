// READING AND WRITING AND WRITING IN FILES IN GO!

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// FILE CREATION FUNCTION!
	// fileCreation("dummy2.txt")

	// FILE WRITING!
	// fileWriting("./dummy2.txt","Hello")

	// READING FROM THE FILE!
	fileReading("./dummy2.txt")
}

// FUNCTION THAT WILL CREATE A FILE!
func fileCreation(fileName string) {

	// CREATING A FILE!
	_, err := os.Create(fileName)

	// HANDLING ERROR!
	ErrorHandler(err)

	fmt.Println("File Creation Success!")
}

// FUNCTION THAT WILL WRITE TO FILE!
func fileWriting(fileName string, content string) {
   // OPENING THE FILE IN WRITE MODE!
    file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666);

	ErrorHandler(err);

	l,err := io.WriteString(file,content);

	fmt.Println("The length of the file: ",l);
    // CLOSING THE FILE!
	defer file.Close();
}

// FUNCTION THAT WILL READ FROM THE FILE!
func fileReading(fileName string) {
    
	fileInByte,err := ioutil.ReadFile(fileName);
	ErrorHandler(err);
	fmt.Println("The content of the file is: ",string(fileInByte));

}

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
