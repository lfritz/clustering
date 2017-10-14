// Package clustering provides implementations of clustering algorithms for 2-D points.
package clustering

// Pre-defined cluster IDs. IDs for actual clusters start at 0.
const (
	Unclassified = -1 // points that are not (yet) classified
	Noise        = -2 // points that don't belong to any cluster
)
