package geometry

// A BoundingBox is an axis-aligned bounding box.
type BoundingBox struct {
	From, To [2]float64
}

// Contains returns true if the point is in the bounding box.
func (bb *BoundingBox) Contains(p [2]float64) bool {
	return bb.From[0] <= p[0] && p[0] < bb.To[0] &&
		bb.From[1] <= p[1] && p[1] < bb.To[1]
}
