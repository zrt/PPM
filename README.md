# PPM

[![GoDoc](https://godoc.org/github.com/zrt/PPM?status.svg)](https://godoc.org/github.com/zrt/PPM)
[![Go Report Card](https://goreportcard.com/badge/github.com/zrt/PPM)](https://goreportcard.com/report/github.com/zrt/PPM)
[![License](https://img.shields.io/badge/LICENSE-GLWTPL-green.svg)](https://github.com/zrt/PPM/blob/master/LICENSE)

A [Progressive Photon Mapping](https://www.ci.i.u-tokyo.ac.jp/~hachisuka/ppm.pdf) implementation in Go with:
- [x]  CPU-only
- [x] Uses all CPU cores in parallel
- [] Supports OBJ file
- [x] Supports sphere
- [x] Supports triangle
- [] Supports bezier surface
- [] Supports textures, bump maps and normal maps
- [] Polish API
- [] Using pprof to improve performance
- [] Supports SPPM(?)
- [] GPU mode
- [] Accelerated by kdtree
- [] Anti-aliasing


## Usage

```bash
go get -u github.com/zrt/PPM
```


```go
import "github.com/zrt/PPM"
````

## Example

```go
package main

import . "github.com/zrt/PPM"

func main() {
	// https://en.wikipedia.org/wiki/Cornell_box

	world := World{}
	v := *NewV(0.99, 0.99, 0.99)
	c1 := *NewV(.75, .75, .75)
	c2 := *NewV(.25, .25, .75)
	c3 := *NewV(.75, .25, .25)

	b1Material2 := MirrorMaterial(v)
	b2Material2 := DiffuseMaterial(v)
	b3Material2 := GlassMaterial(v)

	ball := Ball{V{27, 16.5, 47}, 16.5, b1Material2, "ball1_mirror"}
	ball2 := Ball{V{50, 8.5, 60}, 8.5, b2Material2, "ball2_diff"}
	ball3 := Ball{V{73, 16.5, 88}, 16.5, b3Material2, "ball2_glass"}
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
```

![example](https://github.com/zrt/PPM/raw/master/_history/t6.png)

## Links

- [smallppm.cpp](http://users-cs.au.dk/toshiya/smallppm_exp.cpp)

- [Progressive Photon Mapping](https://www.ci.i.u-tokyo.ac.jp/~hachisuka/ppm.pdf)

## License

This project is licensed under [GLWTPL](https://github.com/me-shaon/GLWTPL).

