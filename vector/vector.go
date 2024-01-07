package vector

import "math"

type Number interface {
	~int | ~float64
}

type Vector[T Number] struct {
	X T
	Y T
}

func (a Vector[T]) Add(b Vector[T]) Vector[T] {
	return Vector[T]{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector[T]) Sub(b Vector[T]) Vector[T] {
	return Vector[T]{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector[T]) Mul(m T) Vector[T] {
	return Vector[T]{
		X: a.X * m,
		Y: a.Y * m,
	}
}

func (a Vector[T]) Normalize() Vector[T] {
	return a.Mul(1 / T(a.Len()))
}

func (a Vector[T]) Opposite() Vector[T] {
	return Vector[T]{
		X: -a.X,
		Y: -a.Y,
	}
}

func (a Vector[T]) Len() float64 {
	return math.Sqrt(float64(a.X*a.X) + float64(a.Y*a.Y))
}