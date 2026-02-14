package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
)

func main() {

    f1, err := os.Open("enero.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f1.Close()

    dfEnero := dataframe.ReadCSV(f1)

    f2, err := os.Open("febrero.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f2.Close()

    dfFebrero := dataframe.ReadCSV(f2)

	dfUnido := dfEnero.RBind(dfFebrero)

    fmt.Println("\nDataFrame combinado:")
    fmt.Println(dfUnido)

    fmt.Printf("\nFilas totales: %d\n", dfUnido.Nrow())

}