package service

import (
	"github.com/mmagr/planets/internal/model"
)

type Polygon interface {
	// sum of all sides
	Perimeter() float64

	// true if point is inside the polygon
	Includes(p model.Point) bool

	// true if polygon is valid
	Valid() bool
}

type PolygonFactory interface {
	FromPoints(points ...model.Point) Polygon
}

type TriangleFactory struct {}

func (tf TriangleFactory) FromPoints(points ...model.Point) Polygon {
	if len(points) != 3 {
		return nil
	}

	return model.TriangleFromPoints(points[0], points[1], points[2])
}

type Line interface {
	Includes(p model.Point) bool
}

type LineFactory interface {
	FromPoints(p1, p2 model.Point) Line
}

type ILineFactory struct {}

func (lf ILineFactory) FromPoints(p1, p2 model.Point) Line {
	return model.LineFromPoints(p1, p2)
}
