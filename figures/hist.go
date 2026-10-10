package main

import (
	"math"
	"math/rand/v2"

	"github.com/go-plot/plot"
)

func init() { register("hist", hist) }

func hist() *plot.Figure {
	rng := rand.New(rand.NewPCG(1, 2))
	a, b := make([]float64, 5000), make([]float64, 5000)
	for i := range a {
		a[i] = rng.NormFloat64()
		b[i] = 1.5 + 0.6*rng.NormFloat64()
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	ax.Hist(a, plot.Bins(40), plot.Density(), plot.Alpha(0.6), plot.Label(`$\mathcal{N}(0, 1)$`))
	ax.Hist(b, plot.Bins(40), plot.Density(), plot.Alpha(0.6), plot.Label(`$\mathcal{N}(1.5, 0.6^2)$`))
	x := linspace(-4, 4, 200)
	pdf := make([]float64, len(x))
	for i, v := range x {
		pdf[i] = math.Exp(-v*v/2) / math.Sqrt(2*math.Pi)
	}
	ax.Plot(x, pdf, plot.Color("k"), plot.LineWidth(1), plot.Label("density of the first"))
	ax.SetTitle("Hist with 40 bins, normalised")
	ax.Legend()
	return fig
}
