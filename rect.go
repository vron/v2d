package v2d

// Eq returns true if r and q are equal withing tol.
func (r Rect) Eq(q Rect, tol float32) bool {
	return r.Min.Eq(q.Min, tol) && r.Max.Eq(q.Max, tol)
}
