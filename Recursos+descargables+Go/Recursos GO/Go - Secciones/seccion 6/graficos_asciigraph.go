package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"

    "github.com/go-gota/gota/dataframe"
    "github.com/guptarohit/asciigraph"
)

type par struct {
    edad    float64
    ingreso float64
}
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

    df := dfEnero.RBind(dfFebrero)

    edadesStr := df.Col("edad").Records()
    ingresosStr := df.Col("ingresos").Records()

    var pares []par
    for i := range edadesStr {
        edad, err1 := strconv.ParseFloat(edadesStr[i], 64)
        ingreso, err2 := strconv.ParseFloat(ingresosStr[i], 64)
        if err1 != nil || err2 != nil {
            continue
        }
        pares = append(pares, par{edad, ingreso})
    }

    sort.Slice(pares, func(i, j int) bool {
        return pares[i].edad < pares[j].edad
    })
    var ingresosOrdenados []float64
    for _, p := range pares {
        ingresosOrdenados = append(ingresosOrdenados, p.ingreso)
    }

    graph := asciigraph.Plot(
        ingresosOrdenados,
        asciigraph.Caption("Ingresos ordenados por edad"),
        asciigraph.Height(10),
    )

    fmt.Println(graph)


}