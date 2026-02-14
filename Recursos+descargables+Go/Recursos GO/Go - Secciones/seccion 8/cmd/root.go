package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "Herramienta CLI para análisis de datos",
    Long:  `mycli es una herramienta de línea de comandos para automatizar tareas de análisis de datos en Go.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
