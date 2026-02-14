package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
    "github.com/go-gota/gota/series"
)

func main() {
    f, err := os.Open("datos.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)

    fmt.Println("\nDataFrame original:")
    fmt.Println(df)

	dfSelected := df.Select([]string{"nombre", "edad"})
	fmt.Println(dfSelected)

	dfFiltered := df.Filter(
    	dataframe.F{
        	Colname:    "edad",
        	Comparator: series.Greater,
        	Comparando: 30,
    	},
	)
	fmt.Println(dfFiltered)

	dfResult := df.
   		Filter(dataframe.F{
       		Colname:    "edad",
      		Comparator: series.GreaterEq,
    	    Comparando: 30,
    	}).
    	Select([]string{"nombre", "edad"})
	fmt.Println(dfResult)

}

/*
series.Eq (igual)

series.Neq (no igual)

series.Greater (mayor)

series.Less (menor)

series.GreaterEq (mayor o igual)

series.LessEq (menor o igual)

series.In (valor dentro de lista)

series.NotIn (no está en lista)”
*/