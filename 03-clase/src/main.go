//paquete fmt

package main

import "fmt"

func main() {
	//variables

	hellomessage := "hello"
	world := "world"
	//println impresion en una linea
	println(hellomessage, world)
	println(hellomessage, world)

	//printf
	nombre := "platzi"
	cursos := 500
	//%s string %d numero %v si no sabemos que tipo de dato ira
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	//sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)
	println(message)

	//saber el tipo de dato
	fmt.Printf("hellomessage %T \n", hellomessage)
	isAuth := false
	fmt.Printf("isAUth %T", isAuth)

}
