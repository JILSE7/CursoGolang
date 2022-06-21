package main

import (
	"fmt"
	moto "test/13-modules/src"
	hsd "test/test/mypackage"
)

func main() {
	var hfdfd hsd.AvionPublic

	var myMoto moto.MotoPublic

	hfdfd.Brand = "hjdfhjdf"
	myMoto.Brand = "jfdjkdf"
	myMoto.Year = 4545

	fmt.Println(myMoto, hfdfd)

	hsd.PrintMessage("hola mundo desde otra funcion")

}
