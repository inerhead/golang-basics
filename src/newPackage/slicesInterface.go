package myPackage

import "fmt"

func ShowSliceInterface() {
	inter := []interface{}{2, "33", true}
	fmt.Printf("\n\n")
	fmt.Println("SLICE INTERFACE")
	for in, v := range inter {
		fmt.Println(v, in)
	}

	fmt.Printf("\n\n")
}
