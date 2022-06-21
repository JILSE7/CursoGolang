package main

import (
	"fmt"

	phonePack "test/14-structsPunteros/package"
)

type pc struct {
	ram   int
	disk  int
	Brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.Brand, "pong")
}

func (mylot *pc) duplicateRAM() {

	mylot.ram *= 2
}

func main() {
	a := 50
	b := &a // & accede ala direccion en memoria que tiene la variable a

	fmt.Println(a, b)

	fmt.Println(*b) //accede al valor que tiene esa direccion en memoria

	myPc := pc{ram: 16, disk: 256, Brand: "hp"}

	myPc.ping()

	myPc.duplicateRAM()

	fmt.Println(myPc)

	var myNewPhone phonePack.PhonePublic

	myNewPhone.Brand = "samsung"
	myNewPhone.Memory = 16
	myNewPhone.Ram = 32
	myNewPhone.Price = 16000

	fmt.Println("mynewPhone", myNewPhone)

	//duplicar
	myNewPhone.IncrementMemoryPublic()
	myNewPhone.PrintFeaturesPublic()

}
