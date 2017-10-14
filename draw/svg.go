// Package draw provides functions to draw 2-D points as SVG.
package draw

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/lfritz/clustering"
	"io"
)

func toInt(points [][2]float64, scale float64) [][2]int {
	result := make([][2]int, len(points))
	for i, p := range points {
		result[i] = [2]int{
			int(scale * p[0]),
			int(scale * p[1]),
		}
	}
	return result
}

func color(clusterID int) string {
	colors := []string{
		"blue",
		"red",
		"forestgreen",
		"tomato",
		"deeppink",
		"darkviolet",
		"brown",
	}
	switch clusterID {
	case clustering.Unclassified:
		return "black"
	case clustering.Noise:
		return "gray"
	default:
		return colors[clusterID%len(colors)]
	}
}

// ToSVG draws points to SVG, with colors indicating clusters.
func ToSVG(floatPoints [][2]float64, cl []int, w io.Writer) {
	const radius = 4
	const scale = 500
	const margin = 20

	points := toInt(floatPoints, scale)

	canvas := svg.New(w)
	width, height := scale+2*margin, scale+2*margin
	canvas.Start(width, height)

	canvas.Rect(0, 0, width, height, "fill:none;stroke:black;stroke-width:4")

	for i, p := range points {
		x := margin + p[0]
		y := height - margin - p[1]
		canvas.Circle(x, y, radius, fmt.Sprintf("fill:%s", color(cl[i])))
	}

	canvas.End()
}
