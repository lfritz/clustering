package main

import (
	"fmt"
	"github.com/lfritz/clustering/dbscan"
	"github.com/lfritz/clustering/draw"
	"github.com/lfritz/clustering/generate"
	"github.com/lfritz/clustering/index"
	"github.com/lfritz/clustering/kmeans"
	"math"
	"os"
)

func main() {
	var ps1 [][2]float64
	ps1 = append(ps1, generate.Point(50, [2]float64{0.2, 0.2}, 0.05)...)
	ps1 = append(ps1, generate.Point(50, [2]float64{0.8, 0.5}, 0.05)...)
	ps1 = append(ps1, generate.Point(50, [2]float64{0.5, 0.6}, 0.05)...)
	runBoth("example-1", ps1)

	var ps2 [][2]float64
	ps2 = append(ps2, generate.CircularArc(
		200, [2]float64{0.5, 0.5}, 0.4, -math.Pi/4, math.Pi/2, 0.02)...)
	ps2 = append(ps2, generate.Point(50, [2]float64{0.5, 0.6}, 0.05)...)
	ps2 = append(ps2, generate.Point(50, [2]float64{0.2, 0.3}, 0.05)...)
	runBoth("example-2", ps2)
}

func runBoth(name string, points [][2]float64) {
	index := index.NewBasicKDTree(points)

	dbscanClustering := dbscan.Dbscan(index, 0.04, 4)
	kmeansClustering, _ := kmeans.Repeat(points, 3, 3)

	save(points, dbscanClustering, fmt.Sprintf("%s-dbscan", name))
	save(points, kmeansClustering, fmt.Sprintf("%s-kmeans", name))
}

func save(points [][2]float64, clustering []int, name string) {
	svgFile, err := os.Create(fmt.Sprintf("%s.svg", name))
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err.Error())
		return
	}
	draw.ToSVG(points, clustering, svgFile)
}
