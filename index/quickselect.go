package index

import (
	"math/rand"
	"sort"
)

func partition(data sort.Interface, left, right, pivot int) int {
	data.Swap(pivot, right)
	storeIndex := left
	for i := left; i < right; i++ {
		if data.Less(i, right) {
			data.Swap(storeIndex, i)
			storeIndex++
		}
	}
	data.Swap(right, storeIndex)
	return storeIndex
}

// QuickSelect finds the kth smalles element in an unordered slice. It re-arranges the elements of
// data so that data[k] is the element that would be at position k if data were sorted, data[i] <=
// data[k] for i < k, and data[i] >= data[k] for i > k.
//
// It implements the Quickselect algorithm (https://en.m.wikipedia.org/wiki/Quickselect).
func QuickSelect(data sort.Interface, k int) {
	left := 0
	right := data.Len() - 1
	for right > left {
		pivot := left + rand.Intn(right-left+1)
		pivot = partition(data, left, right, pivot)
		if k == pivot {
			return
		}
		if k < pivot {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
}
