package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("scatter", scatter) }

func scatter() *plot.Figure {
	n := 120
	x, y := make([]float64, n), make([]float64, n)
	c, s := make([]float64, n), make([]float64, n)
	for i := range n {
		t := float64(i) / float64(n) * 6 * math.Pi
		r := 1 + t/4
		x[i], y[i] = r*math.Cos(t), r*math.Sin(t)
		c[i] = t
		s[i] = 10 + 4*t
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	m := ax.Scatter(x, y, plot.CValues(c), plot.Sizes(s), plot.Cmap("plasma"), plot.Alpha(0.8))
	ax.SetAspect(1)
	ax.SetTitle("Scatter: size and colour per point")
	fig.Colorbar(m, ax).SetYLabel(`angle $\theta$ (rad)`)
	return fig
}
