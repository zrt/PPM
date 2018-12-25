package main

import . "github.com/zrt/PPM"

func main() {

	// Cornell Box
	// https://en.wikipedia.org/wiki/Cornell_box

	world := World{}

	greyMaterial := DiffuseMaterial(Grey(0.7))

	redMaterial := DiffuseMaterial(V{0.5, 0, 0})

	greenMaterial := DiffuseMaterial(V{0, 0.5, 0})

	b2Material2 := MirrorMaterial(Grey(1))
	b2Material2.WR = 0.6
	// purpleMaterial := MirrorMaterial(V{116, 52, 129}.Div(255))
	// purpleMaterial.WR = 0.3
	ball := Ball{V{0, -1, 0}, 1, b2Material2, "ball1"}
	// ball2 := Ball{V{0, -0.5, -1}, 0.5, b2Material2, "ball2"}
	world.Add([]Object{ball /*ball2*/})

	wallR := float64(1000 - 3)

	wall1 := Ball{V{0, 1000, 0}, wallR, greyMaterial, "w_up"}
	wall2 := Ball{V{0, -1000, 0}, wallR, greyMaterial, "w_down"}
	wall3 := Ball{V{0, 0, 1000}, wallR, greyMaterial, "w_front"}
	wall4 := Ball{V{0, 0, -1000}, wallR, greyMaterial, "w_back"}
	wall5 := Ball{V{1000, 0, 0}, wallR, greenMaterial, "w_right"}
	wall6 := Ball{V{-1000, 0, 0}, wallR, redMaterial, "w_left"}

	world.Add([]Object{wall1, wall2, wall3, wall4, wall5, wall6})

	light := Ball{V{0, 3.2, -0.5}, 0.4, LightMaterial(Grey(1)), "L1"}
	light2 := Ball{V{0, 2, 3}, 0.5, LightMaterial(Grey(1)), "L2"}
	world.AddLight([]Object{light, light2})

	size := 500
	camera := Camera{V{0, 0, 3}, V{0, 0, -1}, 2, 2, size, size}

	// sampler := SamplerBalanced{size, size}

	// renderer := PathTracing{world, camera, sampler, "test.png"}
	// renderer.Render()
	pm := NewPhotonMapping(world, camera, "test.png")
	pm.Render()

}
