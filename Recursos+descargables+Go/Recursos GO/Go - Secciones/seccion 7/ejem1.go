// Gorutines: funcion que corre de forma concurrente
// Canales: permite que las gorutines se comuniquen de forma segura.
// Herramientas de sincronización

/*
Concurrencia: multiples tareas al mismo tiempo
C -> plato1, plato2 y plato3
Paralelismo: Ocurre a nivel de cpu
C1 -> plato1
C2 -> plato2
C3 -> plato3
*/

package main

import (
    "fmt"
    "time"
)
func saluda(nombre string) {
    fmt.Printf("Hola, %s!\n", nombre)
}

func main() {

    saluda("Módulo 6")

    go saluda("Goroutine 1")
    go saluda("Goroutine 2")
    go saluda("Goroutine 3")

    time.Sleep(1 * time.Second)

    fmt.Println("¡Todas las goroutines terminaron!")

}
