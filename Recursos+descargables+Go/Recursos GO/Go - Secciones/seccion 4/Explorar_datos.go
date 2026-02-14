package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
)

func main() {
    f, err := os.Open("datos.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    df := dataframe.ReadCSV(f)
 //   fmt.Println(df)

//	fmt.Println("Número de filas:", df.Nrow())
//	fmt.Println("Número de columnas:", df.Ncol())
//	fmt.Println("Nombres de columnas:", df.Names())
//	fmt.Println(df.Subset(rangeInts(0, 5)))

//    fmt.Println(df.Subset([]int{0, 1, 2}))
//	n := df.Nrow()
//	fmt.Println("\nÚltimas 3 filas:")
//	fmt.Println(df.Subset(rangeInts(n-3, n)))

//	fmt.Println("\nColumna edad:")
//	fmt.Println(df.Col("edad"))

	fmt.Println("\nResumen estadístico:")
	fmt.Println(df.Describe())


}


func rangeInts(from, to int) []int {
    size := to - from
    indexes := make([]int, size)
    for i := 0; i < size; i++ {
        indexes[i] = from + i
    }
    return indexes
}

