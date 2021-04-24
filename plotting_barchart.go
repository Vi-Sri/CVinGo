package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
)

func main() {

		// Create the plot values and labels.
		values := plotter.Values{0.5, 10, 20, 30}
		verticalLabels := []string{"A", "B", "C", "D"}
		horizontalLabels := []string{"Label A", "Label B", "Label C", "Label D"}

		// Create a vertical BarChart
		p1 := plot.New()
		verticalBarChart, err := plotter.NewBarChart(values, 1*vg.Centimeter)
		if err != nil {
			log.Panic(err)
		}
		p1.Add(verticalBarChart)
		p1.NominalX(verticalLabels...)
		err = p1.Save(200, 200, "verticalBarChart.png")
		if err != nil {
			log.Panic(err)
		}

		p2 := plot.New()
		horizontalBarChart, err := plotter.NewBarChart(values, 1*vg.Centimeter)
		horizontalBarChart.Horizontal = true // Specify a horizontal BarChart.
		if err != nil {
			log.Panic(err)
		}
		p2.Add(horizontalBarChart)
		p2.NominalY(horizontalLabels...)
		err = p2.Save(200, 200, "horizontalBarChart.png")
		if err != nil {
			log.Panic(err)
		}
}
