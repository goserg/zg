package vector

// Lerp - линейная интерполяция между двух векторов
func Lerp(start Vector, target Vector, t float64) Vector {
	difX := target.X - start.X
	difY := target.Y - start.Y

	return Vector{
		X: start.X + difX*t,
		Y: start.Y + difY*t,
	}
}
