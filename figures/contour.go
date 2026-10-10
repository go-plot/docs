package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("contour", contour) }

func contour() *plot.Figure {
	x, y := linspace(-3, 3, 61), linspace(-2, 2, 41)
	z := make([][]float64, len(y))
	for i, v := range y {
		z[i] = make([]float64, len(x))
		for j, u := range x {
			z[i][j] = (1-u/2+math.Pow(u, 5)+math.Pow(v, 3))*math.Exp(-u*u-v*v)
		}
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	m := ax.Contourf(x, y, z, plot.NLevels(10), plot.Cmap("RdBu_r"))
	ax.Contour(x, y, z, plot.NLevels(10), plot.Color("k"), plot.LineWidth(0.5))
	fig.Colorbar(m, ax)
	ax.SetTitle("Contourf and Contour")
	return fig
}
