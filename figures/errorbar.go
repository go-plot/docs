package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("errorbar", errorbar) }

func errorbar() *plot.Figure {
	x := linspace(0, 10, 11)
	y, e := make([]float64, len(x)), make([]float64, len(x))
	for i, v := range x {
		y[i] = 2 + math.Sin(v/2)
		e[i] = 0.1 + 0.03*v
	}
	lo, hi := make([]float64, len(x)), make([]float64, len(x))
	for i := range x {
		lo[i], hi[i] = y[i]-2*e[i], y[i]+2*e[i]
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	ax.FillBetween(x, lo, hi, plot.Alpha(0.2), plot.Label(`$\pm 2\sigma$`))
	ax.ErrorBar(x, y, plot.YErr(e), plot.CapSize(3), plot.Marker("o"), plot.Label(`measured $\pm\sigma$`))
	ax.Axhline(2, plot.Color("gray"), plot.LineStyle("--"), plot.LineWidth(1))
	ax.Annotate("baseline", 8, 2, 6.5, 1.3, plot.Arrow())
	ax.SetTitle("ErrorBar, FillBetween, Axhline, Annotate")
	ax.Legend(plot.Loc("upper left"))
	return fig
}
