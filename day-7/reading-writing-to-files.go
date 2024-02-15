package main

import (
	"fmt"
	"os"
)


func main(){

	// CHECK IF THE FILE EXISTS!
	_,err := os.Stat("file.txt");
	if os.IsNotExist(err){
		fmt.Println("File not exists!")
		return;
	}

	// OPENS THE FILE!
	file,err := os.OpenFile("file.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644);

	if err != nil {
        fmt.Println("Error:", err)
        return
    }
	defer file.Close()

	// ELSE WE WILL WRITE TO THE FILE!
	text := []byte("I am writing in file!");
	_,err = file.Write(text);
   if err != nil {
       fmt.Println("Success!")
   }
}