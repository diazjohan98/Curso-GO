package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
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
		nil,
		false,
	)
	fmt.Printf("\nModelo de regresion lineal:\n")
	fmt.Printf("\nIngresos = %.2f + %.2f\n * Edad\n", intercepto, pendiente)

	r2 := stat.RSquared(edades, ingresos, nil, intercepto, pendiente)
	fmt.Printf("R2: %.4f\n", r2)

	edadEjemplo := 35.0
	ingresoPredicho := intercepto + pendiente*edadEjemplo
	fmt.Printf("ingreso estimado para edad %.1f: %.2f\n", edadEjemplo, ingresoPredicho)
}
