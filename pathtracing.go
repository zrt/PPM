package PPM

import (
	"fmt"
	"runtime"
)

type PathTracing struct {
	World    World
	Camera   Camera
	Sampler  Sampler
	SavePath string
}

type renderResult struct {
	Pos   PixelPos
	Color V
}

const TRACEDEPLIMIT = 10000
const WEIGHTEPS = 0.0001

func Trace(pt *PathTracing, r *Ray, dep int, weight float64) V {
	// println("pt")
	// r.Pos.Print()
	// r.Dir.Print()
	// println(weight)

	dep -= 1
	if dep < 0 || weight < WEIGHTEPS {
		// 反射次数过多，超过limit
		return V{}
	}
	background := V{} // 默认为黑色
	miss, inshadow, wR, wT, localColor, ReflectedRay, TransmittedRay := pt.World.Collide(r)
	if miss {
		return background
	}
	c := V{}
	if !inshadow {
		c = localColor.Mul(1 - wR - wT).Mul(weight)
	} else {
		// c = V{1, 0, 0}
	}
	if ReflectedRay != nil {
		// println(weight, wR)
		cR := Trace(pt, ReflectedRay, dep-1, weight*wR)
		c.Add_(cR)
	}
	if TransmittedRay != nil {
		// w_t
		// do nothing //
		println(wT)
	}

	return c
}

func PathTracingWorker(pt *PathTracing, p PixelPos, c chan renderResult) {
	ray := pt.Camera.Look(p.Cx, p.Cy)
	result := renderResult{p, Trace(pt, &ray, TRACEDEPLIMIT, 1.0)}
	c <- result
}

func (pt *PathTracing) Render() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumCPU())

	var img Image
	img.Init(pt.Sampler.Shape())
	lenPixelList := img.W * img.H
	pixelList := make(chan PixelPos, 32)
	go pt.Sampler.GenPixelList(pixelList)
	fmt.Println(lenPixelList, "pixels")

	c := make(chan renderResult, 32)

	go func() {
		for i := 0; i < lenPixelList; i++ {
			p, ok := <-pixelList
			if ok {
				go PathTracingWorker(pt, p, c)
			}
		}
		// close(c)
	}()

	bar := Pbar{}
	bar.Init(lenPixelList)

	cnt := 0
	for i := 0; i < lenPixelList; i++ {
		result, ok := <-c
		if !ok {
			fmt.Println("[!] channel error")
			break
		}
		cnt += 1
		img.AddPixelColor(result.Pos.Px, result.Pos.Py, result.Color.Mul(result.Pos.Weight))
		bar.Step(cnt)
	}
	img.Save(pt.SavePath)
}
