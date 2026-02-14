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
    f, err := os.Open("datos.csv") // usa tu archivo actualizado
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)

    fmt.Println("\nDataFrame original:")
    fmt.Println(df)

    fmt.Println("Tipos detectados:")
    fmt.Println(df.Types())

edades := df.Col("edad").Records()
var edadesLimpias []int

for _, e := range edades {
    if e == "" {
        edadesLimpias = append(edadesLimpias, 0)
    } else {
        edad, err := strconv.Atoi(e)
        if err != nil {
            edad = 0
        }
        edadesLimpias = append(edadesLimpias, edad)
    }
}

serieEdad := series.Ints(edadesLimpias)
serieEdad.Name = "edad_limpia"
df = df.Mutate(serieEdad)

fmt.Println("\nDataFrame con edad limpia:")
fmt.Println(df)

ingresos := df.Col("ingresos").Records()
var ingresosNum []float64

for _, i := range ingresos {
    if i == "" {
        ingresosNum = append(ingresosNum, 0.0)
    } else {
        valor, err := strconv.ParseFloat(i, 64)
        if err != nil {
            valor = 0.0
        }
        ingresosNum = append(ingresosNum, valor)
    }
}

serieIngresos := series.Floats(ingresosNum)
serieIngresos.Name = "ingresos_num"

df = df.Mutate(serieIngresos)

fmt.Println("\nDataFrame con ingresos convertidos a numérico:")
fmt.Println(df)


}