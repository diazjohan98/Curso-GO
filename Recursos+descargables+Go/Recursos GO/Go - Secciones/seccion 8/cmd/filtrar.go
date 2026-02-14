package cmd

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"

    "github.com/go-gota/gota/dataframe"
    "github.com/spf13/cobra"
)

var edadMinima int
var salida string

var filtrarCmd = &cobra.Command{
    Use:   "filtrar [archivo]",
    Short: "Filtrar filas por edad mínima y exportar resultado",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        archivo := args[0]

        f, err := os.Open(archivo)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()

        df := dataframe.ReadCSV(f)

        // Filtrar edad >= edadMinima
        filtro := df.Filter(
            dataframe.F{
                Colname:    "edad",
                Comparator: ">=",
                Comparando: edadMinima,
            },
        )

        fmt.Printf("[%s] Filas originales: %d\n", archivo, df.Nrow())
        fmt.Printf("Filas filtradas (edad >= %d): %d\n", edadMinima, filtro.Nrow())

        // Crear archivo de salida
        outFile, err := os.Create(salida)
        if err != nil {
            log.Fatal(err)
        }
        defer outFile.Close()

        w := csv.NewWriter(outFile)
        w.Write(filtro.Names()) // encabezado

        records := filtro.Records()
        for _, row := range records {
            w.Write(row)
        }
        w.Flush()

        fmt.Printf("Archivo filtrado guardado en: %s\n", salida)
    },
}

func init() {
    rootCmd.AddCommand(filtrarCmd)

    filtrarCmd.Flags().IntVarP(&edadMinima, "edad", "e", 30, "Edad mínima para filtrar")
    filtrarCmd.Flags().StringVarP(&salida, "salida", "o", "filtrado.csv", "Nombre del archivo CSV de salida")
}
