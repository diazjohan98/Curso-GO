package cmd

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

var sqliteCmd = &cobra.Command{
    Use:   "sqlite",
    Short: "Demo de conexión a SQLite",
    Run: func(cmd *cobra.Command, args []string) {
        // Abrir o crear base de datos local
        db, err := sql.Open("sqlite3", "./demo.db")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()

        // Crear tabla si no existe
        sqlStmt := `
        CREATE TABLE IF NOT EXISTS usuarios (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            nombre TEXT,
            edad INTEGER
        );
        `
        _, err = db.Exec(sqlStmt)
        if err != nil {
            log.Fatal(err)
        }

        // Insertar datos de ejemplo
        _, err = db.Exec(`INSERT INTO usuarios (nombre, edad) VALUES (?, ?)`, "Jorge", 30)
        if err != nil {
            log.Fatal(err)
        }

        // Leer registros
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
    rootCmd.AddCommand(sqliteCmd)
}
