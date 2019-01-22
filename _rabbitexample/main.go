package main

import . "github.com/zrt/PPM"
import (
	// "fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// https://en.wikipedia.org/wiki/Cornell_box
	// TestMat3d()
	// panic("")
	场景 := World{}
	白色 := V{0.99, 0.99, 0.99}
	灰色 := V{.75, .75, .75}
	蓝色 := V{.25, .25, .75}
	// 绿色 := V{.25, .75, .25}
	红色 := V{.75, .25, .25}

	金属材质 := MirrorMaterial(白色)
	// 漫反射材质 := DiffuseMaterial(白色)
	玻璃材质 := GlassMaterial(白色)

	金属球 := Ball{V{15, 8, -20}, 8, 金属材质, "金属球"}
	漫反射球 := Ball{V{55, 8.5, -20}, 8.5, DiffuseMaterial(V{.7, .44, 0}), "漫反射球"}
	玻璃球 := Ball{V{73, 15, -60}, 15, 玻璃材质, "玻璃球"}

	// // 单向三角面片 := NewTriangle(V{73, 40, 88}, V{50, 50, 60}, V{27, 16.5, 47}, 金属材质, "单向三角面片")

	场景.Add([]Object{金属球, 漫反射球, 玻璃球})
	// 场景.Add([]Object{玻璃球})
	// println(&玻璃球)
	points := []V{{0, 0, 0}, {-5, 2, 0}, {1, 3.4, 0}, {-1, 7.8, 0}}
	for i := 0; i < len(points); i++ {
		points[i] = points[i].Mul(7).Add(V{35, 0, -30})
	}
	贝塞尔 := NewBezier(points, 3, V{0, 1, 0}, V{35, 0, -37.5}, DiffuseMaterial(V{0.8, 0.79, 0.39}), "贝塞尔")

	// r := Ray{V{50, 40, 0}, V{0, 0, -1}}

	// a, b, c := 贝塞尔.Collide(&r)
	// fmt.Println(a, b, c)

	// panic("")
	// 场景.Add(贝塞尔.GenBalls(0.02, 0.3))
	场景.Add([]Object{贝塞尔})
	// (贝塞尔.GetMin()).Print()
	// (贝塞尔.GetMax()).Print()

	// 添加兔子

	是否添加兔兔 := false

	if 是否添加兔兔 {
		兔子色 := V{.87, .7, .51}
		f, err := os.Open("tmp.txt")
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		str := string(b)
		lines := strings.Split(str, "\n")
		if len(lines[len(lines)-1]) == 0 {
			lines = lines[:len(lines)-1]
		}
		// println(lines[len(lines)-1])
		rabbit := make([]Object, len(lines))
		tf := func(s string) float64 {
			ret, err := strconv.ParseFloat(s, 64)
			if err != nil {
				panic(err)
			}
			return ret
		}
		tr := func(v V) V {
			return v.Mul(140).Add(V{80, 0, -20})
		}
		for i := 0; i < len(lines); i++ {
			x := strings.Split(lines[i], " ")
			rabbit[i] = NewTriangle(tr(V{tf(x[0]), tf(x[1]), tf(x[2])}), tr(V{tf(x[3]), tf(x[4]), tf(x[5])}), tr(V{tf(x[6]), tf(x[7]), tf(x[8])}), DiffuseMaterial(兔子色), "兔兔")
			// fmt.Println("%#v\n", rabbit[i])
			// println(rabbit[i])
			// break
		}
		// println(rabbit)
		场景.Add(rabbit)
	}

	// 添加墙

	天花板1 := NewTriangle(V{0, 100, 0}, V{0, 100, -100}, V{100, 100, 0}, DiffuseMaterial(灰色), "天花板1")
	天花板2 := NewTriangle(V{100, 100, 0}, V{0, 100, -100}, V{100, 100, -100}, DiffuseMaterial(灰色), "天花板2")

	地板材质 := DiffuseMaterial(灰色)
	地板材质.SetTexture("tex.jpeg")
	地板1 := NewTriangle(V{0, 0, 0}, V{100, 0, 0}, V{0, 0, -100}, 地板材质, "地板1")
	地板2 := NewTriangle(V{100, 0, 0}, V{100, 0, -100}, V{0, 0, -100}, 地板材质, "地板2")

	前方的墙1 := NewTriangle(V{0, 0, -100}, V{100, 100, -100}, V{0, 100, -100}, DiffuseMaterial(灰色), "前方的墙1")
	前方的墙2 := NewTriangle(V{0, 0, -100}, V{100, 0, -100}, V{100, 100, -100}, DiffuseMaterial(灰色), "前方的墙2")

	身后的墙1 := NewTriangle(V{0, 0, 0}, V{0, 100, 0}, V{100, 100, 0}, DiffuseMaterial(灰色), "身后的墙1")
	身后的墙2 := NewTriangle(V{0, 0, 0}, V{100, 100, 0}, V{100, 0, 0}, DiffuseMaterial(灰色), "身后的墙2")

	左侧的墙1 := NewTriangle(V{0, 0, 0}, V{0, 100, -100}, V{0, 100, 0}, DiffuseMaterial(红色), "左侧的墙1")
	左侧的墙2 := NewTriangle(V{0, 0, 0}, V{0, 0, -100}, V{0, 100, -100}, DiffuseMaterial(红色), "左侧的墙2")

	右侧的墙1 := NewTriangle(V{100, 0, 0}, V{100, 100, 0}, V{100, 100, -100}, DiffuseMaterial(蓝色), "右侧的墙1")
	右侧的墙2 := NewTriangle(V{100, 0, 0}, V{100, 100, -100}, V{100, 0, -100}, DiffuseMaterial(蓝色), "右侧的墙2")

	场景.Add([]Object{天花板1, 天花板2, 地板1, 地板2, 前方的墙1, 前方的墙2, 身后的墙1, 身后的墙2, 右侧的墙1, 右侧的墙2, 左侧的墙1, 左侧的墙2})

	光源位置1 := V{50, 90, -55}
	// 光源位置2 := V{0, 90, -55}
	// 光源位置3 := V{100, 90, -55}
	场景.AddLight([]V{光源位置1})
	// 场景.AddLight([]V{光源位置1, 光源位置2, 光源位置3})

	长宽 := 1000
	相机 := Camera{V{50, 45, 90}, V{0, 0, -1}, 0.9, 0.9, 长宽, 长宽}

	光子映射 := NewPhotonMapping(&场景, &相机, "test.png", 2, 2500)
	光子映射.Render(1000000)
}
