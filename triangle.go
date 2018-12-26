package PPM

type Triangle struct {
	Point [3]V     // 位置
	Mat   Material // 材质
	Name  string   // 名称
}

// 是否相交，交点
func (t Triangle) Collide(r *Ray) (bool, float64) {
	return false, -1.
}

// 取得局部颜色(PT)
func (t Triangle) GetLocalColor() V {
	// todo
	return t.Mat.Color
}

// 获得反射光线(在已经相交时)
func (t Triangle) GetRefRay(r *Ray) *Ray {
	panic("not implement")
	return &Ray{}
}

func (t Triangle) GetTransRay(*Ray) *Ray {
	return nil
}

func (t Triangle) VFrom(pos V) V {
	return V{}
}

func (t Triangle) GetName() string {
	return t.Name
}
