package model

type Triangle struct {
	points []Point
	sides  []Line
}

func (t *Triangle) Valid() bool {
	if t == nil {
		return false
	}

	s1 := LineFromPoints(t.points[0], t.points[1])
	if s1 == nil {
		// we need three different points for a triangle
		return false
	}

	// the third point cannot be in the same line as the other two
	if s1.Distance(t.points[2]) == 0 {
		return false
	}

	s2 := LineFromPoints(t.points[0], t.points[2])
	s3 := LineFromPoints(t.points[1], t.points[2])
	if s2 == nil || s3 == nil {
		return false
	}

	return true
}

// checks if origin (0,0) is inside the triangle
func (t Triangle) Includes(p Point) bool {

	above := 0
	right := 0

	for _, side := range t.sides {
		// if point is in one of the sides, it is contained
		if side.Includes(p) {
			return true
		}

		if yAxis, err := side.Above(p); yAxis && err == nil {
			above += 1
		}

		if xAxis, err := side.Right(p); xAxis && err == nil {
			right += 1
		}
	}

	// The trick here is that we need at least one edge above, one below, and and to each side
	// This should work for triangles, but will not work for shapes with holes or convex ones
	return (above > 0) && (above < 3) && (right > 0) && (right < 3)
}

func (t Triangle) Perimeter() float64 {
	result := 0.0
	for i, _ := range t.points {
		result += t.points[i].Distance(t.points[(i+1)%len(t.points)])
	}
	return result
}

func TriangleFromPoints(p1 Point, p2 Point, p3 Point) *Triangle {
	s1 := LineFromPoints(p1, p2)
	if s1 == nil {
		// we need three different points for a triangle
		return nil
	}

	// the third point cannot be in the same line as the other two
	if s1.Distance(p3) == 0 {
		return nil
	}

	s2 := LineFromPoints(p1, p3)
	s3 := LineFromPoints(p2, p3)
	if s2 == nil || s3 == nil {
		return nil
	}

	return &Triangle{
		points: []Point{p1, p2, p3},
		sides:  []Line{*s1, *s2, *s3},
	}
}
