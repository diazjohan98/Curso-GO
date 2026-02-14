package main

import (
    "log"
    "os"
    "strconv"
	"time"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"

    "github.com/go-gota/gota/dataframe"
)


func main() {
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


    edadesStr := df.Col("edad").Records()
    ingresosStr := df.Col("ingresos").Records()

    pts := make(plotter.XYs, 0, len(edadesStr))

    for i := range edadesStr {
        edad, err1 := strconv.ParseFloat(edadesStr[i], 64)
        ingreso, err2 := strconv.ParseFloat(ingresosStr[i], 64)
        if err1 != nil || err2 != nil {
            continue
        }
        pts = append(pts, plotter.XY{X: edad, Y: ingreso})
    }


    p := plot.New()
    p.Title.Text = "Relación Edad - Ingresos"
    p.X.Label.Text = "Edad"
    p.Y.Label.Text = "Ingresos"

    s, err := plotter.NewScatter(pts)
    if err != nil {
        log.Fatal(err)
    }

    p.Add(s)	
	filename := "output/scatter_" + time.Now().Format("20060102_150405") + ".png"

    if err := p.Save(6*vg.Inch, 6*vg.Inch, filename); err != nil {
        log.Fatal(err)
    }

    log.Println("Gráfico guardado como edad_ingresos.svg")

}


/*
png
svg

*/

