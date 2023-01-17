package myPackage

import "fmt"

func ShowPointer() {

	myPc := pc{
		ram:   2,
		brand: "DELL",
	}
	myPc.duplicateRam(2)
	fmt.Printf("\n\n")
	fmt.Println("POINTER")
	fmt.Println("PC", myPc)
	fmt.Printf("\n\n")

}

func (myPc pc) String() string {

	data := fmt.Sprintf("HOLA %d", myPc.ram)
	return data

}

func (myPc *pc) duplicateRam(amount int) {
	myPc.ram = myPc.ram * amount
}

type pc struct {
	ram   int
	brand string
}
