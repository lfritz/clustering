package index

import (
	"reflect"
	"sort"
	"testing"
)

// ints implement sort.Interface for a slice of ints
type ints []int

func (a ints) Len() int           { return len(a) }
func (a ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ints) Less(i, j int) bool { return a[i] < a[j] }

func TestPartition(t *testing.T) {
	cases := []struct {
		input              []int
		left, right, pivot int
		expectedOutput     []int
		expectedPivot      int
	}{
		{[]int{4, 6}, 0, 1, 0, []int{4, 6}, 0},
		{[]int{6, 4}, 0, 1, 0, []int{4, 6}, 1},
		{[]int{8, 3, 9, 4, 1, 7}, 0, 5, 4, []int{1, 3, 9, 4, 7, 8}, 0},
		{[]int{8, 3, 9, 4, 1, 7}, 0, 5, 2, []int{8, 3, 7, 4, 1, 9}, 5},
		{[]int{8, 3, 9, 4, 1, 7}, 0, 5, 3, []int{3, 1, 4, 7, 8, 9}, 2},
		{[]int{8, 3, 9, 4, 1, 7}, 1, 3, 3, []int{8, 3, 4, 9, 1, 7}, 2},
	}
	for _, c := range cases {
		output := make([]int, len(c.input))
		copy(output, c.input)
		got := partition(ints(output), c.left, c.right, c.pivot)
		ok := sameElements(c.input, output) &&
			output[got] == c.input[c.pivot] &&
			lessOrEqual(output[c.left:got], c.input[c.pivot]) &&
			greaterOrEqual(output[got:c.right+1], c.input[c.pivot]) &&
			reflect.DeepEqual(c.input[:c.left], output[:c.left]) &&
			reflect.DeepEqual(c.input[c.right+1:], output[c.right+1:])
		if !ok {
			t.Errorf("partition(%v, %v, %v, %v) returned %v, produced %v",
				c.input, c.left, c.right, c.pivot, got, output)
		}
	}
}

func lessOrEqual(xs []int, y int) bool {
	for _, x := range xs {
		if x > y {
			return false
		}
	}
	return true
}

func greaterOrEqual(xs []int, y int) bool {
	for _, x := range xs {
		if x < y {
			return false
		}
	}
	return true
}

func duplicate(s []int) []int {
	d := make([]int, len(s))
	copy(d, s)
	return d
}

func sameElements(slice1, slice2 []int) bool {
	sorted1 := duplicate(slice1)
	sorted2 := duplicate(slice2)
	sort.Sort(ints(sorted1))
	sort.Sort(ints(sorted2))
	return reflect.DeepEqual(sorted1, sorted2)
}

func TestQuickSelect(t *testing.T) {
	cases := []struct {
		input []int
		k     int
	}{
		{[]int{5}, 0},
		{[]int{8, 5}, 0},
		{[]int{3, 7, 2}, 1},
		{[]int{8, 3, 1, 5, 9, 7, 2}, 3},
	}
	for _, c := range cases {
		output := make([]int, len(c.input))
		copy(output, c.input)
		QuickSelect(ints(output), c.k)
		ok := sameElements(c.input, output) &&
			lessOrEqual(output[:c.k], output[c.k]) &&
			greaterOrEqual(output[c.k:], output[c.k])
		if !ok {
			t.Errorf("QuickSelect(%v, %v) produced %v", c.input, c.k, output)
		}
	}
}
