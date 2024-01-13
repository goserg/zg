package vector

import (
	"math"

	"github.com/goserg/zg/angle"
)

func Angle(a Vector, b Vector) angle.Angle {
	cosA := Dot(a, b) / (a.Len() * b.Len())
	A := math.Acos(cosA)
	return angle.FromRads(A)
}
