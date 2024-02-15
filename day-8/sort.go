package main

import (
	"fmt"
	"sort"
)

func main() {

	// USING SORT PACKAGE!

	n := []int{7, 6, 5, 8, 3}
	// THAT WILL SORT THE INT SLICE!
	sort.Ints(n)
	fmt.Println(n)

	fruits := []string{"apple", "orange", "banana", "grape", "kiwi"}
	// THAT WILL SORT THE STRING SLICE IN ASCENDING MANNER!
    sort.Strings(fruits)
    fmt.Println("Sorted strings:", fruits)

	// SORTING IN DESCENDING!
	sort.Sort(sort.Reverse(sort.IntSlice(n)));
	fmt.Println(n)
}
