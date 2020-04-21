package v2d

import "github.com/chewxy/math32"

type Transform interface {
	// TransVec applies this Transform to a vector
	TransVec(Vec) Vec
	// TransRect applies this transform to a rectangle. Note that since rectangles always are axis
	// aligned the transformed rectangle will fully enclose the original rectangle.
	TransRect(Rect) Rect
}

type TransformRotateMove struct {
	Sin, Cos float32
	Around   Vec
	Offset   Vec
}

// Rotate move creates a Transform that rotates with an angle theta about around and
// subsequently adds offset.
func RotateMove(theta float32, around, offset Vec) (t TransformRotateMove) {
	t.Sin = math32.Sin(theta)
	t.Cos = math32.Cos(theta)
	t.Around = around
	t.Offset = offset
	return
}

// TransVec applies this Transform to a vector
func (t TransformRotateMove) TransVec(v Vec) (w Vec) {
	v = v.Sub(t.Around)
	w = V(v.X*t.Cos-v.Y*t.Sin, v.X*t.Sin+v.Y*t.Cos)
	w = w.Add(t.Around)
	w = w.Add(t.Offset)
	return
}

// TransRect applies this transform to a rectangle. Note that since rectangles always are axis
// aligned the transformed rectangle will fully enclose the original rectangle.
func (t TransformRotateMove) TransRect(r Rect) (tr Rect) {
	// simplest approach is to transform all 4 corners, and find min and
	// max x, y coordinates.
	// TODO: Since we know the angle we should be able to do something smarter? We know
	// which corner will be where - but do that and check to this function to be sure...

	c1 := r.Min
	c2 := V(r.Max.X, r.Min.Y)
	c3 := r.Max
	c4 := V(r.Min.X, r.Max.Y)

	c1 = t.TransVec(c1)
	c2 = t.TransVec(c2)
	c3 = t.TransVec(c3)
	c4 = t.TransVec(c4)

	tr.Min = V(math32.MaxFloat32, math32.MaxFloat32)
	tr.Max = V(-math32.MaxFloat32, -math32.MaxFloat32)

	// manually unrolled loops

	if c1.X < tr.Min.X {
		tr.Min.X = c1.X
	}
	if c2.X < tr.Min.X {
		tr.Min.X = c2.X
	}
	if c3.X < tr.Min.X {
		tr.Min.X = c3.X
	}
	if c4.X < tr.Min.X {
		tr.Min.X = c4.X
	}

	if c1.Y < tr.Min.Y {
		tr.Min.Y = c1.Y
	}
	if c2.Y < tr.Min.Y {
		tr.Min.Y = c2.Y
	}
	if c3.Y < tr.Min.Y {
		tr.Min.Y = c3.Y
	}
	if c4.Y < tr.Min.Y {
		tr.Min.Y = c4.Y
	}

	if c1.X > tr.Max.X {
		tr.Max.X = c1.X
	}
	if c2.X > tr.Max.X {
		tr.Max.X = c2.X
	}
	if c3.X > tr.Max.X {
		tr.Max.X = c3.X
	}
	if c4.X > tr.Max.X {
		tr.Max.X = c4.X
	}

	if c1.Y > tr.Max.Y {
		tr.Max.Y = c1.Y
	}
	if c2.Y > tr.Max.Y {
		tr.Max.Y = c2.Y
	}
	if c3.Y > tr.Max.Y {
		tr.Max.Y = c3.Y
	}
	if c4.Y > tr.Max.Y {
		tr.Max.Y = c4.Y
	}

	return
}
