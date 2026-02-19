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
	fmt.Println(df)

	fmt.Println("Numero de filas:", df.Nrow())
	fmt.Println("Numero de columnas:", df.Ncol())
	fmt.Println("Nombre de columnas:", df.Names())

	/*
	 import (
	 	"string"
	 )



	*/

}
