package grid_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matjam/girder/grid"
)

func debugPrint(g *grid.Grid[val]) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			v, _ := g.Get(x, y)
			print(v.V)
		}
		println()
	}
}

type val struct {
	X, Y, V int
}

var val_zero = val{0, 0, 0}
var val_one = val{1, 1, 1}
var val_two = val{2, 2, 2}

func TestGrid(t *testing.T) {
	g := grid.New[val](3, 3)
	if g.Width() != 3 || g.Height() != 3 {
		t.Errorf("expected width and height to be 3, got %d and %d", g.Width(), g.Height())
	}

	g.Set(1, 1, val_one)
	g.Set(2, 2, val_two)

	v, ok := g.Get(1, 1)
	if !ok || !cmp.Equal(v, val_one) {
		t.Errorf("expected to get %v at 1,1, got %v", val_one, v)
	}

	v, ok = g.Get(2, 2)
	if !ok || !cmp.Equal(v, val_two) {
		t.Errorf("expected to get 2 at 2,2, got %d", v)
	}

	v, ok = g.Get(0, 0)
	if !ok || !cmp.Equal(v, val_zero) {
		t.Errorf("expected to get 0 at 0,0, got %d", v)
	}

	v, ok = g.Get(3, 3)
	if ok || !cmp.Equal(v, val_zero) {
		t.Errorf("expected to get 0 at 3,3, got %d", v)
		debugPrint(g)
	}

	g.Set(-1, -1, val_two)
	v, ok = g.Get(-1, -1)
	if ok || !cmp.Equal(v, val_zero) {
		t.Errorf("expected to get 0 at -1,-1, got %d", v)
		debugPrint(g)
	}

	neighbors := g.GetNeighbors(1, 1, val_zero)
	expected := [8]val{val_zero, val_zero, val_zero, val_zero, val_zero, val_zero, val_zero, val_two}
	if !cmp.Equal(neighbors, expected) {
		t.Errorf("expected neighbors to be %v, got %v", expected, neighbors)
		debugPrint(g)
	}

	neighbors = g.GetNeighbors(2, 2, val_zero)
	expected = [8]val{val_one, val_zero, val_zero, val_zero, val_zero, val_zero, val_zero, val_zero}
	if !cmp.Equal(neighbors, expected) {
		t.Errorf("expected neighbors to be %v, got %v", expected, neighbors)
		debugPrint(g)
	}

}

func TestGetBitmask(t *testing.T) {
	g := grid.New[val](3, 3)
	isSet := func(v val) bool {
		return v.V != 0
	}

	if bitmask := g.GetBitmask(1, 1, val_zero, isSet); bitmask != 0 {
		t.Errorf("expected bitmask to be %08b, got %08b", 0, bitmask)
		debugPrint(g)
	}

	// set a couple of values
	g.Set(2, 0, val_one)
	g.Set(2, 2, val_one)

	if bitmask := g.GetBitmask(1, 1, val_zero, isSet); bitmask != 132 {
		t.Errorf("expected bitmask to be %8b, got %8b", 132, bitmask)
		debugPrint(g)
	}

	// set all values
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			g.Set(x, y, val_one)
		}
	}

	if bitmask := g.GetBitmask(1, 1, val_zero, isSet); bitmask != 255 {
		t.Errorf("expected bitmask to be %8b, got %8b", 255, bitmask)
		debugPrint(g)
	}
}
