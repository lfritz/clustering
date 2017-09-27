package main

import (
	"github.com/lfritz/clustering"
	"github.com/lfritz/clustering/dbscan"
	"github.com/lfritz/clustering/draw"
	"github.com/lfritz/clustering/rand"
)

func main() {
	points := make([]clustering.Point, 100)
	rand.Norm2D(clustering.Point{X: 0.2, Y: 0.2}, 0.05, points[:50])
	rand.Norm2D(clustering.Point{X: 0.8, Y: 0.5}, 0.05, points[50:])
	clustering := dbscan.Dbscan(points, 0.04, 4)
	draw.ToPNG(points, clustering, "points.png")
}
