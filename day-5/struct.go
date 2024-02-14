// A struct (short for structure) is used to create
// a collection of members of different data types,
// into a single variable.
package main
import "fmt"

func main() {

	// STRUCT SYNTAX:
	// type struct_name struct {
	// 	member1 datatype;
	// 	member2 datatype;
	// 	member3 datatype;
	// 	...
	//   }

	// STRUCT EXAMPLE!
	type Person struct {
		name   string
		age    int
		job    string
		salary int
	}

	// CREATING A STRUCT OBJECT!
	obj := Person{name: "John", age: 30, job: "Engineer", salary: 5000}

	// ACCESSING ANY VALUE!
	fmt.Println(obj.age) // OUTPUTS: 30
}