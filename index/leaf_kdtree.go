package index

import (
	"github.com/lfritz/clustering/geometry"
)

// A LeafKDTree is a k-d tree for 2-D points that stores one point in each leaf node.
type LeafKDTree struct {
	points [][2]float64
	root   leafTreeNode
}

type leafTreeNode interface {
	boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
		level int) []int
}

type innerNode struct {
	value       float64
	left, right leafTreeNode
}

type leafNode struct {
	point int
}

// NewLeafKDTree creates a LeafKDTree for the given points.
func NewLeafKDTree(points [][2]float64) *LeafKDTree {
	indices := make([]int, len(points))
	for i := range indices {
		indices[i] = i
	}
	return &LeafKDTree{points, newLeafTreeNode(points, indices, 0)}
}

func newLeafTreeNode(points [][2]float64, indices []int, level int) leafTreeNode {
	if len(indices) == 0 {
		return (*innerNode)(nil)
	}
	if len(indices) == 1 {
		return leafNode(leafNode{indices[0]})
	}
	middle := len(indices) / 2
	dimension := level % 2
	QuickSelect(&byOneDimension{points, dimension, indices}, middle)
	value := points[indices[middle]][dimension]
	left := newLeafTreeNode(points, indices[:middle], level+1)
	right := newLeafTreeNode(points, indices[middle:], level+1)
	return &innerNode{value, left, right}
}

// Points returns the slice of points.
func (t *LeafKDTree) Points() [][2]float64 {
	return t.points
}

// BoundingBox returns the indices of all points within the given axis-aligned bounding box.
func (t *LeafKDTree) BoundingBox(bb *geometry.BoundingBox) []int {
	result := []int{}
	return t.root.boundingBox(result, t.points, bb, 0)
}

func (n *innerNode) boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
	level int) []int {
	if n == nil {
		return result
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

func (n leafNode) boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
	level int) []int {
	if bb.Contains(points[n.point]) {
		result = append(result, n.point)
	}
	return result
}
