package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	fmt.Println("Enter you age: ");

	reader := bufio.NewReader(os.Stdin);

	age,err := reader.ReadString('\n');
	
	if err!=nil{
		fmt.Println("Error: ", err);
	}

	a,_ := strconv.ParseFloat(strings.TrimSpace(age),64);

	fmt.Println("Your age after 6 years: ",a+5);

}