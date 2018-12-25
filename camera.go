package PPM

type Camera struct {
	// 正交摄像机 暂时只能拍 Ori = (0,1,0)
	Pos, Ori V
	W, H     int // 画幅大小
	X, Y     int // 输出图片大小
}

func (c *Camera) Look(x, y float64) Ray {
	r := Ray{}
	r.Pos = c.Pos
	x = (x - 0.5) * float64(c.W)
	y = (y - 0.5) * float64(c.H)
	r.Dir = c.Ori.Add(V{1, 0, 0}.Mul(x)).Add(V{0, -1, 0}.Mul(y)).Norm()
	return r
}
