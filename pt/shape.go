package pt

import ()

type Shape interface {
	Compile()
	Box() Box
	Intersect(Ray) Hit
	Color(Vector) Color
	Material(Vector) Material
	Normal(Vector) Vector
	RandomPoint() Vector
}

type TransformedShape struct {
	Shape
	matrix  Matrix
	inverse Matrix
}

func NewTransformedShape(s Shape, m Matrix) Shape {
	return &TransformedShape{s, m, m.Inverse()}
}

func (s *TransformedShape) Box() Box {
	return s.matrix.MulBox(s.Shape.Box())
}

func (s *TransformedShape) Intersect(r Ray) Hit {
	hit := s.Shape.Intersect(s.inverse.MulRay(r))
	if !hit.Ok() {
		return hit
	}
	// if s.Shape is a Mesh, the hit.Shape will be a Triangle in the Mesh
	// we need to transform this Triangle, not the Mesh itself
	shape := &TransformedShape{hit.Shape, s.matrix, s.inverse}
	return Hit{shape, hit.T}
}

func (s *TransformedShape) Color(p Vector) Color {
	return s.Shape.Color(s.inverse.MulPosition(p))
}

func (s *TransformedShape) Material(p Vector) Material {
	return s.Shape.Material(s.inverse.MulPosition(p))
}

func (s *TransformedShape) Normal(p Vector) Vector {
	return s.matrix.MulDirection(s.Shape.Normal(s.inverse.MulPosition(p)))
}

func (s *TransformedShape) RandomPoint() Vector {
	return s.matrix.MulPosition(s.Shape.RandomPoint())
}
