package main

import "github.com/go-plot/plot"

func init() { register("bars", bars) }

func bars() *plot.Figure {
	x := []float64{0, 1, 2, 3, 4}
	a := []float64{3, 7, 5, 9, 4}
	b := []float64{2, 4, 6, 3, 5}
	fig, axs := plot.Subplots(1, 2, plot.FigSize(8, 3.6))
	left, right := axs[0][0], axs[0][1]
	left.Bar(x, a, plot.Width(0.4), plot.Label("a"))
	shifted := make([]float64, len(x))
	for i, v := range x {
		shifted[i] = v + 0.4
	}
	left.Bar(shifted, b, plot.Width(0.4), plot.Label("b"))
	left.SetXTicks([]float64{0.2, 1.2, 2.2, 3.2, 4.2}, "Mon", "Tue", "Wed", "Thu", "Fri")
	left.SetTitle("Bar")
	left.Legend()
	right.Barh(x, a, plot.Colors("C0", "C1", "C2", "C3", "C4"))
	right.SetTitle("Barh, one colour per bar")
	fig.TightLayout()
	return fig
}
