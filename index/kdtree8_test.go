package index

import (
	"testing"
)

func TestKDTree8Points(t *testing.T) {
	i := NewKDTree8(points)
	testPoints(t, i)
}

func TestKDTree8BoundingBox(t *testing.T) {
	i := NewKDTree8(points)
	testBoundingBox(t, i)
}

func TestKDTree8Circle(t *testing.T) {
	i := NewKDTree8(points)
	testCircle(t, i)
}
