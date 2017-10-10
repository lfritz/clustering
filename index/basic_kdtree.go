package index

import (
	"github.com/lfritz/clustering/geometry"
)

// A BasicKDTree is a k-d tree for 2-D points that stores one point in each node.
type BasicKDTree struct {
	points [][2]float64
	root   *basicNode
}

type basicNode struct {
	point       int
	value       float64
	left, right *basicNode
}

// NewBasicKDTree creates a BasicKDTree for the given points.
func NewBasicKDTree(points [][2]float64) *BasicKDTree {
	indices := make([]int, len(points))
	for i := range indices {
		indices[i] = i
	}
	return &BasicKDTree{points, newBasicNode(points, indices, 0)}
}

func newBasicNode(points [][2]float64, indices []int, level int) *basicNode {
	if len(indices) == 0 {
		return nil
	}
	middle := len(indices) / 2
	dimension := level % 2
	QuickSelect(&byOneDimension{points, dimension, indices}, middle)
	point := indices[middle]
	value := points[indices[middle]][dimension]
	left := newBasicNode(points, indices[:middle], level+1)
	right := newBasicNode(points, indices[middle+1:], level+1)
	return &basicNode{point, value, left, right}
}

// Points returns the slice of points.
func (t *BasicKDTree) Points() [][2]float64 {
	return t.points
}

// BoundingBox returns the indices of all points within the given axis-aligned bounding box.
func (t *BasicKDTree) BoundingBox(bb *geometry.BoundingBox) []int {
	result := []int{}
	return t.root.boundingBox(result, t.points, bb, 0)
}

func (n *basicNode) boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
	level int) []int {
	if n == nil {
		return result
	}
	if bb.Contains(points[n.point]) {
		result = append(result, n.point)
	}
	dimension := level % 2
	if bb.From[dimension] <= n.value {
		result = n.left.boundingBox(result, points, bb, level+1)
	}
	if bb.To[dimension] >= n.value {
		result = n.right.boundingBox(result, points, bb, level+1)
	}
	return result
}
