package utils

type Vec3 [3]int

type Vec3i struct {
	X int
	Y int
	Z int
}

func NewVec3i(x int, y int, z int) *Vec3i {
	return &Vec3i{
		X: x,
		Y: y,
		Z: z,
	}
}

type Vec3f struct {
	X float32
	Y float32
	Z float32
}

func NewVec3f(x float32, y float32, z float32) Vec3f {
	return Vec3f{
		X: x,
		Y: y,
		Z: z,
	}
}
