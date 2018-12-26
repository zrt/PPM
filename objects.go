package PPM

// object interface 存储物体
type Object interface {
	// interface Object
	Collide(*Ray) (bool, float64)
	GetMaterial() Material
	GetNormal(V) V // 获得某点法向量
	GetName() string
}
