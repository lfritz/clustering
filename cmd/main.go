package main

import (
	"fmt"
	"github.com/lfritz/clustering/dbscan"
	"github.com/lfritz/clustering/draw"
	"github.com/lfritz/clustering/index"
	"github.com/lfritz/clustering/rand"
	"os"
)

func main() {
	points := make([][2]float64, 150)
	rand.Norm2D([2]float64{0.2, 0.2}, 0.05, points[:50])
	rand.Norm2D([2]float64{0.8, 0.5}, 0.05, points[50:100])
	rand.Norm2D([2]float64{0.5, 0.6}, 0.05, points[100:])

	index := index.NewBasicKDTree(points)

	save(points, dbscan.Dbscan(index, 0.04, 4), "dbscan")
	save(points, dbscan.Kmeans(points, 3), "kmeans")
}

func save(points [][2]float64, clustering []int, name string) {
	svgFile, err := os.Create(fmt.Sprintf("%s.svg", name))
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err.Error())
		return
	}
	draw.ToSVG(points, clustering, svgFile)
}
