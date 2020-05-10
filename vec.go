package v2d

import (
	"github.com/chewxy/math32"
)

// Add adds the two vectors.
func (v Vec) Add(w Vec) Vec {
	return Vec{X: v.X + w.X, Y: v.Y + w.Y}
}

// Sub subtracts w from v.
func (v Vec) Sub(w Vec) Vec {
	return Vec{X: v.X - w.X, Y: v.Y - w.Y}
}

// Neg returns -v.
func (v Vec) Neg() Vec {
	return Vec{X: -v.X, Y: -v.Y}
}

// Scale scales the vector with s.
func (v Vec) Scale(s float32) Vec {
	return Vec{X: v.X * s, Y: v.Y * s}
}

// Dot takes the dot-product of v and w.
func (v Vec) Dot(w Vec) float32 {
	return v.X*w.X + v.Y*w.Y
}

// Cross assumes that v, w are (x,y,0) vectors and returns the z component of 3D cross product only.
func (v Vec) Cross(w Vec) float32 {
	return v.X*w.Y - v.Y*w.X
}

// CrossZ returns the cross product of (x, y, 0) and (0, 0, z).
func (v Vec) CrossZ(z float32) Vec {
	return Vec{X: v.Y * z, Y: -v.X * z}
}

// CrossZ returns the cross product of (0, 0, z) and (x, y, 0).
func CrossZ(z float32, v Vec) Vec {
	return Vec{X: -v.Y * z, Y: v.X * z}
}

// LengthSquared returns the length of the vector squared.
func (v Vec) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Length returns the length of the vector.
func (v Vec) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Unit returns a unit vector in the same direction.
func (v Vec) Unit() Vec {
	// pre-scaling to not get underflow / overflow in
	// the sqrt in .length
	v = v.Scale(min(1/v.X, 1/v.Y))

	return v.Scale(1.0 / v.Length())
}

// Orth returns a vector w of same length as v that is orthogonal to v such that
// (v, w) are "right handed".
func (v Vec) Orth() (w Vec) {
	w.X = -v.Y
	w.Y = v.X
	return
}

// Eq returns true if v and w are equal withing tol.
func (v Vec) Eq(w Vec, tol float32) bool {
	return math32.Abs(v.X-w.X) <= tol && math32.Abs(v.Y-w.Y) <= tol
}

// IsZero checks if the vector is (0,0)
func (v Vec) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// Angle returns the angle of the vector (from x axis)
func (v Vec) Angle() float32 {
	return math32.Atan2(v.Y, v.X)
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
