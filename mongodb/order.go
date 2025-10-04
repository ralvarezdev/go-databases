package mongodb

type (
	// Order represents the order of the index
	Order int
)

const (
	// Ascending order
	Ascending Order = 1

	// Descending order
	Descending Order = -1
)

// OrderInt converts the Order type to an integer
//
// Returns:
//
// An integer representing the order
func (o Order) OrderInt() int {
	return int(o)
}
