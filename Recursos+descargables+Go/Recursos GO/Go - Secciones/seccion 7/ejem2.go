package main

import (
    "fmt"
)

func trabajar(result chan string) {
    // Simular trabajo
    result <- "Trabajo completado"
}

func main() {
    // Crear canal
    canal := make(chan string)

    // Lanzar goroutine que enviará resultado
    go trabajar(canal)

    // Recibir resultado
    mensaje := <-canal

    fmt.Println("Canal recibió:", mensaje)
}
