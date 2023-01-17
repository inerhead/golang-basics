package myPackage

import "fmt"

func ShowZeroValue() {

	var i int
	var f float32
	var p *string
	var s string
	fmt.Printf("\n\n")
	fmt.Println("ZERO VALUES FOR")
	fmt.Printf("%T = %v\n", i, i)
	fmt.Printf("%T = %v\n", f, f)
	fmt.Printf("%T = %v\n", p, p)
	fmt.Printf("%T = %v\n", s, s)
	fmt.Printf("\n\n")

}
