package mongodb

// Order represents the order of the index
type Order int

const (
	// Ascending order
	Ascending Order = 1

	// Descending order
	Descending Order = -1
)

// OrderInt converts the Order type to an integer
func (o Order) OrderInt() int {
	return int(o)
}
