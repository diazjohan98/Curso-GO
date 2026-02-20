package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/stat"
)

func main() {
	f, err := os.Open("datos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	ciudades := df.Col("ciudad").Records()
	ciudadesSet := make(map[string]bool)

	for _, c := range ciudades {
		ciudadesSet[c] = true
	}

	for ciudad := range ciudadesSet {

		dfCiudad := df.Filter(dataframe.F{
			Colname:    "ciudad",
			Comparator: series.Eq,
			Comparando: ciudad,
		})

		ingresos := dfCiudad.Col("ingresos").Records()
		var ingresosNum []float64

		for _, i := range ingresos {
			valor, err := strconv.ParseFloat(i, 64)
			if err != nil {
				valor = 0.0
			}
			ingresosNum = append(ingresosNum, valor)
		}

		promedio := stat.Mean(ingresosNum, nil)

		fmt.Printf("Ciudad: %-15s | Promedio Ingresos: %.2f\n", ciudad, promedio)
	}
}
