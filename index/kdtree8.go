package index

import (
	"github.com/lfritz/clustering/geometry"
)

// A KDTree8 is a k-d tree for 2-D points that stores 8 points in each leaf node.
type KDTree8 struct {
	points [][2]float64
	root   node8
}

type node8 interface {
	boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
		level int) []int
}

type inner8 struct {
	value       float64
	left, right node8
}

type leaf8 struct {
	count  int
	points [8]int
}

// NewKDTree8 creates a KDTree8 for the given points.
func NewKDTree8(points [][2]float64) *KDTree8 {
	indices := make([]int, len(points))
	for i := range indices {
		indices[i] = i
	}
	return &KDTree8{points, newNode8(points, indices, 0)}
}

func newNode8(points [][2]float64, indices []int, level int) node8 {
	count := len(indices)
	if count == 0 {
		return (*inner8)(nil)
	}
	if count <= 8 {
		result := leaf8{count: count}
		copy(result.points[:], indices)
		return result
	}
	middle := len(indices) / 2
	dimension := level % 2
	QuickSelect(&byOneDimension{points, dimension, indices}, middle)
	value := points[indices[middle]][dimension]
	left := newNode8(points, indices[:middle], level+1)
	right := newNode8(points, indices[middle:], level+1)
	return &inner8{value, left, right}
}

// Points returns the slice of points.
func (t *KDTree8) Points() [][2]float64 {
	return t.points
}

// BoundingBox returns the indices of all points within the given axis-aligned bounding box.
func (t *KDTree8) BoundingBox(bb *geometry.BoundingBox) []int {
	result := []int{}
	return t.root.boundingBox(result, t.points, bb, 0)
}

func (n *inner8) boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
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

func (n leaf8) boundingBox(result []int, points [][2]float64, bb *geometry.BoundingBox,
	level int) []int {
	for _, i := range n.points[:n.count] {
		if bb.Contains(points[i]) {
			result = append(result, i)
		}
	}
	return result
}
