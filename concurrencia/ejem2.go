package main

import "fmt"

func trabajar(result chan string) {
	result <- "Trabajo completado"
}

func main() {
	canal := make(chan string)

	go trabajar(canal)

	mensaje := <-canal

	fmt.Println("Cabal recibió:", mensaje)
}
