---
title: "TeX labels"
weight: 50
description: "Math in titles, labels, legends and text, set by a TeX typesetter rather than approximated."
tags: [text, tex]
---

Any string go-plot draws — a title, an axis label, a legend entry, a colorbar
label, `Text`, `Annotate` — may contain **TeX math between dollar signs**:

```go
ax.SetXLabel(`time $t$ (s)`)
ax.Plot(x, y, plot.Label(`$e^{-t/4}\cos 2t$`))
```

Use Go's raw strings (backquotes), so that backslashes reach TeX unescaped.

## A TeX typesetter, not mathtext

matplotlib sets `$...$` with **mathtext**, its own approximation of TeX.
go-plot sets it with [go-tex/math](https://github.com/go-tex/math), a real TeX
math typesetter using the OpenType MATH table of **STIX Two Math**: fractions,
radicals, big operators, limits, matrices and `\left...\right` are laid out as
TeX lays them out, which mathtext only approximates. go-tex/math's output is
parsed back into paths and scaled to the label's size, so TeX is drawn as
outlines in both PNG and SVG and needs no font installed.

{{< plotfig "tex" "The standard normal density with a TeX legend entry, a shaded area between -1 and 1, and an integral written as text" >}}

## Tick labels keep matplotlib's look

The one exception is `\mathdefault{...}`, which is what matplotlib's tick
formatters emit (`10^{3}`, `2\times10^{1}`). It is set in **DejaVu Sans** with
mathtext's own superscript rules (shift = max(sup1 · x-height, depth +
x-height/4), scripts at 0.7 size), so tick labels match matplotlib's look and
position. The decades on the [log-log figure]({{< relref "plot-types#log-scales" >}})
are set this way.

## Plain text

Text outside dollar signs is set in **DejaVu Sans**, matplotlib's default
font, from go-opentype/fonts, with kerning from DejaVu's GPOS and kern tables.
Glyphs are drawn as filled outlines, so rotated labels and SVG output need no
font. Positions follow matplotlib's `Text._get_layout` (line metrics from the
OS/2 table, alignment, rotation, anchoring), and like Agg the pen is put on a
whole pixel. How the pixels of text differ from matplotlib's is on
[Differences]({{< relref "differences#text" >}}).

## Text from data: no TeX

A string is read as matplotlib reads it: TeX between an even number of
unescaped dollar signs. Names that come from data ("100$", a team called
`$x$`) should not be. go-plot follows matplotlib's `parse_math`:

| go-plot | matplotlib |
| --- | --- |
| `plot.NewFigure(plot.ParseMath(false))`, `fig.SetParseMath(false)` | `rcParams["text.parse_math"] = False` |
| `plot.PlainText()` on one call | `parse_math=False` on that text |
| `plot.MathText()` on one call | `parse_math=True` on that text |

The calls that take it are `SetTitle`, `SetXLabel` and `SetYLabel`. So do
`SetXTickLabels` and `SetYTickLabels` (category names under bars), `Legend`
(every entry), `Text`, `Annotate` and `Pie`. A plain string is drawn
character for character as glyph outlines. It never reaches the TeX
typesetter, and `<script>` becomes shapes, not markup. The tick labels
go-plot's formatters write (`10³` on a log axis) stay math.

## Cost

STIX Two Math is parsed only when a label contains TeX: a figure without any
does not pay for it ([performance]({{< relref "performance" >}})).

## Replacing the renderer

`Figure.Text` is a `text.Renderer` with two methods, `Layout(s, size)` and
`LineMetrics(size)`. Setting another one changes the fonts of the whole figure.
