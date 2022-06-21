package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (myCuadrado cuadrado) area() float64 {
	return myCuadrado.base * myCuadrado.base
}

func (t rectangulo) area() float64 {
	return t.base * t.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main() {

	myC := cuadrado{base: 65}
	myT := rectangulo{base: 32, altura: 16}

	calcular(myC)
	calcular(myT)

	//Lista de interfaces
	myInterface := []interface{}{"hola", 12, 4.1}

	fmt.Println(myInterface...)

	slice := []cuadrado{{base: 36}, {base: 78}}

	fmt.Println(slice)
}
