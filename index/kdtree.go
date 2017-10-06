package index

import (
	"github.com/lfritz/clustering"
	"sort"
)

type KDTree struct {
	points []clustering.Point
	node   *node
}

type node struct {
	p     int
	value float64
	left  *node
	right *node
}

func NewKDTree(points []clustering.Point) *KDTree {
	t := &KDTree{points, nil}
	indices := make([]int, len(points))
	for i := range indices {
		indices[i] = i
	}
	t.node = newNode(points, indices, false)
	return t
}

type byX struct {
	points  []clustering.Point
	indices []int
}

func (a byX) Len() int           { return len(a.indices) }
func (a byX) Swap(i, j int)      { a.indices[i], a.indices[j] = a.indices[j], a.indices[i] }
func (a byX) Less(i, j int) bool { return a.points[a.indices[i]].X < a.points[a.indices[j]].X }

type byY struct {
	points  []clustering.Point
	indices []int
}

func (a byY) get(i int) clustering.Point { return a.points[a.indices[i]] }
func (a byY) Len() int                   { return len(a.indices) }
func (a byY) Swap(i, j int)              { a.indices[i], a.indices[j] = a.indices[j], a.indices[i] }
func (a byY) Less(i, j int) bool         { return a.points[a.indices[i]].Y < a.points[a.indices[j]].Y }

func newNode(points []clustering.Point, indices []int, xAxis bool) *node {
	// base case: no points left in this area
	if len(indices) == 0 {
		return nil
	}

	// sort points and select median point
	if xAxis {
		sort.Sort(byX{points, indices})
	} else {
		sort.Sort(byY{points, indices})
	}
	i := len(indices) / 2
	for i > 0 && points[indices[i]].Equal(points[indices[i-1]]) {
		i--
	}
	p := indices[i]

	// create node and child nodes
	node := &node{p: p}
	if xAxis {
		node.value = points[p].X
	} else {
		node.value = points[p].Y
	}
	node.left = newNode(points, indices[:i], !xAxis)
	node.right = newNode(points, indices[i+1:], !xAxis)

	return node
}

func (t *KDTree) Points() []clustering.Point {
	return t.points
}

func (t *KDTree) BoundingBox(x0, x1, y0, y1 float64) []int {
	return t.node.bb(t.points, x0, x1, y0, y1, false)
}

func (n *node) bb(points []clustering.Point, x0, x1, y0, y1 float64, xAxis bool) []int {
	result := []int{}
	if n == nil {
		return result
	}

	point := points[n.p]
	if point.X >= x0 && point.X < x1 && point.Y >= y0 && point.Y < y1 {
		result = append(result, n.p)
	}

	lookLeft := xAxis && x0 < n.value || !xAxis && y0 < n.value
	lookRight := xAxis && x1 >= n.value || !xAxis && y1 >= n.value
	if lookLeft {
		result = append(result, n.left.bb(points, x0, x1, y0, y1, !xAxis)...)
	}
	if lookRight {
		result = append(result, n.right.bb(points, x0, x1, y0, y1, !xAxis)...)
	}

	return result
}

func (i *KDTree) Circle(center clustering.Point, radius float64) []int {
	return circle(i, center, radius)
}
