package vector

func Dot(a Vector, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}
