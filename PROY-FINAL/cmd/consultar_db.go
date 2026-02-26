package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

var consultasDBCmd = &cobra.Command{
	Use:   "consultar-db",
	Short: "Muestra los registros validos almacenados en la base de datos",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", "demo.db")
		if err != nil {
			log.Fatalf("Error al conectar con la base de datos: %v", err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT nombre, apellido, correo, genero, ip FROM usuarios_limpios")
		if err != nil {
			log.Fatalf("Error al ejecutar la consulta: %v", err)
		}
		defer rows.Close()

		fmt.Println("Registros validos almacenados")

		var nombre, apellido, correo, genero, ip string
		for rows.Next() {
			err := rows.Scan(&nombre, &apellido, &correo, &genero, &ip)
			if err != nil {
				log.Println("Error al leer fila:", err)
				continue
			}
			fmt.Printf("- %s %s | %s | %s | IP: %s\n", nombre, apellido, correo, genero, ip)
		}

		err = rows.Err()
		if err != nil {
			log.Fatalf("Error al recorrer los resultados: %v", err)
		}
	},
}
