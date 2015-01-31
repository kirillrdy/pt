package pt

import (
	"math"
	"math/rand"
)

type Triangle struct {
	v1, v2, v3 Vector
	color      Color
	material   Material
	normal     Vector
}

func NewTriangle(v1, v2, v3 Vector, color Color, material Material) Shape {
	normal := v1.Sub(v2).Cross(v2.Sub(v3))
	return &Triangle{v1, v2, v3, color, material, normal}
}

func (t *Triangle) Box() Box {
	min := t.v1.Min(t.v2).Min(t.v3)
	max := t.v1.Max(t.v2).Max(t.v3)
	return Box{min, max}
}

func (me *Triangle) Intersect(r Ray) float64 {
	e1 := me.v2.Sub(me.v1)
	e2 := me.v3.Sub(me.v1)
	h := r.Direction.Cross(e2)
	a := e1.Dot(h)
	if math.Abs(a) < EPS {
		return INF
	}
	f := 1 / a
	s := r.Origin.Sub(me.v1)
	u := f * s.Dot(h)
	if u < 0 || u > 1 {
		return INF
	}
	q := s.Cross(e1)
	v := f * r.Direction.Dot(q)
	if v < 0 || u + v > 1 {
		return INF
	}
	t := f * e2.Dot(q)
	if t > EPS {
		return t
	}
	return INF
}

func (t *Triangle) Color(v Vector) Color {
	return t.color
}

func (t *Triangle) Material(v Vector) Material {
	return t.material
}

func (t *Triangle) Normal(v Vector) Vector {
	return t.normal
}

func (t *Triangle) RandomPoint(rnd *rand.Rand) Vector {
	return Vector{} // TODO: fix
}