package main

import "github.com/go-plot/plot"

func init() { register("pie", pie) }

func pie() *plot.Figure {
	spend := []float64{42, 27, 18, 13}
	names := []string{"Salaries", "Rent", "Travel", "Other"}
	fig, axs := plot.Subplots(1, 2, plot.FigSize(8, 4))
	left, right := axs[0][0], axs[0][1]
	if err := left.Pie(spend, plot.Labels(names...), plot.AutoPct("%1.1f%%"),
		plot.PlainText(), plot.Explode(0.05, 0, 0, 0)); err != nil {
		panic(err)
	}
	left.SetTitle("Pie")
	if err := right.Donut(spend, 0.35, plot.Labels(names...), plot.AutoPct("%.0f%%"),
		plot.PlainText(), plot.EdgeColor("w"), plot.StartAngle(90), plot.CounterClock(false)); err != nil {
		panic(err)
	}
	right.SetTitle("Donut, from 12 o'clock, clockwise")
	return fig
}
