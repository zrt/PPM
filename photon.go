package PPM

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const ALPHA = 0.7

type PhotonMapping struct {
	world        *World
	camera       *Camera
	path         string
	img          Image
	hitpoints    *list.List
	table        *HashTable
	lock         sync.Mutex
	lumi         float64
	antiAliasing int
}

type HPoint struct { // hit point
	f, pos, nrm, flux V
	r2                float64
	n                 uint // n = N / ALPHA in the paper
	pix               int
}

func NewPhotonMapping(world *World, camera *Camera, path string, antiAliasing int, lumi float64) *PhotonMapping {
	pm := &PhotonMapping{world, camera, path, Image{}, list.New(), nil, *new(sync.Mutex), lumi, antiAliasing}
	camera.X *= antiAliasing
	camera.Y *= antiAliasing
	pm.img.Init(camera.X, camera.Y)
	println("#lights: ", len(world.Lights), "#objects: ", len(world.Objects), "anti-aliasing: ", antiAliasing)
	println(path)
	println(camera.X, camera.Y)
	world.Tree = BuildOBJKDTree(world.Objects, 0)
	return pm
}

func (pm *PhotonMapping) Pass1() (cnt int, irad float64) {

	BEZIERSTRICT = true
	bar := Pbar{}
	bar.Init(pm.camera.X * pm.camera.Y)
	fmt.Println("calc hitpoints...")
	for i := 0; i < pm.camera.X; i++ {
		for j := 0; j < pm.camera.Y; j++ {
			//println(i, j)
			pm.Trace(pm.camera.Look(i, j), 0, false, *NewV(0, 0, 0), *NewV(1, 1, 1), i*pm.camera.Y+j)
			bar.Tick()
		}
	}
	fmt.Println()

	// build_hash_grid
	BEZIERSTRICT = false

	mx := V{-1e20, -1e20, -1e20}
	mn := V{1e20, 1e20, 1e20}
	cnt = 0
	for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*HPoint)
		mn.X = math.Min(mn.X, hp.pos.X)
		mn.Y = math.Min(mn.Y, hp.pos.Y)
		mn.Z = math.Min(mn.Z, hp.pos.Z)
		mx.X = math.Max(mx.X, hp.pos.X)
		mx.Y = math.Max(mx.Y, hp.pos.Y)
		mx.Z = math.Max(mx.Z, hp.pos.Z)
		cnt++
	}

	ssize := mx.Sub(mn)
	fmt.Print("SIZE: ")
	ssize.Print()
	irad = float64(len(pm.world.Lights)) * (ssize.X + ssize.Y + ssize.Z) / 3. / ((float64(pm.camera.X + pm.camera.Y)) / 2.0) * 2.0

	pm.table = NewHashTable(pm.hitpoints, cnt, irad)
	fmt.Println("hitpoints num:", cnt)

	bar.Init(cnt)

	for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*HPoint)
		//println(hp.pos.X, hp.pos.Y, hp.pos.Z)
		hp.r2 = irad * irad
		hp.n = 0
		hp.flux = V{}
		bar.Tick()
	}
	fmt.Println("\nfinish pass1..")
	return
}

func (pm *PhotonMapping) genPhoton() (*Ray, *V) {
	//if len(pm.world.Lights) != 1 {
	//	panic(fmt.Errorf("only support 1 light, but found %d", len(pm.world.Lights)))
	//}
	lightID := rand.Intn(len(pm.world.Lights)) // 随机选一个灯
	f := NewV(1, 1, 1).Mul(pm.lumi * 4.0 * math.Pi)
	p := 2. * math.Pi * rand.Float64()
	t := 2. * math.Acos(math.Sqrt(rand.Float64()))
	st := math.Sin(t)

	return &Ray{pm.world.Lights[lightID], *NewV(math.Cos(p)*st, math.Cos(t), math.Sin(p)*st)}, &f
}

func (pm *PhotonMapping) Pass2(photonNum int) {
	fmt.Println("pass2..")
	bar := Pbar{}
	bar.Init(photonNum)

	const gnum = 1000
	c := make(chan int, gnum)

	for i := 0; i < photonNum; i++ {
		//println(i)
		for j := 0; j < gnum; j++ {
			go func() {
				ray, flux := pm.genPhoton()
				pm.Trace(ray, 0, true, *flux, *NewV(1, 1, 1), -1)
				c <- i
			}()
		}

		for j := 0; j < gnum; j++ {
			<-c
		}
		bar.Tick()

		if i > 0 && i%2000 == 0 {
			for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
				hp := e.Value.(*HPoint)
				pm.img.Pixel[hp.pix].Add_(hp.flux.Mul(1.0 / (math.Pi * hp.r2 * float64(i) * 1000.)))
			}
			pm.img.DebugSave("latest_debug_"+pm.path, fmt.Sprint("algorithm:ppm, p_num:", i)) // debug
			pm.img.Save("latest_full_" + pm.path)
			pm.img.SaveSmall("latest_"+pm.path, pm.antiAliasing)
			for j := 0; j < len(pm.img.Pixel); j++ {
				pm.img.Pixel[j] = V{}
			}
		}
	}

	fmt.Println()
	// density estimation

	for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*HPoint)
		//fmt.Printf("\n%#v\n", hp.flux)
		//println(hp.r2)
		pm.img.Pixel[hp.pix].Add_(hp.flux.Mul(1.0 / (math.Pi * hp.r2 * float64(photonNum) * 1000.)))
	}
}

func (pm *PhotonMapping) Render(photonNum int) {
	rand.Seed(42)

	//photonNum *= 1000
	startT := time.Now()
	cnt, r := pm.Pass1()
	pass1T := time.Since(startT)
	pm.Pass2(photonNum)
	pass2T := time.Since(startT) - pass1T

	pm.img.DebugSave("debug_"+pm.path, fmt.Sprint("algorithm:ppm, p_num:", photonNum, "\n", "time:", pass1T, "|", pass2T, "\n",
		"hitpoints:", cnt, " R:", r)) // debug
	pm.img.Save("full_" + pm.path)
	pm.img.SaveSmall(pm.path, pm.antiAliasing)

}
