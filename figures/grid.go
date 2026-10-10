package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("grid", grid) }

func grid() *plot.Figure {
	x := linspace(0, 1, 200)
	fig, axs := plot.Subplots(2, 2)
	for r := range 2 {
		for c := range 2 {
			k := float64(1 + 2*r + c)
			y := make([]float64, len(x))
			for i, v := range x {
				y[i] = math.Sin(2*math.Pi*k*v) * math.Exp(-k*v)
			}
			ax := axs[r][c]
			ax.Plot(x, y, plot.Color("C"+string(rune('0'+2*r+c))))
			ax.SetTitle(`$k = ` + string(rune('0'+int(k))) + `$`)
			ax.SetXLabel("x")
			ax.SetYLabel(`$\sin(2\pi k x)\,e^{-kx}$`)
		}
	}
	fig.TightLayout()
	return fig
}
