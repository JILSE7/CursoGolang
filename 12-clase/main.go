package main

import (
	clase "clase12/src"
	"fmt"
)

func main() {
	var myCar clase.CarPublic

	myCar.Brand = "ford"
	myCar.Year = 2022

	fmt.Println(myCar)

}
