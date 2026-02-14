package cmd

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/go-gota/gota/dataframe"
    "github.com/spf13/cobra"
)

var edadMin int

var exportarCmd = &cobra.Command{
    Use:   "exportar [archivo]",
    Short: "Filtrar por edad mínima, exportar y registrar log",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        archivo := args[0]

        f, err := os.Open(archivo)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()

        df := dataframe.ReadCSV(f)

        // Filtrar
        filtrado := df.Filter(
            dataframe.F{
                Colname:    "edad",
                Comparator: ">=",
                Comparando: edadMin,
            },
        )

        fmt.Printf("Filas originales: %d | Filtradas: %d\n", df.Nrow(), filtrado.Nrow())

        // Crear carpeta de salida
        os.MkdirAll("output", os.ModePerm)

        // Nombre con fecha/hora
        timestamp := time.Now().Format("20060102_150405")
        outputFile := fmt.Sprintf("output/filtrado_%d_%s.csv", edadMin, timestamp)

        // Exportar CSV filtrado
        out, err := os.Create(outputFile)
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()

        w := csv.NewWriter(out)
        w.Write(filtrado.Names())
        records := filtrado.Records()
        for _, row := range records {
            w.Write(row)
        }
        w.Flush()

        fmt.Printf("Archivo generado: %s\n", outputFile)

        // Registrar log
        logFile := "output/historial.log"
        lf, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        defer lf.Close()

        logger := log.New(lf, "INFO: ", log.Ldate|log.Ltime)
        logger.Printf("Exportado %s con edad >= %d\n", outputFile, edadMin)
    },
}

func init() {
    rootCmd.AddCommand(exportarCmd)

    exportarCmd.Flags().IntVarP(&edadMin, "edad", "e", 30, "Edad mínima para filtrar")
}
