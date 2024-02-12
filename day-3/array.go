package main

import (
	"fmt"
)

// !ARRAYS EXAMPLES IN GO!

func main() {

	// ARRAY OF STRINGS! ->
	arr := [3]string{"😢", "🤪", "😆"}
	fmt.Println(arr) // OUTPUTS:[😢 🤪 😆]
	// ARRAY OF NUMBERS! ->
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// ACCESSING ARRAY ELEMENTS BY ITS INDEX!
	fmt.Println(arr[0]) // OUTPUTS: 😢

	// MUTATING ARRAY VALUES!
	arr[0] = "😀"
	fmt.Println(arr)

	// ARRAY INITIALIZATION!
	arr3 := [5]int{}              //not initialized  -> OUTPUTS: [0 0 0 0 0]
	arr4 := [5]int{1, 2}          //partially initialized -> OUTPUTS: [1 2 0 0 0]
	arr5 := [5]int{1, 2, 3, 4, 5} //fully initialized -> OUTPUTS: [1 2 3 4 5]

	fmt.Println(arr3, arr4, arr5)

	// Initialize Only Specific Elements ( CAN BE DONE USING THEIR INDEX)
	arr6 := [5]int{1: 10, 2: 40} // -> OUTPUTS: [0 10 40 0 0]
	fmt.Println(arr6)

	// LENGTH OF ARRAY CAN BE FOUND LIKE USING LEN() METHOD!
	fmt.Println(len(arr)) // -> OUTPUTS: 3

}

// NOTE: The length specifies the number of elements to
// store in the array. In Go, arrays have a fixed length.
// The length of the array is either defined by a number
// or is inferred (means that the compiler decides the
// length of the array, based on the number of values)
