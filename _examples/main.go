package main

import . "github.com/zrt/PPM"

func main() {
	// https://en.wikipedia.org/wiki/Cornell_box

	world := World{}
	v := V{0.99, 0.99, 0.99}
	c1 := V{.75, .75, .75}
	c2 := V{.25, .25, .75}
	c3 := V{.75, .25, .25}

	mm := MirrorMaterial(v)
	dm := DiffuseMaterial(v)
	gm := GlassMaterial(v)

	ball := Ball{V{27, 16.5, 47}, 16.5, mm, "ball1_mirror"}
	ball2 := Ball{V{50, 8.5, 60}, 8.5, dm, "ball2_diff"}
	ball3 := Ball{V{73, 16.5, 88}, 16.5, gm, "ball2_glass"}
	tri := NewTriangle(V{73, 40, 88}, V{50, 50, 60}, V{27, 16.5, 47}, DiffuseMaterial(c2), "triangle")

	world.Add([]Object{ball, ball2, ball3, tri})

	R := 1e5
	wall1 := Ball{V{50, -1e5 + 81.6, 81.6}, R, DiffuseMaterial(c1), "w_up"}
	wall2 := Ball{V{50, 1e5, 81.6}, R, DiffuseMaterial(c1), "w_down"}
	wall3 := Ball{V{50, 40.8, -1e5 + 170}, R, DiffuseMaterial(V{}), "w_front"}
	wall4 := Ball{V{50, 40.8, 1e5}, R, DiffuseMaterial(c1), "w_back"}
	wall5 := Ball{V{-1e5 + 99, 40.8, 81.6}, R, DiffuseMaterial(c2), "w_right"}
	wall6 := Ball{V{1e5 + 1, 40.8, 81.6}, R, DiffuseMaterial(c3), "w_left"}

	world.Add([]Object{wall1, wall2, wall3, wall4, wall5, wall6})

	light := V{50, 60, 85}
	world.AddLight([]V{light})

	size := 400
	camera := Camera{V{45, 40, 200}, V{0, 0, -1}, 1, 1, size, size}

	pm := NewPhotonMapping(world, camera, "test.png")
	pm.Render(5000)
}
