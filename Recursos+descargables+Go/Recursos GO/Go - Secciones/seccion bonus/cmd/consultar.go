package cmd

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

var consultarCmd = &cobra.Command{
    Use:   "consultar",
    Short: "Mostrar todos los usuarios registrados",
    Run: func(cmd *cobra.Command, args []string) {
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

        fmt.Println("Usuarios registrados:")
        for rows.Next() {
            var id int
            var nombre string
            var edad int
            rows.Scan(&id, &nombre, &edad)
            fmt.Printf("ID: %d | Nombre: %s | Edad: %d\n", id, nombre, edad)
        }
    },
}

func init() {
    rootCmd.AddCommand(consultarCmd)
}
