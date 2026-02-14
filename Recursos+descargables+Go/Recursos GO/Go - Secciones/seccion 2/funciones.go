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
    resultado := sumar(5, 3)
    fmt.Println("Resultado:", resultado)

	//defer fmt.Println("Esto se ejecuta al final")

	num, err := convertirNumero("42")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Número:", num)
    }

	saludar := func(nombre string) {
    	fmt.Println("Hola,", nombre)
	}

	saludar("Jorge")
}
