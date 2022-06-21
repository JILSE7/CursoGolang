package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	var textReverse string
	fdf := strings.ToLower(text)
	for i := len(fdf) - 1; i >= 0; i-- {
		textReverse += string(fdf[i])
	}
	fmt.Println(textReverse)
	if fdf == textReverse {
		return true
	}
	return false
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPal := isPalindromo("Ama")

	if !isPal {
		fmt.Println("is not palindromo")
	} else {
		fmt.Println("is a palindromo")
	}
}
