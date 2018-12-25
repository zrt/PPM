package PPM

type World struct {
	Objects []Object
	Lights  []Object
}

func (w *World) Add(objs []Object) error {
	w.Objects = append(w.Objects, objs...)
	return nil
}
func (w *World) AddLight(objs []Object) error {
	w.Lights = append(w.Lights, objs...)
	return nil
}

func (w *World) Collide(r *Ray) (miss, inshadow bool, wR, wT float64, localColor V, RefRay, TransRay *Ray) {
	miss = true
	inshadow = true
	RefRay = nil
	TransRay = nil
	localColor = V{}

	mndist := float64(1e20)
	var nearest Object
	for _, obj := range w.Objects {
		ok, pos := obj.Collide(r)
		if ok {
			miss = false
			dist := pos.Disto(r.Pos)
			if dist < mndist {
				mndist = dist
				nearest = obj
			}
			// println(i, dist)
		}
	}

	for _, obj := range w.Lights {
		ok, pos := obj.Collide(r)
		if ok {
			miss = false
			inshadow = false
			// localColor.Add_(obj.GetLocalColor())
			dist := pos.Disto(r.Pos)
			if dist < mndist {
				mndist = dist
				nearest = obj
			}
			// println(i, dist)
		} else if inshadow {
			// inshadow
			v := obj.VFrom(r.Pos)
			ok, pos := obj.Collide(&Ray{r.Pos, v})
			if !ok {
				// TODO FIXME
				// r.Pos.Print()
				// r.Dir.Print()
				// v.Print()
				// switch val := obj.(type) {
				// case Ball:
				// 	val.Pos.Print()
				// }
				// println(obj)
				// panic("error VFrom")

			} else {
				dist := pos.Disto(r.Pos)
				flag := true
				for _, obj2 := range w.Objects {
					ok, pos2 := obj2.Collide(&Ray{r.Pos, v})
					thisDis := pos2.Disto(r.Pos)
					if ok && thisDis > 1e-2 && thisDis < dist {
						flag = false
						break
					}
				}
				if flag {
					inshadow = false
					// localColor.Add_(obj.GetLocalColor())
				}
			}

		}
	}
	if !miss {
		localColor = nearest.GetLocalColor()
		// lastColor = nearest.GetLocalColor()
		wR, wT = nearest.GetMaterialWRWT()
		RefRay = nearest.GetRefRay(r)
		TransRay = nearest.GetTransRay(r)
		// println(nearest.GetName())
	}

	return
}
