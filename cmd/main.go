package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"github.com/lfritz/clustering"
	"github.com/lfritz/clustering/dbscan"
)

func gaussian(mean clustering.Point, stdDev float64, output []clustering.Point) {
	for i := range output {
		x := rand.NormFloat64()*stdDev + mean.X
		y := rand.NormFloat64()*stdDev + mean.Y
		output[i] = clustering.Point{x, y}
	}
}

func toImage(points []clustering.Point, clustering []int) image.Image {
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
		x := int(100 * p.X)
		y := int(100 * (1.0 - p.Y))
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

func toPNG(points []clustering.Point, clustering []int, path string) {
	storeImage(toImage(points, clustering), path)
}

func main() {
	points := make([]clustering.Point, 100)
	gaussian(clustering.Point{0.2, 0.2}, 0.05, points[:50])
	gaussian(clustering.Point{0.8, 0.5}, 0.05, points[50:])
	clustering := dbscan.Dbscan(points, 0.04, 4)
	toPNG(points, clustering, "points.png")
}
