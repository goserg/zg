package collisions

import (
	"math"
	"time"

	"github.com/goserg/zg/vector"
)

type Rect struct {
	Pos  vector.Vector
	Size vector.Vector
}

func NewRect(pos vector.Vector, size vector.Vector) Rect {
	return Rect{
		Pos:  pos,
		Size: size,
	}
}

func PointVsRect(p vector.Vector, r Rect) bool {
	return p.X >= r.Pos.X &&
		p.Y >= r.Pos.Y &&
		p.X < r.Pos.X+r.Size.X &&
		p.Y < r.Pos.Y+r.Size.Y
}

func RectVsRect(r1 Rect, r2 Rect) bool {
	return r1.Pos.X < r2.Pos.X+r2.Size.X &&
		r1.Pos.X+r1.Size.X > r2.Pos.X &&
		r1.Pos.Y < r2.Pos.Y+r2.Size.Y &&
		r1.Pos.Y+r1.Pos.Y > r2.Pos.Y
}

func VectorVsRect(
	start vector.Vector,
	direction vector.Vector,
	rect Rect,
) (
	isColliding bool,
	contactPoint vector.Vector,
	contactNormal vector.Vector,
	tHitNear float64,
) {
	tNearX := (rect.Pos.X - start.X) / direction.X
	if rect.Pos.X == start.X {
		tNearX = 0
	}
	tNearY := (rect.Pos.Y - start.Y) / direction.Y
	if rect.Pos.Y == start.Y {
		tNearY = 0
	}
	tFarX := (rect.Pos.X + rect.Size.X - start.X) / direction.X
	tFarY := (rect.Pos.Y + rect.Size.Y - start.Y) / direction.Y

	if tNearX > tFarX {
		tNearX, tFarX = tFarX, tNearX
	}
	if tNearY > tFarY {
		tNearY, tFarY = tFarY, tNearY
	}
	if tNearX > tFarY || tNearY > tFarX {
		return false, vector.Vector{}, vector.Vector{}, 0
	}

	tHitNear = math.Max(tNearX, tNearY)
	tHitFar := math.Min(tFarX, tFarY)

	if tHitFar <= 0 {
		return false, vector.Vector{}, vector.Vector{}, 0
	}

	contactPoint.X = start.X + tHitNear*direction.X
	contactPoint.Y = start.Y + tHitNear*direction.Y

	if tNearX > tNearY {
		if direction.X < 0 {
			contactNormal = vector.Vector{X: 1}
		} else {
			contactNormal = vector.Vector{X: -1}
		}
	} else if tNearX < tNearY {
		if direction.Y < 0 {
			contactNormal = vector.Vector{Y: 1}
		} else {
			contactNormal = vector.Vector{Y: -1}
		}
	}

	return true, contactPoint, contactNormal, tHitNear
}

func DynamicRectVsRect(
	in Rect,
	speed vector.Vector,
	target Rect,
	dt time.Duration,
) (
	isColliding bool,
	contactPoint vector.Vector,
	contactNormal vector.Vector,
	contactTime float64,
) {
	if speed.X == 0 && speed.Y == 0 {
		return false, vector.Vector{}, vector.Vector{}, contactTime
	}
	expandedTarget := Rect{
		Pos:  target.Pos.Sub(in.Size.Mul(0.5)),
		Size: target.Size.Add(in.Size),
	}
	col, contactPoint, contactNormal, contactTime := VectorVsRect(
		in.Pos.Add(in.Size.Mul(0.5)),
		speed.Mul(dt.Seconds()),
		expandedTarget,
	)
	if col && contactTime < 1 {
		return true, contactPoint, contactNormal, contactTime
	}
	return false, vector.Vector{}, vector.Vector{}, contactTime
}
