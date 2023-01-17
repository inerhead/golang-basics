package main

import (
	"fmt"
	Wild "gossioj/curso/src/newPackage"
)

func main() {
	fmt.Printf("\n\n")
	fmt.Println("Go Lang Basics")
	Wild.ShowZeroValue()
	Wild.ShowArray()
	Wild.ShowSlices()
	Wild.ShowPalindrome()
	Wild.ShowMap()
	Wild.ShowSliceInterface()
	Wild.ShowStruct()

	var myA Wild.Animal
	myA.Name = "Tokio"
	fmt.Println(myA)

	Wild.ShowPointer()

	Wild.MessGo()
	Wild.Channels()

}
