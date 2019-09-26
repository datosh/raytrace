package vector

type Vec3 struct {
	x, y, z float64
}

func Add(a, b Vec3) Vec3 {
	return Vec3{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
	}
}

// func Mul(a, b Vec3) Vec3 {
// 	return Vec3{
// 		x: a.x * b.x,
// 		y: a.y * b.y,
// 		z: a.z * b.z,
// 	}
// }

func Scale(a Vec3, s float64) Vec3 {
	return Vec3{
		x: a.x * s,
		y: a.y * s,
		z: a.z * s,
	}
}
