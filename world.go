package PPM

type World struct {
	Objects []Object
	Lights  []V // 暂时仅支持点光源
	Tree    *OBJKDTree
}

func (w *World) Add(objs []Object) {
	w.Objects = append(w.Objects, objs...)
}
func (w *World) AddLight(objs []V) {
	w.Lights = append(w.Lights, objs...)
}

//func (w *World) Collide(r *Ray) (isCollide bool, dis float64, id int) {
//	isCollide = false
//	dis = 1e20
//	id = 0
//	for k, obj := range w.Objects {
//		ok, d := obj.Collide(r)
//		if ok {
//			isCollide = true
//			if d < dis {
//				dis = d
//				id = k
//			}
//		}
//	}
//	return
//}

func (w *World) Collide(r *Ray) (isCollide bool, dis float64, obj *Object, norm *V) {
	// kd tree
	return getCollide(w.Tree, r)
}
func getCollide(t *OBJKDTree, r *Ray) (isCollide bool, dis float64, obj *Object, norm *V) {
	center := t.GetCenter()
	v := center.Sub(r.Pos)
	p := r.Pos.Add(r.Dir.Mul(v.Dot(r.Dir)))
	dist := p.Disto(center)
	isCollide = false
	dis = 1e20
	obj = nil
	if dist > t.GetMaxBound() {
		return
	}
	ok := false
	d := -1.
	tmpnorm := &V{}
	center = (*t.Obj).GetCenter()
	v = center.Sub(r.Pos)
	p = r.Pos.Add(r.Dir.Mul(v.Dot(r.Dir)))
	dist = p.Disto(center)
	if dist < (*t.Obj).GetMax().Sub((*t.Obj).GetMin()).Maxf() {
		ok, d, tmpnorm = (*t.Obj).Collide(r)
	}
	if ok {
		isCollide = true
		if d < dis {
			dis = d
			obj = t.Obj
			norm = tmpnorm
		}
	}
	if t.Ls != nil {
		a, b, c, d := getCollide(t.Ls, r)
		if a {
			isCollide = true
			if b < dis {
				dis = b
				obj = c
				norm = d
			}
		}
	}
	if t.Rs != nil {
		a, b, c, d := getCollide(t.Rs, r)
		if a {
			isCollide = true
			if b < dis {
				dis = b
				obj = c
				norm = d
			}
		}
	}
	return
}
