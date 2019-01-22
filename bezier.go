package PPM

import "math"
import "math/rand"

// import "fmt"

var BEZIERSTRICT bool = false

type Bezier struct { // 贝塞尔旋转体
	points []V // 控制点
	degree int
	axis   V
	start  V
	mat    Material // 材质
	name   string   // 名称
	center V        // bounding
	mn     V        // bounding
	mx     V        // bounding
}

type mat3d [3][3]float64

func newmat3d(a, b, c, d, e, f, g, h, i float64) *mat3d {
	m := mat3d{}
	m[0][0] = a
	m[0][1] = b
	m[0][2] = c
	m[1][0] = d
	m[1][1] = e
	m[1][2] = f
	m[2][0] = g
	m[2][1] = h
	m[2][2] = i
	return &m
}

func (m *mat3d) det() float64 {
	return m[0][0]*m[1][1]*m[2][2] + m[1][0]*m[2][1]*m[0][2] + m[2][0]*m[0][1]*m[1][2] - m[0][2]*m[1][1]*m[2][0] - m[1][2]*m[2][1]*m[0][0] - m[2][2]*m[0][1]*m[1][0]
}

func (m *mat3d) inv() *mat3d {
	ret := mat3d{}
	a := [3][6]float64{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = m[i][j]
		}
		a[i][i+3] = 1
	}

	for i := 0; i < 3; i++ {
		pos := i
		for j := i + 1; j < 3; j++ {
			if math.Abs(a[j][i]) > math.Abs(a[pos][i]) {
				pos = j
			}
		}
		if pos != i {
			for j := 0; j < 6; j++ {
				a[pos][j], a[i][j] = a[i][j], a[pos][j]
			}
		}
		for j := 5; j >= i; j-- {
			a[i][j] /= a[i][i]
		}
		for j := 0; j < 3; j++ {
			if i != j {
				if math.Abs(a[j][i]) > 1e-10 {
					t := a[j][i]
					for k := 0; k < 6; k++ {
						a[j][k] -= t * a[i][k]
					}
				}
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret[i][j] = a[i][j+3]
		}
	}
	return &ret
}

func (m *mat3d) mul(v V) V {
	ret := V{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret.Set(i, ret.Val(i)+m[i][j]*v.Val(j))
		}
	}
	return ret
}

func (b Bezier) newtonIt(arg V, r *Ray, maxit int, th float64) V {
	const INF = 1e20
	funcval := b.targetFunc(arg, r)
	for i := 0; i < maxit; i++ {
		// fmt.Println(arg, funcval)
		if arg.Y-1 > 0.2 || arg.Y < -0.2 {
			// u 不合法
			return V{INF, INF, INF}
		}
		if funcval.X*funcval.X+funcval.Y*funcval.Y+funcval.Z*funcval.Z < th {
			return arg
		}

		pt := b.getPoint(arg.Y)
		dt := b.getDt(arg.Y)

		f := newmat3d(r.Dir.X, -math.Cos(arg.Z)*dt.X, math.Sin(arg.Z)*(pt.X-b.start.X), r.Dir.Y, -dt.Y, 0, r.Dir.Z, -math.Sin(arg.Z)*dt.X, -math.Cos(arg.Z)*(pt.X-b.start.X))

		if math.Abs(f.det()) < 1e-8 {
			return V{INF, INF, INF}
		}
		f = f.inv()
		arg = arg.Sub(f.mul(funcval))
		funcval = b.targetFunc(arg, r)
	}
	return V{INF, INF, INF}
}

