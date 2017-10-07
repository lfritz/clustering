package index

import (
	"github.com/lfritz/clustering/geometry"
	"sort"
)

// A KDTree implements the k-d tree data structure (https://en.m.wikipedia.org/wiki/K-d_tree).
type KDTree struct {
	points [][2]float64
	node   *node
}

type node struct {
	p     int
	value float64
	left  *node
	right *node
}

// NewKDTree returns a new k-d tree.
func NewKDTree(points [][2]float64) *KDTree {
	t := &KDTree{points, nil}
	indices := make([]int, len(points))
	for i := range indices {
		indices[i] = i
	}
	t.node = newNode(points, indices, 0)
	return t
}

// sortBy implements sort.Interface for sorting a slice of indices to points
// taking into account only one dimension.
type byOneDimension struct {
	points    [][2]float64
	dimension int // 0 or 1
	indices   []int
}

func (a *byOneDimension) Len() int      { return len(a.indices) }
func (a *byOneDimension) Swap(i, j int) { a.indices[i], a.indices[j] = a.indices[j], a.indices[i] }
func (a *byOneDimension) Less(i, j int) bool {
	return a.points[a.indices[i]][a.dimension] < a.points[a.indices[j]][a.dimension]
}

func newNode(points [][2]float64, indices []int, level int) *node {
	// base case: no points left in this area
	if len(indices) == 0 {
		return nil
	}

	// sort points and select median point
	sort.Sort(&byOneDimension{points, level % 2, indices})
	i := len(indices) / 2
	for i > 0 && points[indices[i]] == points[indices[i-1]] {
		i--
	}
	p := indices[i]

	// create node and child nodes
	node := &node{p: p}
	node.value = points[p][level%2]
	node.left = newNode(points, indices[:i], level+1)
	node.right = newNode(points, indices[i+1:], level+1)

	return node
}

// Points returns the slice of points.
func (t *KDTree) Points() [][2]float64 {
	return t.points
}

// BoundingBox returns the indices of all points within the given
// axis-aligned bounding box.
func (t *KDTree) BoundingBox(bb *geometry.BoundingBox) []int {
	return t.node.bb(t.points, bb, 0)
}

func (n *node) bb(points [][2]float64, bb *geometry.BoundingBox, level int) []int {
	result := []int{}
	if n == nil {
		return result
	}

	point := points[n.p]
	if bb.Contains(point) {
		result = append(result, n.p)
	}

	lookLeft := bb.From[level%2] < n.value
	lookRight := bb.To[level%2] >= n.value
	if lookLeft {
		result = append(result, n.left.bb(points, bb, level+1)...)
	}
	if lookRight {
		result = append(result, n.right.bb(points, bb, level+1)...)
	}

	return result
}

// Circle returns the indices of all points in the circle with the
// given center and radius.
func (t *KDTree) Circle(center [2]float64, radius float64) []int {
	return circle(t, center, radius)
}
