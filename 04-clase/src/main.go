package main

func normalFunction(message string) {
	println("hello", message)
}

//a int, b int, c string
//a, b int , c  string
func tripeArguments(a int, b int, c string) {
	println(a, b, c)
}

func returnValue(a int) int {
	return a * 12
}

//retornar 2 valores enteros
func doubleReturn(a int) (c, b int) {
	return a, a * 12
}

func main() {
	normalFunction("said")

	tripeArguments(1, 2, "jdf")

	println(returnValue(6))

	println(doubleReturn(5))

	value1, value2 := doubleReturn(5)

	println(value1, value2)

	//descartar un valor
	value3, _ := doubleReturn(5)

	println(value3)
}
