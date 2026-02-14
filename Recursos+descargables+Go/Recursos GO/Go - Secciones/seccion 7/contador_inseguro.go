/*
Una condición de carrera ocurre cuando varias goroutines acceden y modifican una misma variable al mismo tiempo, sin sincronización.

Esto puede causar resultados inconsistentes, errores aleatorios o datos corruptos.
*/
package main

import (
    "fmt"
    "sync"
)

func main() {
    var contador int
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            contador++ // ❌ posible condición de carrera
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Contador final:", contador)
}
