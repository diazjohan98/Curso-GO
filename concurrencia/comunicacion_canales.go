package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-gota/gota/dataframe"
)

type Resultado struct {
	Archivo string
	Filas   int
}

func procesarArchivo(nombre string, wg *sync.WaitGroup, ch chan Resultado) {
	defer wg.Done()

	f, err := os.Open(nombre)
	if err != nil {
		log.Printf("Error abriendo %s: %v\n", nombre, err)
		return

	}

	defer f.Close()

	df := dataframe.ReadCSV(f)
	filas := df.Nrow()

	ch <- Resultado{Archivo: nombre, Filas: filas}
}

func main() {
	archivos := []string{"enero.csv", "febrero.csv", "marzo.csv"}

	var wg sync.WaitGroup

	ch := make(chan Resultado)

	for _, archivo := range archivos {
		wg.Add(1)
		go procesarArchivo(archivo, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for resultado := range ch {
		fmt.Printf("[%s]Filas: %d\n", resultado.Archivo, resultado.Filas)
	}

}
