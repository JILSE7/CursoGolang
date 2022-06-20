package main

import "fmt"

func main() {
	//declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//declaracion de variables enteras
	base := 12          //crea y asigna un valor
	var altura int = 14 //crea, indica el tipo de dato y asigna valor
	var area int        //solo crea una variable con su tipo de dato

	fmt.Println(base)
	fmt.Println(altura)
	fmt.Println(area)

	//zero values
	var a int     //-> variables enteras, se inicializa en 0
	var b float64 //-> variables con decimales, se inicializa en 0
	var c string  //-> variables con cadenas, se inicializa vacio
	var d bool    //-> variables con booleanos, se inicializa en false

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println(areaCuadrado)
}
