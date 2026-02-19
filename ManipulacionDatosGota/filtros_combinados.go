package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"

	"strings"
)

func main() {
	f, err := os.Open("datos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	fmt.Println("\nDataframe original:")
	fmt.Println(df)

	dfAnd := df.
		Filter(
			dataframe.F{
				Colname:    "edad",
				Comparator: series.GreaterEq,
				Comparando: 25,
			}).
		Filter(
			dataframe.F{
				Colname:    "ciudad",
				Comparator: series.Eq,
				Comparando: "Tulua",
			})

	fmt.Println("\nFiltrado: Edad ≥ 25 y ciudad = Tulua ")
	fmt.Println(dfAnd)

	dfEdad := df.Filter(dataframe.F{
		Colname:    "edad",
		Comparator: series.GreaterEq,
		Comparando: 40,
	})

	dfCiudad := df.Filter(dataframe.F{
		Colname:    "ciudad",
		Comparator: series.Eq,
		Comparando: "Barranquilla",
	})

	dfOR := dfEdad.Concat(dfCiudad)

	dfOR = distincRows(dfOR)

	fmt.Println("\nFiltrado: Edad ≥ 40 o ciudad = Barranquilla ")
	fmt.Println(dfOR)

}

func distincRows(df dataframe.DataFrame) dataframe.DataFrame {
	seen := make(map[string]bool)
	uniqueRows := []map[string]interface{}{}

	for _, row := range df.Maps() {
		var keyParts []string
		for _, colName := range df.Names() {
			keyParts = append(keyParts, fmt.Sprint(row[colName]))
		}
		key := strings.Join(keyParts, "|")

		if !seen[key] {
			seen[key] = true
			uniqueRows = append(uniqueRows, row)
		}
	}
	return dataframe.LoadMaps(uniqueRows)
}
