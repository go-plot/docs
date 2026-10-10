package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("loglog", loglog) }

func loglog() *plot.Figure {
	n := linspace(1, 6, 26)
	x := make([]float64, len(n))
	for i, e := range n {
		x[i] = math.Pow(10, e)
	}
	lin, nlogn, sq := make([]float64, len(x)), make([]float64, len(x)), make([]float64, len(x))
	for i, v := range x {
		lin[i], nlogn[i], sq[i] = v, v*math.Log2(v), v*v
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	ax.Plot(x, lin, plot.Label(`$n$`))
	ax.Plot(x, nlogn, plot.Label(`$n \log_2 n$`))
	ax.Plot(x, sq, plot.Label(`$n^2$`))
	ax.SetXScale("log")
	ax.SetYScale("log")
	ax.Grid(true)
	ax.SetXLabel(`$n$`)
	ax.SetYLabel("operations")
	ax.SetTitle("Log scales: LogLocator and LogFormatterSciNotation")
	ax.Legend()
	return fig
}
