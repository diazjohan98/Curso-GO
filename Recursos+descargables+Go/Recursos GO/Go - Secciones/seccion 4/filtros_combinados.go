package main

import (
    "fmt"
    "log"
    "os"
	"strings"

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

dfAnd := df.
    Filter(dataframe.F{
        Colname:    "edad",
        Comparator: series.GreaterEq,
        Comparando: 25,
    }).
    Filter(dataframe.F{
        Colname:    "ciudad",
        Comparator: series.Eq,
        Comparando: "Quito",
    })

fmt.Println("\nFiltrado: Edad ≥ 25 Y ciudad = Quito")
fmt.Println(dfAnd)

dfEdad := df.Filter(dataframe.F{
    Colname:    "edad",
    Comparator: series.GreaterEq,
    Comparando: 40,
})

dfCiudad := df.Filter(dataframe.F{
    Colname:    "ciudad",
    Comparator: series.Eq,
    Comparando: "Lima",
})

dfOR := dfEdad.Concat(dfCiudad)

dfOR = distinctRows(dfOR)

fmt.Println("\nFiltrado: Edad ≥ 40 O ciudad = Lima (sin duplicados)")
fmt.Println(dfOR)


dfIn := df.Filter(dataframe.F{
    Colname:    "ciudad",
    Comparator: series.In,
    Comparando: []string{"Quito", "Lima"},
})

fmt.Println("\nFiltrado: ciudad = Quito O Lima (IN)")
fmt.Println(dfIn)


}


func distinctRows(df dataframe.DataFrame) dataframe.DataFrame {
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

