package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	f1, err := os.Open("enero.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	dfEnero := dataframe.ReadCSV(f1)

	mesEnero := make([]string, dfEnero.Nrow())
	for i := range mesEnero {
		mesEnero[i] = "Enero"
	}
	serieMesEnero := series.New(mesEnero, series.String, "mes")
	dfEnero = dfEnero.Mutate(serieMesEnero)

	f2, err := os.Open("febrero.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	dfFebrero := dataframe.ReadCSV(f2)

	mesFebrero := make([]string, dfFebrero.Nrow())
	for i := range mesFebrero {
		mesFebrero[i] = "Febrero"
	}
	serieMesFebrero := series.New(mesFebrero, series.String, "mes")
	dfFebrero = dfFebrero.Mutate(serieMesFebrero)

	dfUnido := dfEnero.RBind(dfFebrero)

	fmt.Println("\nDataframe combinado:")
	fmt.Println(dfUnido)

	fmt.Println("\nFilas totales: %d\n", dfUnido.Nrow())

}
