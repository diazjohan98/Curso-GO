package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "Herramienta CLI para analisis de datos",
	Long:  "mycli es una herramienta de línea de comandos para realizar análisis de datos de manera eficiente y sencilla.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