// 是否相交，距离，法向量
func (b Bezier) Collide(r *Ray) (bool, float64, *V) {
	// println("calc")
	//arg := b.merge(0, 1, r)
	mn := V{1e20, 1e20, 1e20}
	TT := 10
	if BEZIERSTRICT {
		TT = 25
	}
	for tt := 0; tt < TT; tt++ {
		// init u random (0,1)
		arginit := V{}
		arginit.X = V{b.center.X, r.Pos.Y, b.center.Z}.Disto(r.Pos) // t
		arginit.Y = rand.Float64()                                  // u
		arginit.Z = rand.Float64() * 2 * math.Pi                    // theta
		res := b.newtonIt(arginit, r, TT+10, 1e-8)
		if res.X < 1e-3 || res.X > 1e19 || res.Y < 0 || res.Y > 1 {
			continue
		}
		if res.X < mn.X {
			mn = res
		}
	}
	// trick?
	if mn.X < 1e19 {
		// 找对面
		TT = 5
		if BEZIERSTRICT {
			TT = 10
		}
		for tt := 0; tt < 2; tt++ {
			// init u random (0,1)
			arginit := V{}
			pt := r.Pos.Add(r.Dir.Mul(mn.X))
			arginit.X = mn.X - pt.Disto(V{b.start.X, pt.Y, b.start.Z})*2
			arginit.Y = mn.Y           // u
			arginit.Z = mn.Z + math.Pi // theta
			res := b.newtonIt(arginit, r, TT+10, 1e-8)
			if res.X < 0 || res.X > 1e19 || res.Y < 0 || res.Y > 1 {
				continue
			}
			if res.X < mn.X {
				mn = res
			}
		}
	}

	if mn.X > 1e19 {
		return false, -1, nil
	}
	// pt := r.Pos.Add(r.Dir.Mul(mn.X))
	dis := mn.X
	dt := b.getDt(mn.Y)
	dt.Z = dt.X * math.Sin(mn.Z)
	dt.X = dt.X * math.Cos(mn.Z)
	norm := dt.Cross(b.axis).Cross(dt).Norm()
	if norm.Dot(r.Dir) > 0 {
		norm = norm.Mul(-1)
	}
	// println(dis)
	return true, dis, &norm
}

// arg: (t,u,theta)
func (b Bezier) targetFunc(arg V, r *Ray) V {
	p1 := b.getPoint3D(arg.Y, arg.Z)
	p2 := r.Pos.Add(r.Dir.Mul(arg.X))
	return p2.Sub(p1)
}

func (b Bezier) GetMaterial() Material {
	return b.mat
}

func (b Bezier) GetName() string {
	return b.name
}

func (b Bezier) GetCenter() V {
	return b.center
}

func (b Bezier) GetMin() V {
	return b.mn
}

func (b Bezier) GetMax() V {
	return b.mx
}

func NewBezier(p []V, d int, axis V, start V, mat Material, name string) Bezier {
	if len(p) != d+1 {
		panic("invalid bezier")
	}
	center := V{}
	center.X = start.X
	center.Z = start.Z
	ysum := 0.
	ymin := 1e20
	ymax := -1e20
	R := 0.
	for i := 0; i <= d; i++ {
		R = math.Max(R, V{start.X, 0, start.Z}.Disto(V{p[i].X, 0, p[i].Z}))
		ysum += p[i].Y
		ymin = math.Min(ymin, p[i].Y)
		ymax = math.Max(ymax, p[i].Y)
	}
	ysum /= float64(d + 1)
	center.Y = ysum
	mn := V{start.X - R, ymin, start.Z - R}
	mx := V{start.X + R, ymax, start.Z + R}
	return Bezier{p, d, axis, start, mat, name, center, mn, mx}
}

func (b Bezier) getPoint(t float64) V {
	tmp := make([]V, len(b.points))
	copy(tmp, b.points)
	for k := 0; k < b.degree; k++ {
		for i := 0; i < b.degree-k; i++ {
			tmp[i] = tmp[i].Mul(1 - t).Add(tmp[i+1].Mul(t))
		}
	}
	return tmp[0]
}

func (b Bezier) getDt(t float64) V {
	tmp := make([]V, len(b.points))
	copy(tmp, b.points)
	for k := 0; k < b.degree-1; k++ {
		for i := 0; i < b.degree-k; i++ {
			tmp[i] = tmp[i].Mul(1 - t).Add(tmp[i+1].Mul(t))
		}
	}
	return tmp[1].Sub(tmp[0]).Mul(float64(b.degree))
}

func (b Bezier) getPoint3D(t float64, theta float64) V {
	p := b.getPoint(t)
	q := p.Sub(V{b.start.X, 0, b.start.Z})
	p.Z = q.X*math.Sin(theta) + b.start.Z
	p.X = q.X*math.Cos(theta) + b.start.X
	return p
}

func (b Bezier) GenBalls(eps, r float64) []Object {
	ret := []Object{}
	for theta := 0.; theta < 2*math.Pi; theta += eps {
		for t := 0.; t < 1; t += eps {
			ret = append(ret, Ball{b.getPoint3D(t, theta), r, DiffuseMaterial(V{1, 0, 0}), b.GetName()})
		}
	}
	return ret
}
