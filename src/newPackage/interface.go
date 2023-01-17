package myPackage

import "fmt"

type figure2D interface {
	area() float32
}

type cuadrado struct {
	base float32
}

type rectangulo struct {
	base   float32
	altura float32
}

func (cua cuadrado) area() float32 {
	return cua.base * 2
}

func (cua rectangulo) area() float32 {
	return cua.base * cua.altura
}

func ShowCalculate(fig figure2D) {
	fmt.Println(fig.area())
}

func inititeInterface() {
	cua := cuadrado{base: 4}
	rect := rectangulo{base: 5, altura: 3}
	inte := []figure2D{cua, rect}
	for _, vec := range inte {
		ShowCalculate(vec)
	}
}
