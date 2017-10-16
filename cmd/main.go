package main

import (
	"fmt"
	"github.com/lfritz/clustering/dbscan"
	"github.com/lfritz/clustering/draw"
	"github.com/lfritz/clustering/index"
	"github.com/lfritz/clustering/kmeans"
	"github.com/lfritz/clustering/random"
	"os"
)

func main() {
	points := make([][2]float64, 150)
	random.Norm2D([2]float64{0.2, 0.2}, 0.05, points[:50])
	random.Norm2D([2]float64{0.8, 0.5}, 0.05, points[50:100])
	random.Norm2D([2]float64{0.5, 0.6}, 0.05, points[100:])

	index := index.NewBasicKDTree(points)

	dbscanClustering := dbscan.Dbscan(index, 0.04, 4)
	kmeansClustering, _ := kmeans.Kmeans(points, 3)

	save(points, dbscanClustering, "dbscan")
	save(points, kmeansClustering, "kmeans")
}

func save(points [][2]float64, clustering []int, name string) {
	svgFile, err := os.Create(fmt.Sprintf("%s.svg", name))
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err.Error())
		return
	}
	draw.ToSVG(points, clustering, svgFile)
}
