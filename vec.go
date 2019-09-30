// 2D vector math with a simple & powerful API. Less bugs & faster development.
package vec

import "math"

// Vec represents a 2-dimensional vector.
type Vec struct {
	X, Y float64
}

// Zero constant - to save some memory.
var Zero = Vec{0, 0}

// New returns a new vector with the given components.
func New(x, y float64) Vec {
	return Vec{X: x, Y: y}
}

// Add compoment-wise adds the vector `a` and vector `b`.
func (a Vec) Add(b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y}
}

// Len returns the length of vector `a`.
func (a Vec) Len() float64 {
	return math.Hypot(a.X, a.Y)
}

// Neg negates the components of vector `a`.
func (a Vec) Neg() Vec {
	return Vec{-a.X, -a.Y}
}

// Sub compoment-wise subtracts the vector `b` from vector `a`.
func (a Vec) Sub(b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y}
}

// Div compoment-wise divides the vector `a` by vector `b`.
func (a Vec) Div(b Vec) Vec {
	return Vec{a.X / b.X, a.Y / b.Y}
}

// Mul compoment-wise multiplies the vector `a` by vector `b`.
func (a Vec) Mul(b Vec) Vec {
	return Vec{a.X * b.X, a.Y * b.Y}
}

// Scale the vector `a` by a factor `b`.
func (a Vec) Scale(b float64) Vec {
	return Vec{a.X * b, a.Y * b}
}

// ScaleTo scales the vector `a` to the given `length`.
func (a Vec) ScaleTo(len float64) Vec {
	return a.Scale(len / a.Len())
}

// Equal checks wether the two vectors are equal.
func (a Vec) Equal(b Vec) bool {
	return a.X == b.X && a.Y == b.Y
}
