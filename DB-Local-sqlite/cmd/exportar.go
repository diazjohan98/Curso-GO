package cmd

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	_ "modernc.org/sqlite"
)

var exportarCmd = &cobra.Command{
	Use:   "exportar",
	Short: "Exportar datos de SQLite a JSON, CSV y Excel",
	Run: func(cmd *cobra.Command, args []string) {
		//conectar con SQLite
		db, err := sql.Open("sqlite", "demo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query(`SELECT id, nombre, edad FROM usuarios`)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		type usuario struct {
			ID     int    `json:"id"`
			Nombre string `json:"nombre"`
			Edad   int    `json:"edad"`
		}

		var usuarios []usuario

		for rows.Next() {
			var u usuario
			rows.Scan(&u.ID, &u.Nombre, &u.Edad)
			usuarios = append(usuarios, u)
		}

		// exportar JSON
		jsonFile, err := os.Create("salida.json")
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()

		enc := json.NewEncoder(jsonFile)
		enc.SetIndent("", "  ")
		enc.Encode(usuarios)
		fmt.Println("Exportado: salida.json")

		// exportar CSV

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

		// Exportar excel
		f := excelize.NewFile()
		sheet := "Usuarios"
		f.NewSheet(sheet)
		f.SetSheetRow(sheet, "A1", &[]interface{}{"ID", "Nombew", "Edad"})

		for i, u := range usuarios {
			row := []interface{}{u.ID, u.Nombre, u.Edad}
			cell := fmt.Sprintf("A%d", i+2)
			f.SetSheetRow(sheet, cell, &row)
		}

		err = f.SaveAs("Salida.xlsx")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Exportado: salida.xlsx")

	},
}

func init() {
	rootCmd.AddCommand(exportarCmd)
}
