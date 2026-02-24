package main

import (
	"fmt"
	"time"
)

func saludar(nombre string) {
	fmt.Printf("Hola, %s!\n", nombre)
}

func main() {
	saludar("Modulo 6")

	go saludar("Goroutine 1")
	go saludar("Goroutine 2")
	go saludar("Goroutine 3")

	time.Sleep(1 * time.Second)

	fmt.Println("Todas las goroutines han terminado.")
}
