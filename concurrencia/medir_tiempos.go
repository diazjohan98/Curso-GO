package main

import (
	"fmt"
	"sync"
	"time"
)

func tarea(id int) {
	fmt.Printf("Procesando tarea #%d\n", id)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	tareas := []int{1, 2, 3, 4, 5}

	inicioSecuencial := time.Now()
	for _, t := range tareas {
		tarea(t)
	}
	duracionSecuencial := time.Since(inicioSecuencial)
	fmt.Printf("Tiempo secuencial: %v\n\n", duracionSecuencial)

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
