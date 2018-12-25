package PPM

type Triangle struct {
	Point [3]V     // 位置
	Mat   Material // 材质
	Name  string   // 名称
}

// 是否相交，交点
func (t Triangle) Collide(r *Ray) (bool, V) {
	return false, V{}
}

// 取得局部颜色(PT)
func (t Triangle) GetLocalColor() V {
	// todo
	return t.Mat.Color
}

// 取得材料透射、折射率
func (t Triangle) GetMaterialWRWT() (float64, float64) {
	return t.Mat.WR, t.Mat.WT
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
