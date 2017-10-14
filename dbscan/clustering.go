package dbscan

// Pre-defined cluster IDs. IDs for actual clusters start at 0.
const (
	Unclassified = -1 // points that are not (yet) classified
	Noise        = -2 // points that don't belong to any cluster
)
