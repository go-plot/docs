---
title: "The API"
linkTitle: "API"
weight: 20
description: "Figure, Subplots and Axes: matplotlib's object model with Go's argument style, and why it has that shape."
tags: [api]
---

## Shaped by what notebooks call

go-plot is the plotting library of a Go notebook ecosystem built on
[go-ndarray](https://github.com/go-ndarray/ndarray). Its API was chosen from a
measurement, not a preference.

A census of **339 scientific notebooks** (measured for this ecosystem) ranks
plotting first among the things that keep them from moving to Go, and gives the
share of notebooks using each call:

| call | notebooks | go-plot |
| --- | --- | --- |
| `pyplot.subplots` | 62 % | `plot.Subplots` |
| `Axes.plot` | 56 % | `Axes.Plot` |
| `pyplot.show` | 52 % | `Figure.Display`, called by the notebook |
| `Axes.set_xlabel` | 47 % | `Axes.SetXLabel` |
| `Axes.legend` | 46 % | `Axes.Legend`, with matplotlib's location names and `"best"` |
| `Axes.set_ylabel` | 45 % | `Axes.SetYLabel` |
| `Axes.set_title` | 37 % | `Axes.SetTitle` |
| `pyplot.figure` | 27 % | `plot.NewFigure` |
| `tight_layout` | 27 % | `Figure.TightLayout` |
| module-level `pyplot.plot`/`xlabel`/`ylabel` | about 22 % | not provided (below) |
| `Axes.set_ylim` | 21 % | `Axes.SetYLim` |
| `Axes.axhline` | 20 % | `Axes.Axhline` |

Most notebooks write `fig, ax = plt.subplots(); ax.plot(...)`. So go-plot is
the **object-oriented API**, and every one of those operations exists under the
same name.

**The module-level, state-machine functions are deliberately not provided.**
`plt.plot` and `plt.xlabel` act on a hidden "current" Axes. They are the
minority form, and hidden global state is what a Go notebook kernel running
cells concurrently should not have.

## Figure, Subplots, Axes

```go
fig, axs := plot.Subplots(2, 3)          // a 2 x 3 grid: axs is a [][]*plot.Axes
ax := axs[1][0]                          // second row, first column

fig := plot.NewFigure(plot.FigSize(6.4, 4.8), plot.DPI(100))
grid := fig.Subplots(1, 2)               // the same, on an existing figure
inset := fig.AddAxes(0.6, 0.6, 0.25, 0.25) // left, bottom, width, height, in figure fractions
```

- `plot.Subplots(rows, cols, ...FigOption)` returns the figure and a
  `[][]*Axes`, always two-dimensional (matplotlib's `squeeze=False`).
- Plot types are `Axes` methods named as in matplotlib: `Plot`, `Scatter`,
  `Bar`, `Barh`, `Hist`, `FillBetween`, `ErrorBar`, `Imshow`, `Contour`,
  `Contourf`, `Axhline`, `Axvline`, `Text`, `Annotate`, `Legend`, `SetXLim`,
  `SetYScale`, `TwinX`... See [plot types]({{< relref "plot-types" >}}).
- `Figure.TightLayout()` is matplotlib's `tight_layout`, ported. It reports
  `false`, changing nothing, when the decorations cannot fit, as matplotlib
  then warns and does nothing.

## Keyword arguments become options

matplotlib's keyword arguments become **functional options**:

```go
ax.Plot(x, y, plot.Color("C1"), plot.LineStyle("--"), plot.Label("data"))
```

is `ax.plot(x, y, color="C1", linestyle="--", label="data")`. Go has no
keywords, and a struct per method would multiply types without adding safety.
Each method documents the options it honours; an option a method does not use
is ignored.

| option | matplotlib keyword |
| --- | --- |
| `Color`, `Colors`, `FaceColor`, `EdgeColor`, `Alpha` | `color`, a list of colours, `facecolor`, `edgecolor`, `alpha` |
| `LineStyle`, `LineWidth`, `Marker`, `MarkerSize`, `MarkerEdgeWidth` | `linestyle`, `linewidth`, `marker`, `markersize`, `markeredgewidth` |
| `Label`, `ZOrder`, `Clip` | `label`, `zorder`, `clip_on` |
| `Width`, `Bottom`, `Align` | `bar`'s `width`, `bottom`, `align` |
| `Bins`, `BinEdges`, `Density` | `hist`'s `bins` (a count, or edges), `density` |
| `XErr`, `YErr`, `CapSize`, `ErrorColor` | `errorbar`'s `xerr`, `yerr`, `capsize`, `ecolor` |
| `Where` | `fill_between`'s `where` |
| `Sizes`, `CValues` | `scatter`'s `s` and `c` |
| `Cmap`, `Vmin`, `Vmax`, `LogNorm` | `cmap`, `vmin`, `vmax`, `norm=LogNorm()` |
| `Extent`, `Origin`, `Aspect` | `imshow`'s `extent`, `origin`, `aspect` |
| `Levels`, `NLevels` | `contour`'s `levels` (a list, or a count) |
| `FontSize`, `HA`, `VA`, `Rotation`, `Arrow` | `fontsize`, `ha`, `va`, `rotation`, `arrowprops=dict(arrowstyle="->")` |
| `Loc` | `legend`'s `loc` |

## Data is `any`

Every data argument accepts `[]float64`, `[]float32`, `[]int`, `[]int32`,
`[]int64`, a scalar, or a `*ndarray.Array` of any of ndarray's thirteen dtypes
(converted with `AsType(Float64)`, flattened in row-major order; 2-D arrays for
`Imshow` and contours, 3-D `(rows, cols, 3|4)` arrays as RGB(A) images, `uint8`
taken as bytes).

A type that cannot be plotted **panics**, as does a length mismatch. Both are
programming errors, exactly as an index out of range is, and returning an
error from every `Plot` call would make every notebook cell three lines longer
for a failure that is never handled.

## Mappables and colorbars

`Imshow`, `Scatter` with `CValues`, `Contour` and `Contourf` return a
`*Mappable` (a colormap and a norm). `Figure.Colorbar(m, ax)` draws it, taking
its space from `ax` exactly as `fig.colorbar(m, ax=ax)` does, and returns the
colorbar's `Axes`, whose label can be set:

```go
m := ax.Imshow(z, plot.Cmap("viridis"))
fig.Colorbar(m, ax).SetYLabel("height (m)")
```

## Output

| method | writes |
| --- | --- |
| `SavePNG(w io.Writer) error` | an anti-aliased raster |
| `SaveSVG(w io.Writer) error` | vector SVG: paths, clip paths, images embedded as PNG, as matplotlib's SVG backend does with `svg.fonttype = "path"`; text is drawn as outlines, so no font is needed to show it |
| `Display() map[string][]byte` | the notebook contract: `image/svg+xml` and `image/png` |

## Text is a seam

`Figure.Text` is a `text.Renderer` (`Layout(s, size)`, `LineMetrics(size)`).
The default sets plain text in DejaVu Sans (matplotlib's default font, from
go-opentype/fonts) and TeX between dollar signs with
[go-tex/math](https://github.com/go-tex/math). Replacing the renderer changes
the fonts of the whole figure. See [TeX labels]({{< relref "tex" >}}).

## Packages

| package | what it does |
| --- | --- |
| `plot` | `Figure`, `Axes` and every plot type; layout, autoscaling, legend, colorbar |
| `ticker` | tick locations and labels: AutoLocator (MaxNLocator), ScalarFormatter, LogLocator, LogFormatterSciNotation |
| `colors` | colour parsing, the tab10 cycle, eight colormaps, `Normalize` and `LogNorm` |
| `contour` | marching squares: iso-lines and filled bands on a rectangular grid |
| `render` | the drawing seam: a `Renderer` interface, a raster back-end and an SVG back-end |
| `text` | strings to glyph outlines: DejaVu Sans for text, go-tex/math for TeX |

The full reference is on [pkg.go.dev](https://pkg.go.dev/github.com/go-plot/plot).
