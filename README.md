# docs

[![docs](https://github.com/go-plot/docs/actions/workflows/docs.yml/badge.svg)](https://github.com/go-plot/docs/actions/workflows/docs.yml)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-1f77b4?style=flat-square)](LICENSE)

**The go-plot documentation site**, published at
<https://go-plot.github.io/docs/>.

A [Hugo](https://gohugo.io/) site with the [tannevaled/hextra](https://github.com/tannevaled/hextra)
fork of the [Hextra](https://github.com/imfing/hextra) theme, built the same way
as the [go-ndarray documentation](https://github.com/go-ndarray/docs), with
go-plot's branding. It documents
[`go-plot/plot`](https://github.com/go-plot/plot).

## Versions

A version of the documentation is **the minor version of
[`go-plot/plot`](https://github.com/go-plot/plot) it describes**, or `main`
while no release is tagged:

| Source | Published under |
| --- | --- |
| branch `main`, with `params.plot.version` in `hugo.yaml` | `https://go-plot.github.io/docs/<MAJOR.MINOR>/`, or `/docs/main/` while it says `main` |
| the newest release (else `main`) | also `https://go-plot.github.io/docs/latest/` |

`https://go-plot.github.io/docs/` redirects to `latest/`. Each push to `main`
replaces the directory of the version it describes, and nothing else. The
pages describe **v0.1.0**, published under `0.1/`.

To document a new release, set `params.plot.version` to its tag, require the
same version in `figures/go.mod`, redraw the figures, and update the pages.

## Figures

Every figure on the site is drawn by go-plot. `figures/` is a Go module, one
file per figure; `go run .` there writes `static/images/figures/<name>.svg`.
The `plotfig` shortcode (`layouts/_shortcodes/plotfig.html`) shows an image
and then **reads the file that drew it**, so the code on the page is the code
that ran. The `figures` job of the CI draws them all again and fails if an
image differs from the committed one, or if `figures/go.mod` requires another
go-plot than the version the pages describe.

```sh
cd figures && go run .
```

## Layout

| Path | What |
| --- | --- |
| `content/en/` | the pages; `_index.md` is the home, `weight` orders the sidebar |
| `i18n/en.yaml` | the site's strings, beside the theme's |
| `hugo.yaml` | `baseURL`, the go-plot version described, the navbar (version › theme › search › GitHub), the brand mounts |
| `figures/` | the Go module that draws every figure, one file each |
| `layouts/_shortcodes/plotfig.html` | a figure and the code that drew it |
| `static/images/figures/` | the figures, as `figures/` writes them |
| `layouts/_partials/custom/version-select.html` | the version selector (reads `versions.json`), from go-ndarray/docs |
| `layouts/_partials/custom/content-begin.html` | the outdated-version banner (reads `versions.json`) |
| `layouts/_partials/favicons.html` | the brand's favicons |
| `assets/css/custom.css` | the brand blue as Hextra's primary colour, the table headers, a white card under figures |
| `themes/hextra` | submodule: [tannevaled/hextra](https://github.com/tannevaled/hextra), pinned to a release tag |
| `branding` | submodule: [go-plot/brand](https://github.com/go-plot/brand) (the mark, its PNGs and ICO) |
| `scripts/publish-version.sh` | puts one build into a `gh-pages` checkout and rewrites `versions.json`, `latest` and the root redirect |

## Working locally

```sh
git clone --recurse-submodules https://github.com/go-plot/docs.git
cd docs
mkdir -p data && sh themes/hextra/scripts/page-history.sh > data/pagehistory.json
hugo server --baseURL http://localhost:1313/docs/0.1/
```

then open <http://localhost:1313/docs/0.1/>. Hugo 0.146 or later; no Python,
no Node. Go 1.27.2 or later only to redraw the figures.

## Publication

`.github/workflows/docs.yml`:

| Job | When | What |
| --- | --- | --- |
| `build` | pull requests, `main` | builds with `--baseURL …/docs/<version>/`, checks internal links and anchors (lychee, offline), uploads the site |
| `links` | pull requests, `main` | fetches every external URL the pages name |
| `figures` | pull requests, `main` | draws every figure again and compares it with the committed image |
| `deploy` | `main` | `scripts/publish-version.sh` into `gh-pages` (started empty on the first deploy), then a commit and a push |

GitHub Pages serves the `gh-pages` branch.

## Licence

BSD-3-Clause.
