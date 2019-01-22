package PPM

import "math"

type Ball struct {
	Pos  V        // 位置
	R    float64  // 半径
	Mat  Material // 材质
	Name string   // 名称
}

// 是否相交，距离
func (b Ball) Collide(r *Ray) (bool, float64, *V) {
	op := b.Pos.Sub(r.Pos)
	bb := op.Dot(r.Dir)
	det := bb*bb - op.Dot(op) + b.R*b.R
	if det < 0 {
		return false, -1., nil
	} else {
		det = math.Sqrt(det)
	}
	if bb-det > 1e-4 {
		x := r.Pos.Add(r.Dir.Mul(bb - det))
		n := x.Sub(b.Pos).Norm()
		return true, bb - det, &n
	} else if bb+det > 1e-4 {
		x := r.Pos.Add(r.Dir.Mul(bb + det))
		n := x.Sub(b.Pos).Norm()
		return true, bb + det, &n
	} else {
		return false, -1., nil
	}
}

func (b Ball) GetMaterial() Material {
	return b.Mat
}
func (b Ball) GetName() string {
	return b.Name
}

func (b Ball) GetCenter() V {
	return b.Pos
}
func (b Ball) GetMin() V {
	return b.Pos.Sub(V{b.R, b.R, b.R})
}
func (b Ball) GetMax() V {
	return b.Pos.Add(V{b.R, b.R, b.R})
}
