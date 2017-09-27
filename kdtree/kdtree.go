package kdtree

import "github.com/lfritz/clustering"

type Tree struct {
	size int
	node *Node
}

type Node struct {
	p     int
	value float64
	left  *Node
	right *Node
}

func New(points []clustering.Point) *Tree {
	t := &Tree{len(points), nil}
	t.node = newNode(points)
	return t
}

type ByX []clustering.Point

func (a ByX) Len() int           { return len(a) }
func (a ByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool { return a[i].X < a[j].X }

type ByY []clustering.Point

func (a ByY) Len() int           { return len(a) }
func (a ByY) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByY) Less(i, j int) bool { return a[i].Y < a[j].Y }

func newNode(points []clustering.Point) *Node {
	if len(points) == 0 {
		return nil
	}
	// TODO
	return nil
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) BoundingBox(x0, x1, y0, y1 float64) []int {
	return t.node.bb(x0, x1, y0, y1, false)
}

func (n *Node) bb(x0, x1, y0, y1 float64, xAxis bool) []int {
	var result []int
	if n == nil {
		return result
	}
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
