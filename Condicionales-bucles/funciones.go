package main

import (
	"fmt"
	"strconv"
)

func sumar(a int, b int) int {
	return a + b
}

func convertirNumero(s string) (int, error) {
	n, err := strconv.Atoi(s)
	return n, err
}

func main() {
	resultado := sumar(3, 5)
	fmt.Println("La suma es:", resultado)

	// * defer fmt.Println("Esta línea se ejecutará al final de la función main")

	num, err := convertirNumero("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Numero:", num)
	}

	saludar := func(nombre string) {
		fmt.Println("Hola,", nombre)
	}

	saludar("Johan")

}
