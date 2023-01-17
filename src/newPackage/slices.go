package myPackage

import "fmt"

func ShowSlices() {

	// Option 1
	s1 := []string{
		"Animal",
		"People",
	}

	// Option 2
	var s2 []string
	s2 = append(s2, "Nuevo")
	fmt.Printf("\n\n")
	fmt.Println("SLICES")
	fmt.Printf("Option 1: %+v \n", s1)
	fmt.Printf("Option 2: %+v \n", s2)
	someLoop(s1)
	fmt.Printf("\n\n")
}

func someLoop(slice []string) {
	fmt.Println(slice[1:2])
	newOne := append(slice, "88")
	last := append([]string{"2"}, newOne...)
	fmt.Println(newOne, last)

	for _, v := range last {
		fmt.Println(v)
	}
}
