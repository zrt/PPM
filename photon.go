package PPM

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const ALPHA = 0.7

type PhotonMapping struct {
	world     World
	camera    Camera
	path      string
	img       Image
	hitpoints *list.List
}

type HPoint struct { // hit point
	f, pos, nrm, flux V
	r2                float64
	n                 uint // n = N / ALPHA in the paper
	pix               int
}

func NewPhotonMapping(world World, camera Camera, path string) *PhotonMapping {
	pm := &PhotonMapping{world, camera, path, Image{}, list.New()}
	pm.img.Init(camera.X, camera.Y)
	return pm
}

func (pm *PhotonMapping) Pass1() {

	bar := Pbar{}
	bar.Init(pm.camera.X * pm.camera.Y)

	for i := 0; i < pm.camera.X; i++ {
		for j := 0; j < pm.camera.Y; j++ {
			pm.Trace(pm.camera.Look(i, j), 0, false, *NewV(0, 0, 0), *NewV(1, 1, 1), i*pm.camera.Y+j)
			bar.Tick()
		}
	}
	fmt.Println()

	// build_hash_grid

	mx := V{-1e20, -1e20, -1e20}
	mn := V{1e20, 1e20, 1e20}
	for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*HPoint)
		mn.X = math.Min(mn.X, hp.pos.X)
		mn.Y = math.Min(mn.Y, hp.pos.Y)
		mn.Z = math.Min(mn.Z, hp.pos.Z)
		mx.X = math.Max(mx.X, hp.pos.X)
		mx.Y = math.Max(mx.Y, hp.pos.Y)
		mx.Z = math.Max(mx.Z, hp.pos.Z)
	}
	ssize := mx.Sub(mn)
	irad := (ssize.X + ssize.Y + ssize.Z) / 3. / ((float64(pm.camera.X + pm.camera.Y)) / 2.0) * 2.0
	for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*HPoint)
		//println(hp.pos.X, hp.pos.Y, hp.pos.Z)
		hp.r2 = irad * irad
		hp.n = 0
		hp.flux = V{}
	}
}

func (pm *PhotonMapping) genPhoton() (*Ray, *V) {
	if len(pm.world.Lights) != 1 {
		panic(fmt.Errorf("only support 1 light, but found %d", len(pm.world.Lights)))
	}

	f := NewV(1, 1, 1).Mul(2500 * 4.0 * math.Pi)
	p := 2. * math.Pi * rand.Float64()
	t := 2. * math.Acos(math.Sqrt(rand.Float64()))
	st := math.Sin(t)

	return &Ray{pm.world.Lights[0], *NewV(math.Cos(p)*st, math.Cos(t), math.Sin(p)*st)}, &f
}

func (pm *PhotonMapping) Pass2(photonNum int) {
	bar := Pbar{}
	bar.Init(photonNum)

	for i := 0; i < photonNum; i++ {
		ray, flux := pm.genPhoton()
		pm.Trace(ray, 0, true, *flux, *NewV(1, 1, 1), -1)
		bar.Tick()
	}
	fmt.Println()
	// density estimation

	for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*HPoint)
		//fmt.Printf("\n%#v\n", hp.flux)
		//println(hp.r2)
		pm.img.Pixel[hp.pix].Add_(hp.flux.Mul(1.0 / (math.Pi * hp.r2 * float64(photonNum))))
	}
}

func (pm *PhotonMapping) Render(photonNum int) {
	rand.Seed(42)

	photonNum *= 1000
	startT := time.Now()
	pm.Pass1()
	pass1T := time.Since(startT)
	pm.Pass2(photonNum)
	pass2T := time.Since(startT) - pass1T

	pm.img.DebugSave("debug_"+pm.path, fmt.Sprint("algo:pm, p_num:", photonNum, ", time:", pass1T, "|", pass2T)) // debug
	pm.img.Save(pm.path)
}
