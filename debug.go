package PPM

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"time"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func (img *Image) SaveWithComment(path string, x, y int, cmt string) {
	//println(img.W, img.H)
	sImg := image.NewRGBA(image.Rect(0, 0, img.W, img.H))
	for i := 0; i < img.W; i++ {
		for j := 0; j < img.H; j++ {
			c := img.Pixel[i][j].Mul(255)
			sImg.Set(i, j, color.RGBA{uint8(c.X), uint8(c.Y), uint8(c.Z), 255})
		}
	}

	addLabel(sImg, x, y, cmt)

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

func (img *Image) DebugSave(path string, cmt string) {
	info := fmt.Sprintln(
		runtime.GOOS,
		runtime.GOARCH,
		fmt.Sprint(runtime.NumCPU())+"cpus",
		runtime.Version(),
		runtime.GOROOT(),
	)
	sImg := image.NewRGBA(image.Rect(0, 0, img.W, img.H))
	for i := 0; i < img.W; i++ {
		for j := 0; j < img.H; j++ {
			c := img.Pixel[i][j].Mul(255)
			sImg.Set(i, j, color.RGBA{uint8(c.X), uint8(c.Y), uint8(c.Z), 255})
		}
	}

	addLabel(sImg, 5, 13, fmt.Sprintln(time.Now(), sImg.Bounds().Size()))
	addLabel(sImg, 5, 26, info)
	addLabel(sImg, 5, 39, cmt)

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
