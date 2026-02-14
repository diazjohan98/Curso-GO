package main

import (
    "fmt"
    "log"
    "os"
    "strconv"

    "github.com/go-gota/gota/dataframe"
    "github.com/go-gota/gota/series"
)


func main() {
    f, err := os.Open("datos.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)

    fmt.Println("\nDataFrame original:")
    fmt.Println(df)

	var edadesDoble []int

	edades := df.Col("edad").Records()

    for _, e := range edades {
        edad, err := strconv.Atoi(e)
        if err != nil {
            edad = 0
        }
        edadesDoble = append(edadesDoble, edad*2)
    }

	nuevaColumna := series.Ints(edadesDoble)
    nuevaColumna.Name = "edad_doble"

	dfNuevo := df.Mutate(nuevaColumna)

    fmt.Println("\nDataFrame con columna nueva:")
    fmt.Println(dfNuevo)

}