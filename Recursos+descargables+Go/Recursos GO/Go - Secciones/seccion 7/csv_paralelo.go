package main

import (
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/go-gota/gota/dataframe"
)

func procesarArchivo(nombre string, wg *sync.WaitGroup) {
    defer wg.Done() // Marcar como hecho cuando termine

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
    // Crear lista de archivos
    archivos := []string{"enero.csv", "febrero.csv", "marzo.csv"}

    // Crear WaitGroup
    var wg sync.WaitGroup

    // Lanzar una goroutine por archivo
    for _, archivo := range archivos {
        wg.Add(1) // Aumentar contador
        go procesarArchivo(archivo, &wg)
    }

    // Esperar a que todas terminen
    wg.Wait()

    fmt.Println("Todos los archivos procesados.")
}

/*
WaitGroup: asegura que main no termine hasta que todas acaben
wg.Add(1):  suma 1 tarea pendiente
defer wg.Done() marca la tarea como completada
go procesarArchivo(...): lanza la tarea concurrente.
*/
