package main

import (
	"math"

	"github.com/go-plot/plot"
)

func init() { register("twinx", twinx) }

func twinx() *plot.Figure {
	t := linspace(0, 24, 97)
	temp, hum := make([]float64, len(t)), make([]float64, len(t))
	for i, v := range t {
		temp[i] = 14 + 6*math.Sin((v-9)*math.Pi/12)
		hum[i] = 70 - 15*math.Sin((v-9)*math.Pi/12)
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	ax.Plot(t, temp, plot.Color("C3"))
	ax.SetXLabel("hour")
	ax.SetYLabel("temperature (°C)")
	right := ax.TwinX()
	right.Plot(t, hum, plot.Color("C0"), plot.LineStyle("--"))
	right.SetYLabel("relative humidity (%)")
	ax.SetTitle("TwinX: two y axes on one x axis")
	return fig
}
