package v2d

import (
	"math"
	"testing"
)

func TestTransformRectangle(t *testing.T) {
	r := R(V(1, 1), V(2, 2))
	tt := RotateMove(math.Pi/4, V(1.5, 1.5), V(-1.5, -1.5))
	rt := tt.TransRect(r)

	ref := R(V(-1.0/math.Sqrt2, -1.0/math.Sqrt2), V(1.0/math.Sqrt2, 1.0/math.Sqrt2))
	if !ref.Eq(rt, 1e-5) {
		t.Error("not equal", ref, rt)
	}
}
