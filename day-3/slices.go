// In Go, there are several ways to create a slice:

// Using the []datatype{values} FORMAT/SYNTAX
// Create a slice from an array
// Using the make() function

package main

import (
	"fmt"
)

func main() {

	// Using the []datatype{values} FORMAT/SYNTAX
	myslice := []string{"Go", "Slices", "Are", "Powerful"}

	fmt.Println(len(myslice)) // OUTPUTS-> 4
	fmt.Println(cap(myslice)) // OUTPUTS-> 4

	// ACCESSING VALUES!
	fmt.Println(myslice[0]) // OUTPUTS-> Go

	// UPDATE VALUES!
	myslice[2] = "Fun"
	fmt.Println(myslice[2]) // OUTPUTS-> Fun

	// REMOVE ELEMENT/RE-SLICING!!
	// THAT WILL REMOVE THE THIRD VALUE AND OUTPUTS->
	myslice = append(myslice[:3], myslice[4:]...)
	fmt.Println(myslice) // OUTPUTS-> [Go Slices Fun]

	// ADD/APPEND ELEMENTS TO SLICES!
	// SYNTAX:-> slice_name = append(slice_name, element1, element2, ...many more)
	myslice = append(myslice, "To", "Code")
	fmt.Println(myslice) // OUTPUTS-> [Go Slices Are Fun To Code]

	// Create a slice from an array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice2 := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice2)       // OUTPUTS: [12 13]
	fmt.Printf("length = %d\n", len(myslice2))   // OUTPUTS: 2
	fmt.Printf("capacity = %d\n", cap(myslice2)) // OUTPUTS: 4

	// IF THE LENGTH INCREASES THAN THE CAPACITY THEN
	// THE CAPACITY WILL CHANGE TO 2X-OF CAPACITY AS BELOW! ->

	// APPENDING MORE ELEMENTS TO SLICE AND THEN SAVING BACK
	// TO SLICE VARIABLE!

	myslice2 = append(myslice2, 1, 3, 4)
	fmt.Printf("myslice = %v\n", myslice2)       // OUTPUTS: [12 13]
	fmt.Printf("length = %d\n", len(myslice2))   // OUTPUTS: 2
	fmt.Printf("capacity = %d\n", cap(myslice2)) // OUTPUTS: 4

	// Using the make() function
	// SYNTAX:->  slice_name := make([]type, length, capacity)
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

}
