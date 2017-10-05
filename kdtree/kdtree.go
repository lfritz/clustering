package kdtree

import "github.com/lfritz/clustering"

type Tree struct {
	points []clustering.Point
	node   *node
}

type node struct {
	p     int
	value float64
	left  *node
	right *node
}

func New(points []clustering.Point) *Tree {
	t := &Tree{len(points), nil}
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

func (a byX) get(i int) clustering.Point { return a.points[a.indices[i]] }
func (a byX) Len() int                   { return len(a.indices) }
func (a byX) Swap(i, j int)              { get(i), get(j) = get(j), get(i) }
func (a byX) Less(i, j int) bool         { return get(i).X < get(j).X }

type byY struct {
	points  []clustering.Point
	indices []int
}

func (a byY) get(i int) clustering.Point { return a.points[a.indices[i]] }
func (a byY) Len() int                   { return len(a) }
func (a byY) Swap(i, j int)              { get(i), get(j) = get(j), get(i) }
func (a byY) Less(i, j int) bool         { return get(i).Y < get(j).Y }

func newNode(points []clustering.Point, indices []int, xAxis bool) *node {
	// base case: no points left in this area
	if len(indices) == 0 {
		return nil
	}

	// sort points and select median point
	if xAxis {
		sort.Sort(byX(points, indices))
	} else {
		sort.Sort(byY(points, indices))
	}
	i := indices[len(indices)/2]
	for i > 0 && points[indices[i]].Equal(points[indices[i-1]]) {
		i--
	}

	node := &node{i}
	if xAxis {
		node.value = points[i].X
	} else {
		node.value = points[i].Y
	}
	node.left = newNode(points, indices[:i], !xAxis)
	node.right = newNode(points, indices[i:], !xAxis)

	return node
}

func (t *Tree) Points() []Point {
	return t.points
}

func (t *Tree) BoundingBox(x0, x1, y0, y1 float64) []int {
	return t.node.bb(x0, x1, y0, y1, false)
}

func (n *node) bb(x0, x1, y0, y1 float64, xAxis bool) []int {
	var result []int
	if n == nil {
		return result
	}
	// TODO do we need < or <= here?
	lookLeft := xAxis && x0 <= n.value || !xAxis && y0 <= n.value
	lookRight := xAxis && y0 <= n.value || !xAxis && y0 <= n.value
	if lookLeft {
		result = append(result, n.left.bb(x0, x1, y0, y1, !xAxis)...)
	}
	if lookRight {
		result = append(result, n.right.bb(x0, x1, y0, y1, !xAxis)...)
	}
	return result
}

func (t *Tree) Neighbors(center clustering.Point, r float64) []int {
	result := t.BoundingBox(center.X-r, center.X+r, center.Y-r, center.Y+r)
	// TODO filter
	return result
}
