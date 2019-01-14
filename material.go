package PPM

import (
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"strings"
)

const (
	MT_DIFF = iota
	MT_SPEC
	MT_REFR
)

type Material struct {
	// Material interface
	color V
	T     int // type 0 DIFF(Lambertian), 1 SPEC(mirror), 2 REFR(glass)
	//WR, WT float64 // 反射率 透射？折射率?
	Texture *image.Image
}

func MirrorMaterial(color V) Material {
	return Material{color, MT_SPEC, nil}
}

func DiffuseMaterial(color V) Material {
	return Material{color, MT_DIFF, nil}
}

func GlassMaterial(color V) Material {
	return Material{color, MT_REFR, nil}
}

func (m *Material) SetTexture(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var img image.Image
	if strings.HasSuffix(filename, "png") {
		img, err = png.Decode(f)
	} else {
		img, err = jpeg.Decode(f)
	}
	if err != nil {
		panic(err)
	}
	m.Texture = &img
}

func (m Material) Color(v V) V {
	if m.Texture == nil {
		return m.color
	} else {
		dx := float64((*m.Texture).Bounds().Dx())
		dy := float64((*m.Texture).Bounds().Dy())
		const k = 10
		u := math.Mod(v.X*k, dx)
		if u < 0 {
			u += dx
		}
		v := math.Mod(v.Z*k, dy)
		if v < 0 {
			v += dy
		}
		//if int(u) != 0 {
		//	println(int(dx), int(dy), int(u), int(v), k)
		//
		//}
		r, g, b, _ := (*m.Texture).At(int(u), int(v)).RGBA()
		//println(r, g, b)
		return *NewV(float64(r)/65535., float64(g)/65535., float64(b)/65535.)
	}
}
