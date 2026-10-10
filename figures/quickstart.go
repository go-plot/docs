package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("quickstart", quickstart) }

// The README's example, with SaveSVG in place of SavePNG.
func quickstart() *plot.Figure {
	x := make([]float64, 200)
	y := make([]float64, 200)
	for i := range x {
		x[i] = float64(i) / 20
		y[i] = math.Sin(x[i])
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	ax.Plot(x, y, plot.Label(`$\sin t$`))
	ax.Plot(x, y, plot.Marker("o"), plot.LineStyle("None"), plot.MarkerSize(2))
	ax.SetXLabel(`time $t$ (s)`)
	ax.Legend()
	return fig
}
