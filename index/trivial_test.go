package index

import (
	"testing"
)

func TestTrivialPoints(t *testing.T) {
	i := NewTrivialIndex(points)
	testPoints(t, i)
}

func TestTrivialBoundingBox(t *testing.T) {
	i := NewTrivialIndex(points)
	testBoundingBox(t, i)
}

func TestTrivialCircle(t *testing.T) {
	i := NewTrivialIndex(points)
	testCircle(t, i)
}
