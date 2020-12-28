package model

import (
	"math"
)

// Describe a planet with a circular orbit around point (0,0)
type Planet struct {
	// Angular speed, in degrees
	Omega float64
	// Initial angle on d = 0, in degrees
	Alpha float64
	// Radius of the circular orbit around (0,0)
	Orbit float64
}

// Given a day, return the cartesian position of the planet, at that time
func (p Planet) Position(day int) Point {
	// we're actually subject to some interesting errors here - round is used so we can at least
	// observe the draught periods more easily
	y, x := math.Sincos(p.Angle(day) * (math.Pi / 180.0))
	return Point{
		X: math.Round(p.Orbit * x),
		Y: math.Round(p.Orbit * y),
	}
}

func (p Planet) Angle(day int) float64 {
	return p.Alpha + (float64(day) * p.Omega)
}
