// Package draw contains functions to draw clustered points.
package draw

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func toImage(points [][2]float64, clustering []int) image.Image {
	colors := []color.RGBA{
		{0x00, 0x00, 0x00, 0xff}, // Unclassified
		{0xaa, 0xaa, 0xaa, 0xff}, // Noise
		{0x44, 0x44, 0xff, 0xff},
		{0x44, 0xff, 0x44, 0xff},
		{0x44, 0xff, 0xff, 0xff},
		{0xff, 0x44, 0x44, 0xff},
		{0xff, 0x44, 0xff, 0xff},
		{0xff, 0xff, 0x44, 0xff},
	}
	im := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for i := range im.Pix {
		im.Pix[i] = 0xff
	}
	for i, p := range points {
		x := int(100 * p[0])
		y := int(100 * (1.0 - p[1]))
		if x < 0 || y < 0 || x >= 100 || y >= 100 {
			continue
		}
		c := colors[clustering[i]%len(colors)]
		im.Set(x, y, c)
	}
	return im
}

func storeImage(image image.Image, path string) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			panic(err)
		}
	}()
	png.Encode(w, image)
}

// ToPNG creates a PNG of the points and writes it to the given path.
func ToPNG(points [][2]float64, clustering []int, path string) {
	storeImage(toImage(points, clustering), path)
}
