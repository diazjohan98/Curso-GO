package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "CLI de ejemplo para bloque BONUS",
    Long:  `Automatización de flujo de datos con base SQLite, exportación y más.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
