package draw

import (
	"fmt"
	"github.com/ajstarks/svgo"
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

// ToSVG draws points to SVG, with colors indicating clusters.
func ToSVG(floatPoints [][2]float64, clustering []int, w io.Writer) {
	const radius = 4
	const scale = 500
	const margin = 20
	colors := []string{
		"black", // Unclassified
		"gray",  // Noise
		"blue",
		"red",
		"forestgreen",
		"tomato",
		"deeppink",
		"darkviolet",
		"brown",
	}

	points := toInt(floatPoints, scale)

	canvas := svg.New(w)
	width, height := scale+2*margin, scale+2*margin
	canvas.Start(width, height)

	canvas.Rect(0, 0, width, height, "fill:none;stroke:black;stroke-width:4")

	for i, p := range points {
		color := colors[clustering[i]%len(colors)]
		x := margin + p[0]
		y := height - margin - p[1]
		canvas.Circle(x, y, radius, fmt.Sprintf("fill:%s", color))
	}

	canvas.End()
}