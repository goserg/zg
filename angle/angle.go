package angle

import "math"

// Angle - радианы
type Angle float64

func FromRads(r float64) Angle {
	return Angle(r)
}

func FromDegrees(d float64) Angle {
	return Angle(d * math.Pi / 180)
}

func (a Angle) Degrees() float64 {
	return float64(a) * 180 / math.Pi
}

func (a Angle) Rads() float64 {
	return float64(a)
}
