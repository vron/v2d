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
