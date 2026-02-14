package main

import (
    "fmt"
    "sync"
)

func main() {
    var contador int
    var wg sync.WaitGroup
    var mu sync.Mutex // Mutex

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            mu.Lock()         // Bloquea acceso exclusivo
            contador++        // Sección crítica protegida
            mu.Unlock()       // Libera el lock
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Contador final:", contador)
}
