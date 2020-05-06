// Pacakge v2d implements 2D float64 vector math
package v2d

import "math"

// Vec is a 2D vector (x, y)
type Vec struct {
	X, Y float64
}

// V is a convenience function to create a Vec.
func V(x, y float64) Vec {
	return Vec{X: x, Y: y}
}

// Split is a convenience function to get the components of a Vec.
func (v Vec) Split() (float64, float64) {
	return v.X, v.Y
}

// A Rect represents a bounding rectangle in the (x, y) plane that is
// always aligned with the x and y axis.
type Rect struct {
	Min, Max Vec
}

// R is a convenience method to create a rectangle
func R(c1, c2 Vec) (tr Rect) {
	tr.Min = V(math.MaxFloat64, math.MaxFloat64)
	tr.Max = V(-math.MaxFloat64, -math.MaxFloat64)

	if c1.X < tr.Min.X {
		tr.Min.X = c1.X
	}
	if c2.X < tr.Min.X {
		tr.Min.X = c2.X
	}

	if c1.Y < tr.Min.Y {
		tr.Min.Y = c1.Y
	}
	if c2.Y < tr.Min.Y {
		tr.Min.Y = c2.Y
	}

	if c1.X > tr.Max.X {
		tr.Max.X = c1.X
	}
	if c2.X > tr.Max.X {
		tr.Max.X = c2.X
	}

	if c1.Y > tr.Max.Y {
		tr.Max.Y = c1.Y
	}
	if c2.Y > tr.Max.Y {
		tr.Max.Y = c2.Y
	}

	return
}
