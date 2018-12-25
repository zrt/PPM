package PPM

// object interface 存储物体
type Object interface {
	// interface Object
	Collide(*Ray) (bool, V)
	GetLocalColor() V
	GetMaterialWRWT() (float64, float64)
	GetRefRay(*Ray) *Ray
	GetTransRay(*Ray) *Ray
	VFrom(V) V
	GetName() string
}
