package PPM

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Image struct {
	Pixel []V
	X, Y  int
}

func (img *Image) Init(x, y int) {
	img.Pixel = make([]V, x*y)

	img.X = x
	img.Y = y
}

func (img *Image) At(x, y int) *V {
	//println(x, y)
	pos := x*img.Y + y
	//println(pos)
	return &img.Pixel[pos]
}

type PixelPos struct {
	Px, Py int
	Cx, Cy float64
	Weight float64
}

func (img *Image) AddPixelColor(x, y int, c V) {
	img.At(x, y).Add_(c)
}
func toInt(x float64) uint8 {
	return uint8(int(math.Pow(1-math.Exp(-x), 1/2.2)*255 + .5))
}
func (img *Image) Save(path string) {
	sImg := image.NewRGBA(image.Rect(0, 0, img.X, img.Y))
	for i := 0; i < img.X; i++ {
		for j := 0; j < img.Y; j++ {
			c := img.At(i, j)
			//c.Print()
			sImg.Set(i, j, color.RGBA{toInt(c.X), toInt(c.Y), toInt(c.Z), 255})
		}
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, sImg)
	if err != nil {
		panic(err)
	}
}

func (img *Image) SaveSmall(path string, ratio int) {
	sImg := image.NewRGBA(image.Rect(0, 0, img.X/ratio, img.Y/ratio))
	tmp := 1. / float64(ratio*ratio)
	for i := 0; i < img.X/ratio; i++ {
		for j := 0; j < img.Y/ratio; j++ {
			r, g, b := 0., 0., 0.
			for k1 := 0; k1 < ratio; k1++ {
				for k2 := 0; k2 < ratio; k2++ {
					c := img.At(i*ratio+k1, j*ratio+k2)
					r += tmp * c.X
					g += tmp * c.Y
					b += tmp * c.Z
				}
			}
			sImg.Set(i, j, color.RGBA{toInt(r), toInt(g), toInt(b), 255})
		}
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, sImg)
	if err != nil {
		panic(err)
	}
}
