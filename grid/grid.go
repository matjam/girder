package grid

// package grid implements a generic grid data structure with some
// useful methods for manipulating it.

type Grid[T any] struct {
	width, height int
	data          []T
}

// New creates a new grid with the given width and height of the given type.
func New[T any](width, height int) *Grid[T] {
	return &Grid[T]{
		width:  width,
		height: height,
		data:   make([]T, width*height),
	}
}

// Width returns the width of the grid.
func (g *Grid[T]) Width() int {
	return g.width
}

// Height returns the height of the grid.
func (g *Grid[T]) Height() int {
	return g.height
}

// Get returns the value at the given x and y coordinates. If the coordinates
// are out of bounds, the zero value of the type is returned along with false.
func (g *Grid[T]) Get(x, y int) (T, bool) {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		var t T
		return t, false
	}
	return g.data[y*g.width+x], true
}

// Set sets the value at the given x and y coordinates. If the coordinates
// are out of bounds, nothing is done.
func (g *Grid[T]) Set(x, y int, v T) {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		return
	}
	g.data[y*g.width+x] = v
}

// a matrix of the 8 locations of the neighbors of a given x and y coordinate
var neighborOffset = [8][2]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

// GetNeighbors returns the values of the 8 neighbors of the given x and y
// coordinates. If the coordinates are out of bounds, def is returned as the
// value for that neighbor.
//
// The neighbors are returned in the following order:
//
//	0 1 2
//	3 x 4
//	5 6 7
func (g *Grid[T]) GetNeighbors(x, y int, def T) [8]T {
	var neighbors [8]T
	var ok bool

	// go go gadget compiler and inline this sucker
	for i, loc := range neighborOffset {
		if neighbors[i], ok = g.Get(x + loc[0], y + loc[1]); !ok {
			neighbors[i] = def
		}
	}

	return neighbors
}

// GetBitmask takes a pointer to a function that accepts a value of the grid
// and returns a bool. This function is called for each neighbor of the given
// x and y coordinates. If the function returns true, the corresponding bit in
// the bitmask is set to 1. The bitmask is then returned.
//
// The neighbors are returned in the following order:
//
//	0 1 2
//	3 x 4
//	5 6 7
//
// the def parameter is used as the value for neighbors that are out of bounds
// of the grid.
func (g *Grid[T]) GetBitmask(x, y int, def T, f func(T) bool) uint8 {
	n := g.GetNeighbors(x, y, def)
	var bitmask uint8

	for i := range n {
		if f(n[i]) {
			bitmask |= (1 << uint(i))
		}
	}

	return bitmask
}
