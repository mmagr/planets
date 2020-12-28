package model

import (
	"errors"
	"math"
)

// Represents a line of form ax + by + c = 0
type Line struct {
	A float64
	B float64
	C float64
}

func (l Line) Distance(p Point) float64 {
	dem := math.Sqrt(math.Pow(l.A, 2) + math.Pow(l.B, 2))
	nom := math.Abs((l.A * p.X) + (l.B * p.Y) + l.C)
	return nom / dem
}

// True if point is in line
func (l Line) Includes(p Point) bool {
	return (l.A*p.X)+(l.B*p.Y)+l.C == 0
}

func (l Line) y(x float64) float64 {
	return ((l.A * x) + l.C) / (-1 * l.B)
}

func (l Line) x(y float64) float64 {
	return ((l.B * y) + l.C) / (-1 * l.A)
}

// True if point is above line
func (l Line) Above(p Point) (bool, error) {
	if l.B == 0 {
		// it doesn't make sense for anything to be above a vertical line on the plane
		return false, errors.New("invalid assessment for line")
	}

	return l.y(p.X) < p.Y, nil
}

// True if point is to the right of line
func (l Line) Right(p Point) (bool, error) {
	if l.A == 0 {
		// it doesn't make sense for anything to be to the right of a horizontal line on the plane
		return false, errors.New("invalid assessment for line")
	}

	return l.x(p.Y) < p.X, nil
}

func LineFromPoints(p1 Point, p2 Point) *Line {
	if p1.X == p2.X {
		if p1.Y == p2.Y {
			// two points on the same coordinates do not produce a line
			return nil
		}

		return &Line{-1, 0, p1.X}
	}

	result := Line{
		A: (p1.Y - p2.Y) / (p1.X - p2.X),
		B: -1,
	}

	result.C = p1.Y - (result.A * p1.X)
	return &result
}
