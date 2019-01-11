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

	金属球 := Ball{V{27, 16.5, 47}, 16.5, 金属材质, "ball1_mirror"}
	漫反射球 := Ball{V{50, 8.5, 60}, 8.5, 漫反射材质, "ball2_diff"}
	玻璃球 := Ball{V{73, 16.5, 88}, 16.5, 玻璃材质, "ball2_glass"}
	单向三角面片 := NewTriangle(V{73, 40, 88}, V{50, 50, 60}, V{27, 16.5, 47}, DiffuseMaterial(蓝色), "triangle")

	场景.Add([]Object{金属球, 漫反射球, 玻璃球, 单向三角面片})

	半径 := 1e5
	wall1 := Ball{V{50, -1e5 + 81.6, 81.6}, 半径, DiffuseMaterial(灰色), "w_up"}
	wall2 := Ball{V{50, 1e5, 81.6}, 半径, DiffuseMaterial(灰色), "w_down"}
	wall3 := Ball{V{50, 40.8, -1e5 + 170}, 半径, DiffuseMaterial(V{}), "w_front"}
	wall4 := Ball{V{50, 40.8, 1e5}, 半径, DiffuseMaterial(灰色), "w_back"}
	wall5 := Ball{V{-1e5 + 99, 40.8, 81.6}, 半径, DiffuseMaterial(蓝色), "w_right"}
	wall6 := Ball{V{1e5 + 1, 40.8, 81.6}, 半径, DiffuseMaterial(红色), "w_left"}

	场景.Add([]Object{wall1, wall2, wall3, wall4, wall5, wall6})

	光源位置 := V{50, 60, 85}
	场景.AddLight([]V{光源位置})

	长宽 := 400
	相机 := Camera{V{45, 40, 200}, V{0, 0, -1}, 1, 1, 长宽, 长宽}

	光子映射 := NewPhotonMapping(场景, 相机, "test.png")
	光子映射.Render(1000)
}
