package main

import (
	"math"

	"github.com/go-ndarray/ndarray"
	"github.com/go-plot/plot"
)

func init() { register("imshow", imshow) }

func imshow() *plot.Figure {
	const rows, cols = 120, 160
	data := make([]float64, rows*cols)
	for i := range rows {
		for j := range cols {
			data[i*cols+j] = math.Sin(float64(i)/11) * math.Cos(float64(j)/17)
		}
	}
	z, err := ndarray.FromSlice(data, rows, cols)
	if err != nil {
		panic(err)
	}
	fig, axs := plot.Subplots(1, 1)
	ax := axs[0][0]
	m := ax.Imshow(z, plot.Cmap("viridis"))
	fig.Colorbar(m, ax)
	ax.SetTitle(`Imshow of a $120 \times 160$ ndarray`)
	return fig
}
