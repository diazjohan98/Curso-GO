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
/*
    // Leer cabecera
    header, err := reader.Read()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Encabezado:", header)

    // Procesar registros uno por uno
    for {
        record, err := reader.Read()
        if err != nil {
            break
        }

        fmt.Println("Original:", record)
    }
*/

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

    edadDoble := edad * 2

    fmt.Printf("Nombre: %s | Edad original: %d | Edad x2: %d | Ciudad: %s\n",
        nombre, edad, edadDoble, ciudad)
}


}

/*
if edad > 30 {
    fmt.Println(nombre, "tiene más de 30 años")
}
	*/