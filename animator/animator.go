package animator

import (
	"time"

	"github.com/goserg/zg/zmath"
)

const defaultDuration = time.Second

type EaseFunc func(t float64) float64
type AnimaFunc[T any] func(start T, target T, dt float64) T

type Animator[T any] struct {
	duration time.Duration
	easeFunc EaseFunc

	isAnimating bool
	start       T
	target      T
	animTime    time.Duration

	f AnimaFunc[T]
}

type Options struct {
	Duration time.Duration
	EaseFunc EaseFunc
}

func New[T any](f AnimaFunc[T], opts *Options) *Animator[T] {
	a := Animator[T]{
		f:        f,
		duration: defaultDuration,
	}
	if opts != nil {
		a.duration = opts.Duration
		a.easeFunc = opts.EaseFunc
	}
	return &a
}

func (a *Animator[T]) Start(start T, target T) {
	a.isAnimating = true
	a.start = start
	a.target = target
	a.animTime = 0
}

func (a *Animator[T]) Animate(dt time.Duration) T {
	if a.animTime >= a.duration {
		a.isAnimating = false
		return a.target
	}
	a.animTime += dt

	t := zmath.Clamp(float64(a.animTime) / float64(a.duration))

	if a.easeFunc != nil {
		t = a.easeFunc(t)
	}

	return a.f(a.start, a.target, t)
}

func (a *Animator[T]) IsAnimating() bool {
	return a.isAnimating
}

func EaseIn(t float64) float64 {
	return t * t
}

func EaseOut(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

func CubicInOut(t float64) float64 {
	return t * t * (3 - 2*t)
}

// CustomEase - a и b - наклон графика в начале и в конце пути соотственно
func CustomEase(a float64, b float64) EaseFunc {
	return func(t float64) float64 {
		c3 := a + b - 2
		c2 := 3 - 2*a - b
		t2 := t * t
		t3 := t2 * t
		return c3*t3 + c2*t2 + a*t
	}
}
