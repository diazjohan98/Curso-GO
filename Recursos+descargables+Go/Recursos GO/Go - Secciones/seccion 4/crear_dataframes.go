package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
)

func main() {
    f, err := os.Open("datos.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)
    fmt.Println(df)

fmt.Println("Número de filas:", df.Nrow())
fmt.Println("Número de columnas:", df.Ncol())
fmt.Println("Nombres de columnas:", df.Names())

/*
import (
    "strings"
)
csvStr := `nombre,edad,ciudad
Jorge,30,Bogotá
Ana,25,Lima
Luis,40,Quito`

df := dataframe.ReadCSV(strings.NewReader(csvStr))
fmt.Println(df)

*/


}
