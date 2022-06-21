package main

import "fmt"

func main() {
	m := make(map[string]int) //llaves-string, valores-enteros is like  {'said': 25}

	m["jose"] = 14
	m["said"] = 25

	fmt.Println(m)

	for i, val := range m {
		fmt.Println(i, val)
	}

	//encontrar un valor
	value, ok := m["josep"]
	//ok indica si el valor fue encontrado, si no lo es, retorna un 0
	//porque le valor zero value de un entero es cero
	if !ok {
		fmt.Println(value)
	}

}
