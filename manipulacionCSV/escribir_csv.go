package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("datos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	//leer cabecera
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cabecera: ", header)

	// crear archivo salida

	outline, err := os.Create("datos_transformados.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outline.Close()

	writer := csv.NewWriter(outline)
	defer writer.Flush()

	//escribir nueva cabecera
	newReader := append(header, "edad_x2")
	writer.Write(newReader)

	// Procesar y escribir registros
	for {
		reacord, err := reader.Read()
		if err != nil {
			break
		}

		nombre := reacord[0]
		edadStr := reacord[1]
		ciudad := reacord[2]

		edad, err := strconv.Atoi(edadStr)
		if err != nil {
			log.Println("Error al convertir edad:", edadStr)
			continue
		}
		edadDoble := strconv.Itoa(edad * 2)

		newRecord := []string{nombre, edadStr, ciudad, edadDoble}

		err = writer.Write(newRecord)
		if err != nil {
			log.Println("Error al escribir:", err)
		}
	}

	fmt.Println("Archivo transformado guardado como datos_transformados.csv")
}
