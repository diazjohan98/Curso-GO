package main

/*
Covarianza: mide si dos variables varían juntas (signo positivo → se mueven igual, signo negativo → se mueven opuesto).
Correlación: es la versión normalizada de la covarianza (valor entre -1 y 1).

*/

import (
    "fmt"
    "log"
    "os"
    //"sort"
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

	edadesStr := df.Col("edad").Records()
    ingresosStr := df.Col("ingresos").Records()
    var ingresos []float64
    var edades []float64

	for _, e := range edadesStr {
        valor, err := strconv.ParseFloat(e, 64)
        if err != nil {
            valor = 0.0
        }
        edades = append(edades, valor)
    }
    for _, i := range ingresosStr {
        valor, err := strconv.ParseFloat(i, 64)
        if err != nil {
            valor = 0.0
        }
        ingresos = append(ingresos, valor)
    }

    cov := stat.Covariance(edades, ingresos, nil)
    corr := stat.Correlation(edades, ingresos, nil)

    fmt.Printf("\nCovarianza edad-ingresos: %.2f\n", cov)
    fmt.Printf("Correlación edad-ingresos: %.2f\n", corr)
}