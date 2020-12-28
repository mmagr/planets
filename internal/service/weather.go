package service

import (
	"github.com/mmagr/planets/internal/model"
	"github.com/mmagr/planets/internal/model/conditions"
)

type Weather interface {
	ConditionsOn(day int) (string, float64)
}

// TODO get this from config
type climatempo struct {
	p1 model.Planet
	p2 model.Planet
	p3 model.Planet
	lf LineFactory
	pf PolygonFactory
}

func NewClimatempo(p1, p2, p3 model.Planet, lf LineFactory, pf PolygonFactory) Weather {
	return climatempo{p1: p1, p2: p2, p3: p3, lf: lf, pf: pf}
}

func (c climatempo) ConditionsOn(day int) (string, float64) {

	p1 := c.p1.Position(day)
	p2 := c.p2.Position(day)
	p3 := c.p3.Position(day)

	triangle := c.pf.FromPoints(p1, p2, p3)
	if triangle.Valid() == false {
		// we have either a draught, or perfect weather
		line := c.lf.FromPoints(p1, p2)
		if line.Includes(model.Point{0, 0}) {
			return conditions.Draught, 0.0
		}

		return conditions.Perfect, 0.0
	}

	if triangle.Includes(model.Point{0, 0}) {
		return conditions.Rain, triangle.Perimeter()
	}

	return conditions.Cloudy, triangle.Perimeter()
}
