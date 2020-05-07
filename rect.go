package v2d

// Eq returns true if r and q are equal withing tol.
func (r Rect) Eq(q Rect, tol float32) bool {
	return r.Min.Eq(q.Min, tol) && r.Max.Eq(q.Max, tol)
}

// W returns the width (x dimension) of the rectangle.
func (r Rect) W() float32 {
	return r.Max.X - r.Min.X
}

// H returns the height (y dimension) of the rectangle.
func (r Rect) H() float32 {
	return r.Max.Y - r.Min.Y
}

// C0 returns the first corner of the rectangle.
func (r Rect) C0() Vec {
	return r.Min
}

// C1 returns the second corner of the rectangle.
func (r Rect) C1() Vec {
	r.Min.X = r.Max.X
	return r.Min
}

// C2 returns the third corner of the rectangle.
func (r Rect) C2() Vec {
	return r.Max
}

// C3 returns the forth corner of the rectangle.
func (r Rect) C3() Vec {
	r.Min.Y = r.Max.Y
	return r.Min
}
