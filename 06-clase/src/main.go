//valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales.
//valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2.
//valor1 < valor2: Retorna TRUE si valor1 es menor que valor2
//valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2
//valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2
//valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2.
//AND 1>0 && 2>0: retorna true si se cumplen las 2 condiciones
//OR 2<0 || 1>0 retorna TRUE porque una de las 2 se cumple
//NOT "!" myBool := true if( !myBool ){}

package main

import (
	"log"
	"strconv"
)

func main() {
	valor := 2
	valor2 := 3
	if valor == 1 && valor2 == 3 {
		println("el valor es uno")
	} else {
		println("No es uno")
	}

	if valor == 2 || valor2 == 3 {
		println("los valores son correctos")
	}

	//convertir texo a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}

	println("Value: ", value)

	println(isPair(3))

	println(isValidPass("papotas"))

}

func isPair(num int) string {
	if num%2 == 0 {
		return "is pair"
	}

	return "isnt pair"
}

func isValidPass(pass string) bool {
	if pass == "papitas" {
		return true
	}
	return false
}
