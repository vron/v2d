package v2d

// Eq returns true if r and q are equal withing tol.
func (r Rect) Eq(q Rect, tol float64) bool {
	return r.Min.Eq(q.Min, tol) && r.Max.Eq(q.Max, tol)
}

// W returns the width (x dimension) of the rectangle.
func (r Rect) W() float64 {
	return r.Max.X - r.Min.X
}

// H returns the height (y dimension) of the rectangle.
func (r Rect) H() float64 {
	return r.Max.Y - r.Min.Y
}
