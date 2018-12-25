package PPM

import "math"

type Ball struct {
	Pos  V        // 位置
	R    float64  // 半径
	Mat  Material // 材质
	Name string   // 名称
}

// 是否相交，交点
func (b Ball) Collide(r *Ray) (bool, V) {
	//https://www.jianshu.com/p/1b008ed86627
	Voc := b.Pos.Sub(r.Pos)
	k := Voc.Dot(r.Dir)
	Poc := r.Dir.Mul(k)
	d2 := Voc.Dot(Voc) - Poc.Dot(Poc)
	r2 := b.R * b.R
	if d2 < r2 && k > 0 {
		// Vd := Poc.Sub(Voc)
		Pd := r.Dir.Mul(-math.Sqrt(r2 - d2))
		// Vs := Vd.Add(Pd)
		return true, r.Pos.Add(Poc).Add(Pd)
	} else {
		return false, V{}
	}
}

// 取得局部颜色(PT)
func (b Ball) GetLocalColor() V {
	// todo
	return b.Mat.Color
}

// 取得材料透射、折射率
func (b Ball) GetMaterialWRWT() (float64, float64) {
	return b.Mat.WR, b.Mat.WT
}

// 获得反射光线(在已经相交时)
func (b Ball) GetRefRay(r *Ray) *Ray {
	//https://www.jianshu.com/p/1b008ed86627
	Voc := b.Pos.Sub(r.Pos)
	k := Voc.Dot(r.Dir)
	Poc := r.Dir.Mul(k)
	d2 := Voc.Dot(Voc) - Poc.Dot(Poc)
	r2 := b.R * b.R
	if d2 < r2 && k > 0 {
		Vd := Poc.Sub(Voc)
		Pd := r.Dir.Mul(-math.Sqrt(r2 - d2))
		Vs := Vd.Add(Pd)
		ret := Ray{}
		ret.Pos = r.Pos.Add(Poc).Add(Pd)
		Vss := Vs.Norm().Mul(b.R - Vd.Dot(Vs.Norm()))
		ret.Dir = Pd.Neg().Add(Vss.Mul(2)).Norm()
		return &ret
	} else {
		panic("err GetRefRay")
		return &Ray{}
	}
}

func (b Ball) GetTransRay(*Ray) *Ray {
	return nil
}

func (b Ball) VFrom(pos V) V {
	return b.Pos.Sub(pos)
}

func (b Ball) GetName() string {
	return b.Name
}
