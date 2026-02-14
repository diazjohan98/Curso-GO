/*
time.Now()
time.Since()
*/

package main

import (
    "fmt"
    "sync"
    "time"
)

// Simula procesar un archivo, tarda 500ms
func tarea(id int) {
    fmt.Printf("Procesando tarea #%d\n", id)
    time.Sleep(500 * time.Millisecond)
}

func main() {
    tareas := []int{1, 2, 3, 4, 5}

    // Ejecutar secuencialmente
    inicioSecuencial := time.Now()
    for _, t := range tareas {
        tarea(t)
    }
    duracionSecuencial := time.Since(inicioSecuencial)
    fmt.Printf("Tiempo secuencial: %v\n\n", duracionSecuencial)

    // Ejecutar concurrentemente
    var wg sync.WaitGroup
    inicioConcurrente := time.Now()

    for _, t := range tareas {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            tarea(id)
        }(t)
    }

    wg.Wait()
    duracionConcurrente := time.Since(inicioConcurrente)
    fmt.Printf("Tiempo concurrente: %v\n", duracionConcurrente)
}
