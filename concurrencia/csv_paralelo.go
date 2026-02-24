package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-gota/gota/dataframe"
)

func procesarArchivo(nombre string, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(nombre)
	if err != nil {
		log.Printf("Error al abrir %s: %v\n", nombre, err)
		return

	}

	defer f.Close()

	df := dataframe.ReadCSV(f)
	filas := df.Nrow()

	fmt.Printf("[%s] Filas: %d\n", nombre, filas)
}

func main() {

	archivos := []string{"enero.csv", "febrero.csv", "marzo.csv"}

	var wg sync.WaitGroup

	for _, archivo := range archivos {
		wg.Add(1)
		go procesarArchivo(archivo, &wg)
	}

	wg.Wait()

	fmt.Println("Procesamiento completo.")
}

//* waitGroup asegura que main no termine hasta que todas acaben
//*! wg.Add(1) suma 1 tarea pendiente
//*? defer wg.Done() marca la tarea como completada
//* go procesarArchivo(...) lanza la tarea concurrente
