import "fmt"

func main() {
	// Declare a variable
	var x int = 10

	// Declare a pointer
	var pointerToX *int

	// Assign the address of x to the pointer
	pointerToX = &x

	// Print the value of x and the value at the address stored in the pointer
	fmt.Println("Value of x:", x)
	fmt.Println("Value at the address stored in pointerToX:", *pointerToX)

	// Modify the value of x through the pointer
	*pointerToX = 20

	// Print the updated value of x
	fmt.Println("Updated value of x:", x)
}