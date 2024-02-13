package main

import "fmt"

func main() {
	// Maps are used to store data values in key:value
	// pairs && A map is an unordered and changeable
	// collection that does not allow duplicates

	// SYNTAX-> : map[KeyType]ValueType{key1:value1, key2:value2,...}

	emojis := map[string]string{"sad": "😞", "happy": "😇", "joker": "🃏"}
	fmt.Println(emojis) // OUTPUTS: [sad: 😞 happy: 😇  joker: 🃏]

	// Create Maps Using make()Function
	var a = make(map[string]string) // The map is empty now
	fmt.Println(a)                  // OUTPUTS: map[]
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a) // OUTPUTS: map[brand:Ford model:Mustang year:1964]
	// Accessing Map Elements
	fmt.Println(a["brand"]) // OUTPUTS: Ford

	// REMOVING MAP ELEMENTS!
	delete(a, "model")
	fmt.Println(a) // OUTPUTS: map[brand:Ford year:1964]

	// UPDATING VALUES!
	a["brand"] = "Ford2"
	fmt.Println(a) // OUTPUTS: map[brand:Ford2 year:1964]

	// CHECK IF ELEMENT/KEY EXISTS!
	value, ok := a["brand"]
	fmt.Println(ok, value) // OUTPUTS: true Ford2

	// Maps are references to hash tables.
	// If two map variables refer to the same hash table, changing
	// the content of one variable affect the content of the other.
	var a2 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a2

	fmt.Println(a2) // OUTPUTS: [brand:Ford model:Mustang year:1964]
	fmt.Println(b)  // OUTPUTS: [brand:Ford model:Mustang year:1964]

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a2) // OUTPUTS: [brand:Ford model:Mustang year:1970]
	fmt.Println(b)  // OUTPUTS: [brand:Ford model:Mustang year:1970]

	// ITERATING OVER THE MAP!
	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}
