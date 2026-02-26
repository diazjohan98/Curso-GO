// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/go-gota/gota/dataframe"
// )

// func main() {

// 	//* definir bandera
// 	archivo := flag.String("archivo", "", "Ruta del archivo a procesar")
// 	flag.Parse()

// 	//* validar
// 	if *archivo == "" {
// 		log.Fatal("Debe proporcionar la ruta con -archivo=nombre.csv")
// 	}

// 	f, err := os.Open(*archivo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	df := dataframe.ReadCSV(f)
// 	fmt.Printf("[%s] filas: %d\n", *archivo, df.Nrow())

// }