package PPM

import (
	"fmt"
	"math"
)

type V struct {
	X, Y, Z float64
}

func sqr(X float64) float64 {
	return X * X
}

func (v *V) Disto(v2 V) float64 {
	return math.Sqrt(sqr(v.X-v2.X) + sqr(v.Y-v2.Y) + sqr(v.Z-v2.Z))
}

func (v *V) Dot(v2 V) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v *V) Len() float64 {
	return math.Sqrt(sqr(v.X) + sqr(v.Y) + sqr(v.Z))
}

func (v V) Norm() V {
	len := v.Len()
	return V{v.X / len, v.Y / len, v.Z / len}
}

func (v *V) Print() {
	fmt.Printf("(%.2f, %.2f, %.2f)\n", v.X, v.Y, v.Z)
}

func (v V) Add(v2 V) V {
	return V{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v *V) Add_(v2 V) {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
}

func (v *V) Mulv(v2 *V) V {
	return V{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v V) Mul(k float64) V {
	return V{v.X * k, v.Y * k, v.Z * k}
}

func (v V) Div(k float64) V {
	return V{v.X / k, v.Y / k, v.Z / k}
}

func (v *V) Neg() V {
	return V{-v.X, -v.Y, -v.Z}
}

func (v *V) Sub(v2 V) V {
	return V{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func Grey(k float64) V {
	return V{k, k, k}
}

func NewV(x, y, z float64) V {
	return V{x, y, z}
}
