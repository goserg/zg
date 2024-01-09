package vector

// Lerp - линейная интерполяция между двух векторов
func Lerp[T Number](start Vector[T], target Vector[T], t float64) Vector[T] {
	difX := target.X - start.X
	difY := target.Y - start.Y

	return Vector[T]{
		X: T(float64(start.X) + float64(difX)*t),
		Y: T(float64(start.Y) + float64(difY)*t),
	}
}
