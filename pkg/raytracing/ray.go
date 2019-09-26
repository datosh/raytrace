package raytracing

import (
	vector "github.com/datosh/raytrace/pkg/vector"
)

type Ray struct {
	A vector.Vec3
	B vector.Vec3
}

func NewRay(a, b vector.Vec3) Ray {
	return Ray{
		A: a,
		B: b,
	}
}

func (r *Ray) Origin() vector.Vec3 {
	return r.A
}

func (r *Ray) Direction() vector.Vec3 {
	return r.B
}

func (r *Ray) At(t float64) vector.Vec3 {
	return vector.Add(r.A, vector.Scale(r.B, t))
}
