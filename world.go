package PPM

type World struct {
	Objects []Object
	Lights  []V // 暂时仅支持点光源
}

func (w *World) Add(objs []Object) {
	w.Objects = append(w.Objects, objs...)
}
func (w *World) AddLight(objs []V) {
	w.Lights = append(w.Lights, objs...)
}

func (w *World) Collide(r *Ray) (isCollide bool, dis float64, id int) {
	isCollide = false
	dis = 1e20
	id = 0
	for k, obj := range w.Objects {
		ok, d := obj.Collide(r)
		if ok {
			isCollide = true
			if d < dis {
				dis = d
				id = k
			}
		}
	}
	return
}
