package v2d

import "testing"

func TestNumStabUnit(t *testing.T) {
	v := Vec{6.617445e-24, 0}
	v = v.Unit()
	if v.Length() != 1.0 {
		var a = float32(6.617445e-24)
		b := a * (1.0 / a)
		t.Error(v, b)
	}
	if v.X != 1.0 || v.Y != 0.0 {
		t.Error(v)
	}
}
