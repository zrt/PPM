package PPM

type Material struct {
	// Material interface
	Color  V
	WR, WT float64 // 反射率 透射？折射率?
}

func MirrorMaterial(color V) Material {
	return Material{color, 0.8, 0}
}

func DiffuseMaterial(color V) Material {
	return Material{color, 0.01, 0}
}

func LightMaterial(color V) Material {
	return Material{color, 0, 0}
}
