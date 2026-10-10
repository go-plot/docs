package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("lines", lines) }

func lines() *plot.Figure {
	x := linspace(0, 2*math.Pi, 40)
	styles := []string{"-", "--", "-.", ":"}
	markers := []string{"o", "s", "^", "D"}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	for k := range styles {
		y := make([]float64, len(x))
		for i, v := range x {
			y[i] = math.Sin(v + float64(k)*math.Pi/4)
		}
		ax.Plot(x, y, plot.LineStyle(styles[k]), plot.Marker(markers[k]),
			plot.MarkerSize(4), plot.Label(`style "`+styles[k]+`", marker "`+markers[k]+`"`))
	}
	ax.SetTitle("Line styles and markers, colours from the cycle")
	ax.SetYLim(-2, 1.2)
	ax.Legend(plot.Loc("lower center"), plot.FontSize(8))
	return fig
}
