package PPM

import "math"

type Triangle struct {
	Point  [3]V     // 位置
	Mat    Material // 材质
	Norm   V        //法向量
	Name   string   // 名称
	center V
}

// 是否相交，交点
func (t Triangle) Collide(r *Ray) (bool, float64, *V) {
	t0 := t.Norm.Dot(r.Dir)
	if t0 > 0 { // 单向面
		return false, -1, nil
	}
	e1 := t.Point[1].Sub(t.Point[0])
	e2 := t.Point[2].Sub(t.Point[0])
	T := r.Pos.Sub(t.Point[0])
	p := r.Dir.Cross(e2)
	q := T.Cross(e1)
	pe1 := p.Dot(e1)
	if math.Abs(pe1) < 1e-5 {
		return false, -1., nil
	}
	tt := q.Dot(e2) / pe1
	u := p.Dot(T) / pe1
	v := q.Dot(r.Dir) / pe1
	if u < 0 || v < 0 || u+v > 1 || tt < 0 {
		return false, -1., nil
	}
	return true, tt, &t.Norm
}

func (t Triangle) GetMaterial() Material {
	return t.Mat
}

func (t Triangle) GetName() string {
	return t.Name
}

func (t Triangle) GetCenter() V {
	return t.center
}

func (t Triangle) GetMin() V {
	return t.Point[0].Min(t.Point[1]).Min(t.Point[2])
}

func (t Triangle) GetMax() V {
	return t.Point[0].Max(t.Point[1]).Max(t.Point[2])
}

func NewTriangle(a, b, c V, mat Material, name string) *Triangle {
	v1 := b.Sub(a)
	v2 := c.Sub(a)
	norm := v1.Cross(v2).Norm()
	return &Triangle{[3]V{a, b, c}, mat, norm, name, a.Add(b).Add(c).Div(3)}
}
