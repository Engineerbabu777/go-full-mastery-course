
package main;

import "fmt";


func main(){

	fmt.Println("POINTERS IN GO💀");

	// POINTERS ARE USED TO ACTUALLY STORED THE ADDRESSES OF VARIABLES!

	// SYNTAX!
	// var ptr *int;

	var ptrVal int = 6 // INTEGER VARIABLE!
	var ptr *int = &ptrVal; // POINTER TO INTEGER!

	fmt.Println("address of ptrVal in memory:",ptr); // OUTPUTS: 0xc000096068

	// GETTING VALUE! // ACTUALLY DEREFERENCING THE POINTER TO GET ITS VAL!
	fmt.Println("value of ptrVal in memory:",*ptr); // OUTPUTS: 6

	// BY THIS WE CAN CHANGE THE VALUES AT THIS LOCATION IN MEMORY!
	*ptr = 10; // CHANGING THE VALUE OF ptrVal!
	fmt.Println("value of ptrVal in memory:",*ptr); // OUTPUTS: 10

	// POINTER TO POINTER!
	var a int = 7;
	fmt.Println(a) // OUTPUTS: 7

	var aPtr *int = &a;
	fmt.Println(aPtr) // OUTPUTS: 0xc000088c68
	fmt.Println(*aPtr) // OUTPUTS: 7

	var ptrToPtr **int = &aPtr;
	fmt.Println(ptrToPtr) // OUTPUTS: 0xc000a77062
	fmt.Println(**ptrToPtr) // OUTPUTS: 7
}