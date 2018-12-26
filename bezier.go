package PPM

type Bezier struct {
	Point []V      // 位置
	Mat   Material // 材质
	Name  string   // 名称
}

// 是否相交，交点
func (b Bezier) Collide(r *Ray) (bool, float64) {
	return false, -1
}

func (b Bezier) GetMaterial() Material {
	return b.Mat
}

func (b Bezier) GetName() string {
	return b.Name
}

func (t Bezier) GetNormal(v V) V {

}

func NewBezier() *Bezier {
	return &Bezier{}
}
