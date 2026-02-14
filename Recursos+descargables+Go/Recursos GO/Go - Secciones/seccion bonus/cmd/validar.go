package cmd

import (
    "database/sql"
    "errors"
    "fmt"
    "log"

    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

var nombre string
var edad int

func ValidarUsuario(nombre string, edad int) error {
    if nombre == "" {
        return errors.New("el nombre es obligatorio")
    }
    if edad <= 0 {
        return errors.New("la edad debe ser mayor a cero")
    }
    return nil
}

var validarCmd = &cobra.Command{
    Use:   "validar",
    Short: "Insertar usuario con validación de reglas de negocio",
    Run: func(cmd *cobra.Command, args []string) {
        // Usar la función que ahora es testeable
        err := ValidarUsuario(nombre, edad)
        if err != nil {
            log.Fatal("❌ Error:", err)
        }

        db, err := sql.Open("sqlite3", "./demo.db")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()

        _, err = db.Exec(`CREATE TABLE IF NOT EXISTS usuarios (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            nombre TEXT,
            edad INTEGER
        );`)
        if err != nil {
            log.Fatal(err)
        }

        _, err = db.Exec(`INSERT INTO usuarios (nombre, edad) VALUES (?, ?)`, nombre, edad)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Usuario insertado: %s (%d años)\n", nombre, edad)
    },
}

func init() {
    rootCmd.AddCommand(validarCmd)

    validarCmd.Flags().StringVarP(&nombre, "nombre", "n", "", "Nombre del usuario")
    validarCmd.Flags().IntVarP(&edad, "edad", "e", 0, "Edad del usuario")
}
