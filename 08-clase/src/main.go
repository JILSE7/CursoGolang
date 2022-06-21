package main

import "fmt"

func main() {
	//los arrays son estaticos en longitud
	var array [4]int
	array[0] = 1
	array[2] = 45

	fmt.Println(array)
	fmt.Println(len(array), cap(array))
	fmt.Println(array[len(array)-1])

	//los slices son dinamicos en logitud
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Println(slice[0])   //impresion segun la posicion
	fmt.Println(slice[:3])  //imprime hasta el indice 3
	fmt.Println(slice[2:4]) //imprime el indice 2 hasta el indice 4, mas no imprime el 4
	fmt.Println(slice[4:])  //imprime del indice 4 en adelante

	//append aggregar elemento hasta el ultima
	slice = append(slice, 7)
	fmt.Println(slice)

	//agregar otra lista
	newSlice := []int{8, 9, 10}

	slice = append(slice, newSlice...)

	//copiar elementos sin afectar su referencia
	var newboX = make([]int, len(slice))
	copy(newboX, slice)
	slice = append(slice[:1], slice[4:]...)
	fmt.Println("newBox:", newboX, "slice: ", slice)

}
