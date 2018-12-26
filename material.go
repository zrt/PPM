package PPM

const (
	MT_DIFF = iota
	MT_SPEC
	MT_REFR
)

type Material struct {
	// Material interface
	Color V
	T     int // type 0 DIFF(Lambertian), 1 SPEC(mirror), 2 REFR(glass)
	//WR, WT float64 // 反射率 透射？折射率?
}

func MirrorMaterial(color V) Material {
	return Material{color, MT_SPEC}
}

func DiffuseMaterial(color V) Material {
	return Material{color, MT_DIFF}
}

func GlassMaterial(color V) Material {
	return Material{color, MT_REFR}
}
