package myPackage

import "fmt"

func ShowArray() {
	// option 1
	var array0 [2]int
	array0[0] = 5
	array0[1] = 8

	// option 2
	array := [4]int{3, 4}
	fmt.Printf("\n\n")
	fmt.Println("ARRAY")
	fmt.Printf("Option 1: %+v con length %d and Cap %d\n", array0, len(array0), cap(array0))
	fmt.Printf("Option 2: %+v con length %d and Cap %d\n", array, len(array), cap(array))
	fmt.Printf("\n\n")
}
