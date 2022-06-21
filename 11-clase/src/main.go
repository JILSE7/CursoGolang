package main

import "fmt"

//creamos objetos con struct es como si fuera una clase
type car struct {
	brand      string
	year       int
	isVerified bool
}

func main() {

	myCar := car{brand: "ford", year: 56456, isVerified: true}

	fmt.Println(myCar)

	//otra manera de crear objetos a partir del struc
	var otherCar car
	otherCar.year = 2022
	otherCar.brand = "lambo"
	fmt.Println(otherCar)

}
