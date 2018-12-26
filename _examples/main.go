package main

import . "github.com/zrt/PPM"

func main() {

	// Cornell Box
	// https://en.wikipedia.org/wiki/Cornell_box

	world := World{}

	// greyMaterial := DiffuseMaterial(*NewV(0.7, 0.7, 0.7))

	// redMaterial := DiffuseMaterial(V{0.5, 0, 0})

	// greenMaterial := DiffuseMaterial(V{0, 0.5, 0})

	b1Material2 := MirrorMaterial(*NewV(0.99, 0.99, 0.99))
	b2Material2 := DiffuseMaterial(*NewV(0.99, 0.99, 0.99))
	b3Material2 := GlassMaterial(*NewV(0.99, 0.99, 0.99))
	// b2Material2.WR = 0.6
	// purpleMaterial := MirrorMaterial(V{116, 52, 129}.Div(255))
	// purpleMaterial.WR = 0.3
	ball := Ball{V{27, 16.5, 47}, 16.5, b1Material2, "ball1_mirror"}
	ball2 := Ball{V{50, 8.5, 60}, 8.5, b2Material2, "ball2_diff"}
	ball3 := Ball{V{73, 16.5, 88}, 16.5, b3Material2, "ball2_glass"}
	world.Add([]Object{ball, ball2, ball3})

	wallR := 1e5

	wall1 := Ball{V{50, -1e5 + 81.6, 81.6}, wallR, DiffuseMaterial(*NewV(.75, .75, .75)), "w_up"}
	wall2 := Ball{V{50, 1e5, 81.6}, wallR, DiffuseMaterial(*NewV(.75, .75, .75)), "w_down"}
	wall3 := Ball{V{50, 40.8, -1e5 + 170}, wallR, DiffuseMaterial(V{}), "w_front"}
	wall4 := Ball{V{50, 40.8, 1e5}, wallR, DiffuseMaterial(*NewV(.75, .75, .75)), "w_back"}
	wall5 := Ball{V{-1e5 + 99, 40.8, 81.6}, wallR, DiffuseMaterial(*NewV(.25, .25, .75)), "w_right"}
	wall6 := Ball{V{1e5 + 1, 40.8, 81.6}, wallR, DiffuseMaterial(*NewV(.75, .25, .25)), "w_left"}

	world.Add([]Object{wall1, wall2, wall3, wall4, wall5, wall6})

	light := V{50, 60, 85}
	world.AddLight([]V{light})

	size := 500
	// camera := Camera{V{50, 48 - 3, 295.6}, V{0, 0, -1}, 0.5, 0.5, size, size}
	camera := Camera{V{45, 50, 165}, V{0, 0, -1}, 1.3, 1.3, size, size}
	// sampler := SamplerBalanced{size, size}

	// renderer := PathTracing{world, camera, sampler, "test.png"}
	// renderer.Render()
	pm := NewPhotonMapping(world, camera, "test.png")
	pm.Render(100000)

}
