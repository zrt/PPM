package PPM

import "math"

type Ball struct {
	Pos  V        // 位置
	R    float64  // 半径
	Mat  Material // 材质
	Name string   // 名称
}

// 是否相交，距离
func (b Ball) Collide(r *Ray) (bool, float64) {
	op := b.Pos.Sub(r.Pos)
	bb := op.Dot(r.Dir)
	det := bb*bb - op.Dot(op) + b.R*b.R
	if det < 0 {
		return false, -1.
	} else {
		det = math.Sqrt(det)
	}
	if bb-det > 1e-4 {
		return true, bb - det
	} else if bb+det > 1e-4 {
		return true, bb + det
	} else {
		return false, -1.
	}
}

func (b Ball) GetMaterial() Material {
	return b.Mat
}
func (b Ball) GetNormal(x V) V {
	return x.Sub(b.Pos).Norm()
}

func (b Ball) GetName() string {
	return b.Name
}
