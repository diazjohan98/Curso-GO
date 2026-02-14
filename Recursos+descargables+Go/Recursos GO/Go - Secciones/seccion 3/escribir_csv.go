package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    // Abrir archivo original
    file, err := os.Open("datos.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    // Leer cabecera
    header, err := reader.Read()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Cabecera:", header)

    // Crear archivo de salida
    outFile, err := os.Create("datos_transformados.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()

    writer := csv.NewWriter(outFile)
    defer writer.Flush()

    // Escribir nueva cabecera
    newHeader := append(header, "edad_x2")
    writer.Write(newHeader)

    // Procesar y escribir registros
    for {
        record, err := reader.Read()
        if err != nil {
            break
        }

        nombre := record[0]
        edadStr := record[1]
        ciudad := record[2]

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
