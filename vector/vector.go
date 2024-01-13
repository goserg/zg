package vector

import (
	"github.com/goserg/zg/angle"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector) Mul(m float64) Vector {
	return Vector{
		X: a.X * m,
		Y: a.Y * m,
	}
}

func (a Vector) Normalize() Vector {
	return a.Mul(1 / a.Len())
}

func (a Vector) Opposite() Vector {
	return Vector{
		X: -a.X,
		Y: -a.Y,
	}
}

func (a Vector) Rotate(r angle.Angle) Vector {
	return Vector{
		X: a.X*math.Cos(r.Rads()) - a.Y*math.Sin(r.Rads()),
		Y: a.X*math.Sin(r.Rads()) + a.Y*math.Cos(r.Rads()),
	}
}

func (a Vector) Len() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}
