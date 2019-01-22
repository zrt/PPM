package main

import . "github.com/zrt/PPM"

func main() {
	// https://en.wikipedia.org/wiki/Cornell_box

	场景 := World{}
	白色 := V{0.99, 0.99, 0.99}
	灰色 := V{.75, .75, .75}
	蓝色 := V{.25, .25, .75}
	红色 := V{.75, .25, .25}

	金属材质 := MirrorMaterial(白色)
	漫反射材质 := DiffuseMaterial(白色)
	玻璃材质 := GlassMaterial(白色)

	金属球 := Ball{V{27, 16.5, 47}, 16.5, 金属材质, "金属球"}
	漫反射球 := Ball{V{50, 8.5, 60}, 8.5, 漫反射材质, "漫反射球"}
	玻璃球 := Ball{V{73, 16.5, 88}, 16.5, 玻璃材质, "玻璃球"}
	单向三角面片 := NewTriangle(V{73, 40, 88}, V{50, 50, 60}, V{27, 16.5, 47}, 金属材质, "单向三角面片")

	场景.Add([]Object{金属球, 漫反射球, 玻璃球, 单向三角面片})

	半径 := 1e5
	天花板 := Ball{V{50, -1e5 + 81.6, 81.6}, 半径, DiffuseMaterial(灰色), "天花板"}
	地板材质 := DiffuseMaterial(灰色)
	地板材质.SetTexture("tex.png")
	地板 := Ball{V{50, 1e5, 81.6}, 半径, 地板材质, "地板"}
	前方的墙 := Ball{V{50, 40.8, -1e5 + 170}, 半径, DiffuseMaterial(V{}), "前方的墙"}
	身后的墙 := Ball{V{50, 40.8, 1e5}, 半径, DiffuseMaterial(灰色), "身后的墙"}
	右侧的墙 := Ball{V{-1e5 + 99, 40.8, 81.6}, 半径, DiffuseMaterial(蓝色), "右侧的墙"}
	左侧的墙 := Ball{V{1e5 + 1, 40.8, 81.6}, 半径, DiffuseMaterial(红色), "左侧的墙"}

	场景.Add([]Object{天花板, 地板, 前方的墙, 身后的墙, 右侧的墙, 左侧的墙})

	// 光源位置1 := V{50, 60, 85}
	光源位置2 := V{5, 80, 80}
	场景.AddLight([]V{光源位置2})

	长宽 := 200
	相机 := Camera{V{45, 40, 200}, V{0, 0, -1}, 1, 1, 长宽, 长宽}

	光子映射 := NewPhotonMapping(&场景, &相机, "test.png", 1, 2500)
	光子映射.Render(1000)
}
