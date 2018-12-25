package PPM

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	Pixel [][]V
	W, H  int
}

func (img *Image) Init(w, h int) {
	var pix = make([][]V, h)
	for i := range pix {
		pix[i] = make([]V, w)
	}
	img.Pixel = pix
	img.W = w
	img.H = h
}

type PixelPos struct {
	Px, Py int
	Cx, Cy float64
	Weight float64
}

func (img *Image) AddPixelColor(x, y int, c V) {
	img.Pixel[x][y].Add_(c)
}

func (img *Image) Save(path string) {
	sImg := image.NewRGBA(image.Rect(0, 0, img.W, img.H))
	for i := 0; i < img.W; i++ {
		for j := 0; j < img.H; j++ {
			c := img.Pixel[i][j].Mul(255)
			sImg.Set(i, j, color.RGBA{uint8(c.X), uint8(c.Y), uint8(c.Z), 255})
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
