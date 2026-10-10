---
title: "Plot types"
weight: 30
description: "Every plot type of v0, each with the code that drew it."
tags: [api, gallery]
---

Each figure on this page was drawn by go-plot v0.1.0 from the code below it.
The code is not a copy: the page reads it from
[`figures/`](https://github.com/go-plot/docs/tree/main/figures) in this
site's repository, where `go run .` draws every image, and the CI draws them
again and fails if any differs from the one shown. `linspace` is a helper in
[`figures/main.go`](https://github.com/go-plot/docs/blob/main/figures/main.go),
`numpy.linspace`'s equivalent.

| what | go-plot | matplotlib |
| --- | --- | --- |
| lines and markers | `Plot` (17 markers, 4 line styles) | `ax.plot` |
| points | `Scatter` (size and colour per point, colormaps) | `ax.scatter` |
| bars | `Bar`, `Barh` | `ax.bar`, `ax.barh` |
| histograms | `Hist` | `ax.hist` |
| bands, error bars | `FillBetween`, `ErrorBar` | `ax.fill_between`, `ax.errorbar` |
| reference lines | `Axhline`, `Axvline` | `ax.axhline`, `ax.axvline` |
| images | `Imshow` (2-D data or RGB(A)) | `ax.imshow` |
| contours | `Contour`, `Contourf` | `ax.contour`, `ax.contourf` |
| text | `Text`, `Annotate` | `ax.text`, `ax.annotate` |
| scales, axes | `SetXScale("log")`, `TwinX`, `SetXTicks`, `SetAspect` | `ax.set_xscale`, `ax.twinx`, `ax.set_xticks`, `ax.set_aspect` |
| colour scale | `Figure.Colorbar` | `fig.colorbar` |

## Lines and markers

`Plot` draws a line in the next colour of the cycle. `LineStyle` takes `"-"`,
`"--"`, `"-."`, `":"` (and `"None"`); `Marker` takes matplotlib's marker codes.

{{< plotfig "lines" "Four phase-shifted sines, each with a different line style and marker" >}}

## Scatter

With `CValues`, colours come from a colormap and `Scatter` returns a
`*Mappable` for `Colorbar`. `Sizes` gives each point its area in points².

{{< plotfig "scatter" "An Archimedean spiral of points whose size and colour grow with the angle, with a plasma colorbar" >}}

## Bars

`Bar` draws bars 0.8 wide, centred on x, from 0; `Width` and an x offset put
two series side by side. `Barh` is the horizontal form, and `Colors` gives one
colour per bar.

{{< plotfig "bars" "Grouped vertical bars for five weekdays, and horizontal bars in five colours" >}}

## Histograms

`Hist` bins with NumPy's algorithm: 10 equal bins over the data range by
default, `Bins(n)` for n, `BinEdges` for explicit edges, `Density` to
normalise.

{{< plotfig "hist" "Two overlapping normalised histograms of normal samples and the standard normal density" >}}

## Error bars, bands and annotations

{{< plotfig "errorbar" "Points with error bars over a translucent band, a dashed horizontal baseline and an arrow annotation" >}}

## Images

`Imshow` colours 2-D data through a colormap (viridis by default), with pixel
centres at integer coordinates and row 0 at the top; a 3-D array of shape
`(rows, cols, 3 or 4)` is shown as RGB(A). Here the data is a go-ndarray
array.

{{< plotfig "imshow" "A 120 by 160 array of a product of sines shown with viridis and a colorbar" >}}

## Contours

`Contour` draws iso-lines and `Contourf` fills the bands between them, at
levels chosen as matplotlib chooses them (`ContourSet._autolev`). x and y give
the column and row coordinates.

{{< plotfig "contour" "Filled contours of a two-dimensional function in RdBu_r with black contour lines and a colorbar" >}}

## Log scales

`SetXScale("log")` and `SetYScale("log")` use matplotlib's `LogLocator` and
`LogFormatterSciNotation`, so the decades and their labels are matplotlib's.

{{< plotfig "loglog" "n, n log n and n squared on log-log axes with a grid" >}}

## Two y axes

`TwinX` returns a new Axes sharing the x axis, with its y axis on the right and
its own colour cycle.

{{< plotfig "twinx" "Temperature on the left axis and humidity on the right axis over 24 hours" >}}

## A grid, with TightLayout

`Subplots(rows, cols)` lays out a grid with matplotlib's subplot parameters;
`TightLayout` then moves the Axes so that titles and labels do not overlap.

{{< plotfig "grid" "Four damped sines in a 2 by 2 grid with titles and labels, after TightLayout" >}}
