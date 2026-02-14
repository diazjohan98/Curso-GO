package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("datos.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    for _, record := range records {
        fmt.Println(record)
    }
}

/*
for {
    record, err := reader.Read()
    if err != nil {
        break
    }
    fmt.Println(record)
}
*/