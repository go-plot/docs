---
title: "Differences from matplotlib"
linkTitle: "Differences"
weight: 70
description: "What does not equal matplotlib, and what is not in v0."
tags: [matplotlib]
---

What is equal, and how it is checked, is on [Exactness]({{< relref "exactness" >}}).
This page is the rest.

## What does not equal matplotlib

### Text

Glyphs are filled outlines (go-gfx coverage), not FreeType's hinted
rasterisation with matplotlib's text antialiasing, so **the pixels of text
differ**, and a string's box is its advance where matplotlib's is FreeType's
hinted box (up to a few pixels wider on long labels). Positions are
matplotlib's (`Text._get_layout`: line metrics from the OS/2 table, alignment,
rotation, anchoring), and like Agg the pen is put on a whole pixel. Kerning
comes from DejaVu's GPOS/kern tables through go-opentype.

### Math text is TeX, not mathtext

`$...$` is typeset by go-tex/math with STIX Two Math — real TeX layout
(fractions, radicals, big operators, limits, matrices, `\left...\right`), which
matplotlib's mathtext only approximates. The exception is `\mathdefault{...}`,
what matplotlib's tick formatters emit: it is set in DejaVu Sans with
mathtext's own superscript rules, so tick labels match matplotlib's look and
position. See [TeX labels]({{< relref "tex" >}}).

### Strokes

Anti-aliasing is go-gfx's, not Agg's; widths, caps, joins and dashes are
matplotlib's. Pixel snapping follows matplotlib's rule (odd widths at pixel
centres, even at edges) for paths made of horizontal and vertical segments.
Strokes of more than 512 segments are drawn in pieces and get round ends (see
[performance]({{< relref "performance" >}})).

### Images

Images are resampled nearest-neighbour when enlarged and box-averaged when
reduced; matplotlib's default (`interpolation="auto"`, antialiased resampling
in data space) differs when an image is shown smaller than its resolution, or
less than three times enlarged.

### Contours

Contours are marching squares with the saddle resolved by the cell mean, not
contourpy's `serial` algorithm (`corner_mask`, chunking): levels and colours
are matplotlib's, line vertices differ slightly. Filled bands are drawn per
grid cell.

### Legend "best"

`"best"` counts the vertices, segments, boxes and scatter points under each of
the nine candidate positions and takes the first with none, else the least;
matplotlib's badness also weights offsets and is computed in display space with
a pad. Usually the same place; not always.

### Annotate

`Annotate` draws `arrowstyle="->"` only, shrunk 2 points at both ends.

### SVG

The SVG carries the same geometry but not matplotlib's document structure, ids
or metadata.

### No module-level pyplot

`plt.plot`, `plt.xlabel` and the other functions acting on a hidden current
Axes are not provided, by design: see [the API]({{< relref "api" >}}).

## Not in v0

- suptitle, and `TightLayout` with explicit pads or `rect`;
- symlog, logit and date axes;
- minor ticks on linear axes (off by default in matplotlib too) and
  `AutoMinorLocator`;
- `constrained_layout`;
- `subplots` sharing;
- hatching;
- boxplot, violin, pie, quiver, stem, step, `pcolormesh`, 3-D;
- colorbar `extend` and discrete norms;
- image interpolation other than nearest;
- `Text` bounding boxes;
- legend columns, titles and `bbox_to_anchor`;
- native and browser windows ([go-widgets/window](https://github.com/go-widgets/window));
- PDF output.
