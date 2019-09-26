package vector

import (
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func Length(v Vec3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func Unit(v Vec3) Vec3 {
	l := Length(v)
	return Vec3{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l,
	}
}

func Add(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Sub(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Dot(a, b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Scale(a Vec3, s float64) Vec3 {
	return Vec3{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}
