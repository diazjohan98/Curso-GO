package main

import "fmt"

func main() {
    edad := 12

    if edad >= 18 {
        fmt.Println("Es mayor de edad")
    } else if edad >= 13 {
        fmt.Println("Es adolescente")
    } else {
        fmt.Println("Es niño")
    }

	nota := 10

	if nota >= 70 && nota <= 100 {
    	fmt.Println("Aprobado")
	} else {
    	fmt.Println("Reprobado")
	}


	if edad := 25; edad > 18 {
    	fmt.Println("Mayor de edad")
	}



}
