package cmd

import (
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
    "github.com/spf13/cobra"
)

var delimitador string

var contarCmd = &cobra.Command{
    Use:   "contar [archivo]",
    Short: "Contar filas de un archivo CSV",
    Args:  cobra.ExactArgs(1), // Requiere exactamente 1 argumento posicional
    Run: func(cmd *cobra.Command, args []string) {
        archivo := args[0] // El argumento posicional

        f, err := os.Open(archivo)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()

        df := dataframe.ReadCSV(f, dataframe.WithDelimiter([]rune(delimitador)[0]))
		cols := df.Ncol()
		fmt.Printf("[%s] Filas: %d | Columnas: %d | Encabezados: %v\n", archivo, df.Nrow(), cols, df.Names())

    },
}

func init() {
    rootCmd.AddCommand(contarCmd)

    // Flag con valor por defecto
    contarCmd.Flags().StringVarP(&delimitador, "delimitador", "d", ",", "Delimitador del archivo CSV (por defecto ',')")
}
