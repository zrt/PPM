package PPM

import (
	"fmt"
	"math"
	"math/rand"
)

const TRACEDEPLIMIT = 20 // 最大trace深度

// Eye Trace/ Photon Trace
func (pm *PhotonMapping) Trace(r *Ray, dep int, photon bool, flux, adj V, pix int) {
	dep++
	isCollide, dis, id := pm.world.Collide(r)
	if !isCollide || dep >= TRACEDEPLIMIT {
		return
	}
	x := r.Pos.Add(r.Dir.Mul(dis)) // 交点
	obj := pm.world.Objects[id]    // 相交的obj
	objMt := obj.GetMaterial()     // obj的材质
	n := obj.GetNormal(x)          // 得到交点法向量
	nl := n
	if n.Dot(r.Dir) >= 0 {
		nl = n.Mul(-1)
	}
	debugN := 1
	if pix == debugN {
		fmt.Println()
		r.Pos.Print()
		r.Dir.Print()
		println(dep)
		println(photon)
		flux.Print()
		adj.Print()
		println(obj.GetName())
		x.Print()
		fmt.Println()
	}
	//if obj.GetName() != "w_up" {
	//
	//	println(obj.GetName())
	//}

	if objMt.T == MT_SPEC {
		// mirror
		pm.Trace(&Ray{x, r.Dir.Sub(n.Mul(2).Mul(n.Dot(r.Dir)))}, dep, photon, flux.Mulv(objMt.Color), adj.Mulv(objMt.Color), pix)
	} else if objMt.T == MT_REFR {
		// glass
		lr := &Ray{x, r.Dir.Sub(n.Mul(2).Mul(n.Dot(r.Dir)))} // 反射光线
		into := n.Dot(nl) > 0.0
		nc := 1.0
		nt := 1.5
		nnt := 0.
		if !into {
			nnt = nt / nc
		} else {
			nnt = nc / nt
		}
		ddn := r.Dir.Dot(nl)
		cos2t := 1 - nnt*nnt*(1-ddn*ddn)
		if cos2t < 0 {
			// total internal reflection
			pm.Trace(lr, dep, photon, flux, adj, pix)
			return
		}
		// else
		td := V{}
		if into {
			td = (r.Dir.Mul(nnt).Sub(n.Mul(ddn*nnt + math.Sqrt(cos2t)))).Norm()
		} else {
			td = (r.Dir.Mul(nnt).Add(n.Mul(ddn*nnt + math.Sqrt(cos2t)))).Norm()
		}
		a := nt - nc
		b := nt + nc
		R0 := a * a / (b * b)
		c := 0.
		if into {
			c = 1 + ddn
		} else {
			c = 1 - td.Dot(n)
		}
		Re := R0 + (1-R0)*c*c*c*c*c
		P := Re
		rr := &Ray{x, td}
		fa := objMt.Color.Mulv(adj)
		if !photon {
			// eye
			pm.Trace(lr, dep, photon, flux, fa.Mul(Re), pix)
			pm.Trace(rr, dep, photon, flux, fa.Mul(1.-Re), pix)
		} else {
			// photon
			if rand.Float64() < P {
				pm.Trace(lr, dep, photon, flux, fa, pix)
			} else {
				pm.Trace(rr, dep, photon, flux, fa, pix)
			}
		}
	} else if objMt.T == MT_DIFF {
		// diffuse
		//println("lala")
		r1 := 2. * math.Pi * rand.Float64()
		r2 := rand.Float64()
		r2s := math.Sqrt(r2)
		w := nl
		u := V{}
		if math.Abs(w.X) > .1 {
			u = NewV(0, 1, 0).Cross(&w).Norm()
		} else {
			u = NewV(1, 0, 0).Cross(&w).Norm()
		}
		v := w.Cross(&u)
		d := (u.Mul(math.Cos(r1) * r2s).Add(v.Mul(math.Sin(r1) * r2s).Add(w.Mul(math.Sqrt(1 - r2))))).Norm()
		if !photon {
			// eye
			hp := &HPoint{objMt.Color.Mulv(adj), x, n, V{}, 0, 0, pix}
			if pix == debugN {
				fmt.Printf("\n%#v\n", hp)
			}
			pm.hitpoints.PushFront(hp)
		} else {
			// photon
			for e := pm.hitpoints.Front(); e != nil; e = e.Next() {
				hp := e.Value.(*HPoint)
				v = hp.pos.Sub(x)
				// check normals to be closer than 90 degree (avoids some edge brightning)
				if hp.nrm.Dot(n) > 1e-3 && v.Dot(v) < hp.r2 {
					g := (float64(hp.n)*(ALPHA) + (ALPHA)) / (float64(hp.n)*ALPHA + 1.0)
					hp.r2 *= g
					hp.n++
					hp.flux = (hp.flux.Add(hp.f.Mulv(flux).Mul(1. / math.Pi))).Mul(g)
					//println(hp.n, g)
					//hp.flux.Print()
					//fmt.Printf("\n%#v\n", hp)
				}
			}
			p := 0.
			f := objMt.Color
			if f.X > f.Y && f.X > f.Z {
				p = f.X
			} else if f.Y > f.Z {
				p = f.Y
			} else {
				p = f.Z
			}
			if rand.Float64() < p {
				pm.Trace(&Ray{x, d}, dep, photon, flux.Mulv(f).Mul(1./p), adj, pix)
			}
		}
	} else {
		panic(fmt.Errorf("unknown material type: %d, id: %d, name: %s", objMt.T, id, obj.GetName()))
	}
}
