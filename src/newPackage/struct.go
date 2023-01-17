package myPackage

import "fmt"

type Car struct {
	Brand string
	year  int
}

func ShowStruct() {
	// Option 1
	myCar := Car{
		Brand: "Chevrolet",
		year:  2014,
	}

	// Option 2
	mycar2 := new(Car)
	mycar2.Brand = "Toyo"
	//mycar2.year = 2022

	fmt.Printf("\n\n")
	fmt.Println("STRUCTS")
	fmt.Printf("Option 1: %+v \n", myCar)
	fmt.Printf("Option 2: %+v \n", *mycar2)
	fmt.Printf("\n\n")
}
