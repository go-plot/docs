package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("tex", tex) }

func tex() *plot.Figure {
	x := linspace(-4, 4, 400)
	y := make([]float64, len(x))
	for i, v := range x {
		y[i] = math.Exp(-v*v/2) / math.Sqrt(2*math.Pi)
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	ax.Plot(x, y, plot.Label(`$\varphi(x) = \frac{1}{\sqrt{2\pi}}\, e^{-x^2/2}$`))
	ax.FillBetween(x, y, nil, plot.Where(between(x, -1, 1)), plot.Alpha(0.3))
	ax.Text(1.4, 0.25, `$\int_{-1}^{1} \varphi(x)\,dx \approx 0.683$`, plot.FontSize(12))
	ax.SetTitle(`The standard normal density $\varphi$`)
	ax.SetXLabel(`$x$`)
	ax.SetYLabel(`$\varphi(x)$`)
	ax.Legend(plot.Loc("upper left"))
	return fig
}

func between(x []float64, lo, hi float64) []bool {
	m := make([]bool, len(x))
	for i, v := range x {
		m[i] = v >= lo && v <= hi
	}
	return m
}
