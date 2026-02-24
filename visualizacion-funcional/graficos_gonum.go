package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	pts := make(plotter.XYs, len(edadesStr))

	for i := range edadesStr {
		edad, err1 := strconv.ParseFloat(edadesStr[i], 64)
		ingreso, err2 := strconv.ParseFloat(ingresosStr[i], 64)
		if err1 != nil || err2 != nil {
			continue
		}
		pts = append(pts, plotter.XY{X: edad, Y: ingreso})
	}

	p := plot.New()
	p.Title.Text = "Relacion Edad - Ingresos"
	p.X.Label.Text = "Edad"
	p.Y.Label.Text = "Ingresos"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(s)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "edad_ingresos.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("Grafico guardado como edad_ingresos.png")

}
