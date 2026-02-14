package main

import (
    "fmt"
    "log"
	"math"
    "os"
    "sort"
    "strconv"

    "gonum.org/v1/gonum/stat"
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

    media := stat.Mean(ingresos, nil)

    mediana := stat.Quantile(0.5, stat.Empirical, ingresos, nil)

	moda := calcularModa(ingresos)

    desviacion := stat.StdDev(ingresos, nil)

    fmt.Printf("\nMétricas de ingresos:\n")
    fmt.Printf("Media: %.2f\n", media)
    fmt.Printf("Mediana: %.2f\n", mediana)
    fmt.Printf("Moda: %.2f\n", moda)
    fmt.Printf("Desviación estándar: %.2f\n", desviacion)


}


func calcularModa(nums []float64) float64 {
    frecuencia := make(map[float64]int)
    for _, n := range nums {
        frecuencia[n]++
    }

    maxFreq := 0
    moda := math.NaN()
    for valor, freq := range frecuencia {
        if freq > maxFreq {
            maxFreq = freq
            moda = valor
        }
    }
    return moda
}
