/*
make(chan T) 

<- canal 

canal <- x
*/

package main

import (
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/go-gota/gota/dataframe"
)

// Estructura para enviar resultado
type Resultado struct {
    Archivo string
    Filas   int
}

func procesarArchivo(nombre string, wg *sync.WaitGroup, ch chan Resultado) {
    defer wg.Done() // Marcar como completada

    f, err := os.Open(nombre)
    if err != nil {
        log.Printf("Error abriendo %s: %v\n", nombre, err)
        return
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)
    filas := df.Nrow()

    // Enviar resultado al canal
    ch <- Resultado{Archivo: nombre, Filas: filas}
}

func main() {
    archivos := []string{"enero.csv", "febrero.csv", "marzo.csv"}

    var wg sync.WaitGroup

    // Crear canal para recibir resultados
    ch := make(chan Resultado)

    // Lanzar goroutines
    for _, archivo := range archivos {
        wg.Add(1)
        go procesarArchivo(archivo, &wg, ch)
    }

    // Goroutine extra: cerrar canal cuando todas terminen
    go func() {
        wg.Wait()
        close(ch)
    }()

    // Recibir resultados del canal
    for resultado := range ch {
        fmt.Printf("[%s] Filas: %d\n", resultado.Archivo, resultado.Filas)
    }

    fmt.Println("Todos los archivos procesados y resultados recibidos.")
}
