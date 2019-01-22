package PPM

// object interface 存储物体
type Object interface {
	// interface Object
	Collide(*Ray) (bool, float64, *V)
	GetMaterial() Material
	GetCenter() V // 获取中心 (kdtree)
	GetMin() V    // 获取Min (kdtree)
	GetMax() V    // 获取Max (kdtree)
	GetName() string
}
