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

	//operadores aritmeticos
	x := 10
	y := 50

	//suma
	result := x + y

	fmt.Println(result)

	//result
	result = y - x

	fmt.Println(result)

	//Multiplicacion
	result = x * y
	fmt.Println(result)

	//division
	result = x / y
	fmt.Println(result)

	//residuo
	result = y % x
	println("residuo", result)

	//incrementa
	x++
	y++
	println(x, y)

	//decremental
	x--
	y--
	println(x, y)

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

}
