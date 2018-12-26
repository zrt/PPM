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
	"strings"
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
	sImg := image.NewRGBA(image.Rect(0, 0, img.X, img.Y))
	for i := 0; i < img.X; i++ {
		for j := 0; j < img.Y; j++ {
			c := img.At(i, j)
			sImg.Set(i, j, color.RGBA{toInt(c.X), toInt(c.Y), toInt(c.Z), 255})
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
		fmt.Sprint(runtime.NumCPU())+" cpus",
	)
	sImg := image.NewRGBA(image.Rect(0, 0, img.X, img.Y))
	for i := 0; i < img.X; i++ {
		for j := 0; j < img.Y; j++ {
			c := img.At(i, j)
			sImg.Set(i, j, color.RGBA{toInt(c.X), toInt(c.Y), toInt(c.Z), 255})
		}
	}

	addLabel(sImg, 5, 13*1, fmt.Sprintln(time.Now()))
	addLabel(sImg, 5, 13*2, fmt.Sprintln(sImg.Bounds().Size()))
	addLabel(sImg, 5, 13*3, info)
	addLabel(sImg, 5, 13*4, runtime.Version()+os.Getenv("GOPATH"))
	for i, c := range strings.Split(cmt, "\n") {
		addLabel(sImg, 5, 13*(5+i), c)
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
