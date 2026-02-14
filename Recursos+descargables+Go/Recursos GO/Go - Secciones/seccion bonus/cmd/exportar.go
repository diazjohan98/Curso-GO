package cmd

import (
    "database/sql"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
    "github.com/xuri/excelize/v2"
)

var exportarCmd = &cobra.Command{
    Use:   "exportar",
    Short: "Exportar datos de SQLite a JSON, CSV y Excel",
    Run: func(cmd *cobra.Command, args []string) {
        // Conectar SQLite
        db, err := sql.Open("sqlite3", "./demo.db")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()

        rows, err := db.Query(`SELECT id, nombre, edad FROM usuarios`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        type Usuario struct {
            ID     int    `json:"id"`
            Nombre string `json:"nombre"`
            Edad   int    `json:"edad"`
        }

        var usuarios []Usuario

        for rows.Next() {
            var u Usuario
            rows.Scan(&u.ID, &u.Nombre, &u.Edad)
            usuarios = append(usuarios, u)
        }

        // Exportar JSON
        jsonFile, err := os.Create("salida.json")
        if err != nil {
            log.Fatal(err)
        }
        defer jsonFile.Close()

        enc := json.NewEncoder(jsonFile)
        enc.SetIndent("", "  ")
        enc.Encode(usuarios)
        fmt.Println("Exportado: salida.json")

        // Exportar CSV
        csvFile, err := os.Create("salida.csv")
        if err != nil {
            log.Fatal(err)
        }
        defer csvFile.Close()

        writer := csv.NewWriter(csvFile)
        writer.Write([]string{"id", "nombre", "edad"})
        for _, u := range usuarios {
            writer.Write([]string{
                fmt.Sprintf("%d", u.ID),
                u.Nombre,
                fmt.Sprintf("%d", u.Edad),
            })
        }
        writer.Flush()
        fmt.Println("Exportado: salida.csv")

        // Exportar Excel
        f := excelize.NewFile()
        sheet := "Usuarios"
        f.NewSheet(sheet)
        f.SetSheetRow(sheet, "A1", &[]interface{}{"ID", "Nombre", "Edad"})

        for i, u := range usuarios {
            row := []interface{}{u.ID, u.Nombre, u.Edad}
            cell := fmt.Sprintf("A%d", i+2)
            f.SetSheetRow(sheet, cell, &row)
        }

        err = f.SaveAs("salida.xlsx")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Exportado: salida.xlsx")
    },
}

func init() {
    rootCmd.AddCommand(exportarCmd)
}
