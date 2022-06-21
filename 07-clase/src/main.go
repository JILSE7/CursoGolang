package main

func main() {

	switch module := 5 % 2; module {
	case 0:
		println("es par")

	default:
		println("Es impar")

	}

	//sin condicion
	value := 200

	switch {
	case value > 100:
		println("es mayor a 100")
	case value < 0:
		println("es menor a 0")
	default:
		println("default")
	}

	//defer - permite ejecutar algo antes de que la funcion termine
	defer println("hola")
	println("mundo")

	//continue
	for i := 0; i < 10; i++ {
		println(i)

		if i == 2 {
			println("es 2")
			continue
		}
		//break
		if i == 8 {
			break
		}
	}

}
