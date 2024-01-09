package zmath

// Clamp - ограничевает число от 0 до 1
func Clamp(f float64) float64 {
	if f < 0 {
		return 0
	}
	if f > 1 {
		return 1
	}
	return f
}
