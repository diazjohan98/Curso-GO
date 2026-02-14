package main

import (
    "log"
    "os"
    "strconv"

    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/opts"
    "github.com/go-echarts/go-echarts/v2/types"

    "github.com/go-gota/gota/dataframe"
)

func main() {
    // Abrir CSVs
    f1, err := os.Open("enero.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f1.Close()
    dfEnero := dataframe.ReadCSV(f1)

    f2, err := os.Open("febrero.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f2.Close()
    dfFebrero := dataframe.ReadCSV(f2)

    df := dfEnero.RBind(dfFebrero)

    // Extraer columnas
    edadesStr := df.Col("edad").Records()
    ingresosStr := df.Col("ingresos").Records()

    // Crear datos para scatterplot
    var scatterData []opts.ScatterData

    for i := range edadesStr {
        edad, err1 := strconv.ParseFloat(edadesStr[i], 64)
        ingreso, err2 := strconv.ParseFloat(ingresosStr[i], 64)
        if err1 != nil || err2 != nil {
            continue
        }
        scatterData = append(scatterData, opts.ScatterData{Value: [2]float64{edad, ingreso}})
    }

    // Crear scatter chart
    s := charts.NewScatter()
    s.SetGlobalOptions(
        charts.WithTitleOpts(opts.Title{
            Title:    "Relación Edad - Ingresos",
            Subtitle: "Demo con go-echarts",
        }),
        charts.WithInitializationOpts(opts.Initialization{
            Theme: types.ThemeWesteros,
        }),
        charts.WithXAxisOpts(opts.XAxis{
            Name: "Edad",
        }),
        charts.WithYAxisOpts(opts.YAxis{
            Name: "Ingresos",
        }),
    )

    s.SetXAxis([]string{}).AddSeries("Edad-Ingresos", scatterData)

    // Crear carpeta output si no existe
    os.MkdirAll("output", os.ModePerm)

    // Crear archivo HTML
    f, err := os.Create("output/scatter_dashboard.html")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    s.Render(f)

    log.Println(" Dashboard HTML creado en output/scatter_dashboard.html")
}
