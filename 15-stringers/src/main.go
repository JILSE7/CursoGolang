package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	Brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB, %d Disco y es una %s", myPc.ram, myPc.disk, myPc.Brand)
}

func main() {
	myPc := pc{ram: 16, disk: 32, Brand: "dell"}

	fmt.Println(myPc.String())

}
