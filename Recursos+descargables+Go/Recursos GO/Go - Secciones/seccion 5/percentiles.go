package main

/*
Un percentil indica el valor debajo del cual se encuentra un porcentaje X de los datos.
Ejemplo:
El percentil 25: 25% de los datos están por debajo.

*/

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"

    "gonum.org/v1/gonum/stat"

    "github.com/go-gota/gota/dataframe"
)

func main() {
    // Abrir y unir CSVs
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

    ingresosStr := df.Col("ingresos").Records()
    var ingresos []float64
    for _, i := range ingresosStr {
        valor, err := strconv.ParseFloat(i, 64)
        if err != nil {
            valor = 0.0
        }
        ingresos = append(ingresos, valor)
    }

    sort.Float64s(ingresos)

    q1 := stat.Quantile(0.25, stat.Empirical, ingresos, nil)
    q2 := stat.Quantile(0.50, stat.Empirical, ingresos, nil) //mediana
    q3 := stat.Quantile(0.75, stat.Empirical, ingresos, nil)

    fmt.Printf("\nPercentiles de ingresos:\n")
    fmt.Printf("Q1 (25%%): %.2f\n", q1)
    fmt.Printf("Q2 (50%%, Mediana): %.2f\n", q2)
    fmt.Printf("Q3 (75%%): %.2f\n", q3)

    iqr := q3 - q1
    fmt.Printf("Rango intercuartílico (IQR): %.2f\n", iqr)

}