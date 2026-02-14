package main

/*
X: edad
Y: ingresos

Y = a +bX

*/

import (
    "fmt"
    "log"
    "os"
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

    intercepto, pendiente := stat.LinearRegression(
        edades,
        ingresos,
        nil,   // sin pesos
        false, // incluir intercepto
    )


    fmt.Printf("\nModelo de regresión lineal:\n")
    fmt.Printf("Ingresos = %.2f + %.2f * Edad\n", intercepto, pendiente)

    r2 := stat.RSquared(edades, ingresos, nil, intercepto, pendiente)
    fmt.Printf("R²: %.4f\n", r2)

    edadEjemplo := 35.0
    ingresoPredicho := intercepto + pendiente*edadEjemplo
    fmt.Printf("Ingreso estimado para edad %.1f: %.2f\n", edadEjemplo, ingresoPredicho)


}