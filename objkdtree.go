package PPM

import "sort"

type OBJKDTree struct {
	Ls  *OBJKDTree
	Rs  *OBJKDTree
	mn  V
	mx  V
	Obj *Object
}

func BuildOBJKDTree(objs []Object, d int) *OBJKDTree {
	if len(objs) == 0 {
		return nil
	}
	cur := OBJKDTree{}
	// find d mid
	sort.Slice(objs, func(i, j int) bool { return objs[i].GetCenter().Val(d) < objs[j].GetCenter().Val(d) })
	mid := len(objs) / 2
	cur.Obj = &objs[mid]
	cur.mn = objs[mid].GetMin()
	cur.mx = objs[mid].GetMax()
	cur.Ls = BuildOBJKDTree(objs[:mid], (d+1)%3)
	cur.Rs = BuildOBJKDTree(objs[mid+1:], (d+1)%3)
	upd(&cur)
	return &cur
}

func upd(obj *OBJKDTree) {
	if obj.Ls != nil {
		obj.mn = obj.mn.Min(obj.Ls.mn)
		obj.mx = obj.mx.Max(obj.Ls.mx)
	}
	if obj.Rs != nil {
		obj.mn = obj.mn.Min(obj.Rs.mn)
		obj.mx = obj.mx.Max(obj.Rs.mx)
	}
}
func (t *OBJKDTree) GetCenter() V {
	return t.mx.Add(t.mn).Div(2)
}
func (t *OBJKDTree) GetMaxBound() float64 {
	return t.mx.Sub(t.mn).Maxf()
}
