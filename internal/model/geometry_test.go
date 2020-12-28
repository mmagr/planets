package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine(t *testing.T) {
	type entry struct {
		name     string
		p1       Point
		p2       Point
		expected Line
	}

	tests := []entry{
		entry{"positive slope on zero", Point{0, 0}, Point{1, 1}, Line{1, -1, 0}},
		entry{"positive slope not on zero", Point{0, 1}, Point{1, 2}, Line{1, -1, 1}},
		entry{"negative slope", Point{0, 1}, Point{1, 0}, Line{-1, -1, 1}},
		entry{"zero slope", Point{0, 1}, Point{1, 1}, Line{0, -1, 1}},
		entry{"x function", Point{3, 0}, Point{3, 1}, Line{-1, 0, 3}},
	}

	for _, k := range tests {
		t.Run(k.name, func(t *testing.T) {
			result := LineFromPoints(k.p1, k.p2)
			assert.NotNil(t, result, "Valid line was expected as result")
			assert.Equal(t, k.expected, *result, "Calculated line did not match expectations")
		})
	}
}

func TestLineDistance(t *testing.T) {
	type entry struct {
		name     string
		p        Point
		l        Line
		expected float64
	}

	tests := []entry{
		entry{"vertical", Point{0, 1}, Line{0, -1, 0}, 1},
		entry{"horizontal", Point{0, 0}, Line{-1, 0, 1}, 1},
	}

	for _, k := range tests {
		t.Run(k.name, func(t *testing.T) {
			result := k.l.Distance(k.p)
			assert.Equal(t, k.expected, result, "Calculated distance did not match expectations")
		})
	}
}

func TestLineIncludes(t *testing.T) {
	type entry struct {
		name     string
		p        Point
		l        Line
		expected bool
	}

	tests := []entry{
		entry{"vertical", Point{0, 3}, Line{0, -1, 3}, true},
		entry{"!vertical", Point{0, 2}, Line{0, -1, 3}, false},
		entry{"horizontal", Point{1, 0}, Line{-1, 0, 1}, true},
		entry{"!horizontal", Point{0, 0}, Line{-1, 0, 1}, false},
		entry{"angled", Point{1, 1}, Line{1, -1, 0}, true},
		entry{"!angled", Point{1, 0}, Line{1, -1, 0}, false},
	}

	for _, k := range tests {
		t.Run(k.name, func(t *testing.T) {
			result := k.l.Includes(k.p)
			assert.Equal(t, k.expected, result, "Point inclusion did not match")
		})
	}
}

func TestTriangleIncludesOrigin(t *testing.T) {
	type entry struct {
		name     string
		tri      *Triangle
		expected bool
	}

	tests := []entry{
		entry{"vertex", TriangleFromPoints(Point{0, 0}, Point{1, 1}, Point{-1, 1}), true},
		entry{"horizontal", TriangleFromPoints(Point{-1, 0}, Point{1, 0}, Point{0, 3}), true},
		entry{"vertical", TriangleFromPoints(Point{0, -1}, Point{0, 1}, Point{3, 0}), true},
		entry{"negative", TriangleFromPoints(Point{1, -1}, Point{1, 1}, Point{3, 0}), false},
		entry{"positive", TriangleFromPoints(Point{-1, -1}, Point{1, -1}, Point{0, 1}), true},
	}

	for _, k := range tests {
		t.Run(k.name, func(t *testing.T) {
			assert.Equal(t, k.expected, k.tri.Includes(Point{0, 0}), "Origin expectations did not match")
		})
	}
}

func TestLineAbove(t *testing.T) {
	type entry struct {
		name     string
		p        Point
		l        Line
		expected bool
		err      bool
	}

	tests := []entry{
		// y = 3
		entry{"horizontal", Point{0, 5}, Line{0, -1, 3}, true, false},
		entry{"!horizontal", Point{0, 2}, Line{0, -1, 3}, false, false},
		// x = 1
		entry{"vertical", Point{2, 0}, Line{-1, 0, 1}, false, true},
		entry{"!vertical", Point{0, 0}, Line{-1, 0, 1}, false, true},
		// y = x + 1
		entry{"angled", Point{0, 2}, Line{1, -1, 1}, true, false},
		entry{"!angled", Point{0, 0}, Line{1, -1, 1}, false, false},
	}

	for _, k := range tests {
		t.Run(k.name, func(t *testing.T) {
			result, err := k.l.Above(k.p)
			assert.Equal(t, k.err, err != nil, "Error indicator did not meet expectations")
			assert.Equal(t, k.expected, result, "Reference indicator did not match expected value")
		})
	}
}

func TestLineRight(t *testing.T) {
	type entry struct {
		name     string
		p        Point
		l        Line
		expected bool
		err      bool
	}

	tests := []entry{
		// y = 3
		entry{"horizontal", Point{0, 5}, Line{0, -1, 3}, false, true},
		entry{"!horizontal", Point{0, 2}, Line{0, -1, 3}, false, true},
		// x = 1
		entry{"vertical", Point{2, 0}, Line{-1, 0, 1}, true, false},
		entry{"!vertical", Point{0, 0}, Line{-1, 0, 1}, false, false},
		// y = x + 1
		entry{"angled", Point{0, 2}, Line{1, -1, 1}, false, false},
		entry{"!angled", Point{0, 0}, Line{1, -1, 1}, true, false},
	}

	for _, k := range tests {
		t.Run(k.name, func(t *testing.T) {
			result, err := k.l.Right(k.p)
			assert.Equal(t, k.err, err != nil, "Error indicator did not meet expectations")
			assert.Equal(t, k.expected, result, "Reference indicator did not match expected value")
		})
	}
}

func TestPerimeter(t *testing.T) {
	tri := TriangleFromPoints(Point{0, 0}, Point{0, 4}, Point{3, 4})
	assert.Equal(t, 12.0, tri.Perimeter())
}
