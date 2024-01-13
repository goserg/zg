package vector

func Dot(a Vector[float64], b Vector[float64]) float64 {
	return a.X*b.X + a.Y*b.Y
}
