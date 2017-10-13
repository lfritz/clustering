package dbscan

// Pre-defined cluster IDs. The actual cluster IDs start at 2.
const (
	// TODO maybe these should be -1 and -2
	Unclassified = 0 // points that are not (yet) classified
	Noise        = 1 // points that don't belong to any cluster
)
