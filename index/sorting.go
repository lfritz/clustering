package index

// byOneDimension implements sort.Interface for sorting a slice of indices to points taking into
// account only one dimension.
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
