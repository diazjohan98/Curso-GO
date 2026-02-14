// Que es una herramiento CLI
/*
Command Line Interface

- Se ejecuta desde terminal
- Recibe parámetros
- Ejecuta tareas automáticas
*/

package main

import (
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
)

func main() {
    // Definir bandera
    archivo := flag.String("archivo", "", "Ruta del archivo CSV a procesar")
    flag.Parse()

    // Validar
    if *archivo == "" {
        log.Fatal("Debes proporcionar la ruta con -archivo=nombre.csv")
    }

    f, err := os.Open(*archivo)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)
    fmt.Printf("[%s] Filas: %d\n", *archivo, df.Nrow())
}
