/*
Debug
Info
Warn
Error
Fatal

Configuración de un nivel:

logger := logrus.New()

// Salida en JSON estructurado
logger.SetFormatter(&logrus.JSONFormatter{})

// Nivel global: solo muestra >= Info
logger.SetLevel(logrus.InfoLevel)


*/
package cmd

import (
    "database/sql"
    "os"

    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

var nombreLogrus string
var edadLogrus int

var logrusCmd = &cobra.Command{
    Use:   "logrus",
    Short: "Insertar usuario y registrar logs estructurados con niveles",
    Run: func(cmd *cobra.Command, args []string) {
        logger := logrus.New()

        // Salida en JSON
        logger.SetFormatter(&logrus.JSONFormatter{})

        // Nivel mínimo: Info
        logger.SetLevel(logrus.InfoLevel)

        // También guardar en archivo
        file, err := os.OpenFile("logrus_output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err == nil {
            logger.SetOutput(file)
        } else {
            logger.Warn("No se pudo abrir archivo, usando stdout")
        }

        // Mensaje Debug (solo se muestra si SetLevel es DebugLevel)
        logger.Debug("Este es un mensaje de depuración (Debug)")

        // Conexión a DB
        db, err := sql.Open("sqlite3", "./demo.db")
        if err != nil {
            logger.WithError(err).Fatal("Error abriendo base de datos")
        }
        defer db.Close()

        // Validar datos
        if nombreLogrus == "" {
            logger.Error("Validación fallida: nombre vacío")
            logger.Fatal("Abortando por falta de nombre")
        }
        if edadLogrus <= 0 {
            logger.WithFields(logrus.Fields{
                "edad": edadLogrus,
            }).Error("Edad inválida")
            logger.Fatal("Abortando por edad inválida")
        }

        logger.WithFields(logrus.Fields{
            "nombre": nombreLogrus,
            "edad":   edadLogrus,
        }).Info("Validación superada")

        // Insertar usuario
        _, err = db.Exec(`INSERT INTO usuarios (nombre, edad) VALUES (?, ?)`, nombreLogrus, edadLogrus)
        if err != nil {
            logger.WithError(err).Fatal("Error insertando usuario")
        }

        logger.WithFields(logrus.Fields{
            "nombre": nombreLogrus,
            "edad":   edadLogrus,
        }).Info("Usuario insertado correctamente")
    },
}

func init() {
    rootCmd.AddCommand(logrusCmd)

    logrusCmd.Flags().StringVarP(&nombreLogrus, "nombre", "n", "", "Nombre del usuario")
    logrusCmd.Flags().IntVarP(&edadLogrus, "edad", "e", 0, "Edad del usuario")
}
