package PPM

type Camera struct {
	// 正交摄像机 暂时只能拍 Ori = (0,0,-1)
	Pos, Ori V
	W, H     float64 // 画幅大小
	X, Y     int     // 输出图片大小
}

func (c *Camera) Look(a, b int) *Ray { // k 0 ~ 1
	x := float64(a) / float64(c.X-1)
	y := float64(b) / float64(c.Y-1)
	r := Ray{}
	x = (x - 0.5) * c.W
	y = (y - 0.5) * c.H
	r.Dir = c.Ori.Add(V{1, 0, 0}.Mul(x)).Add(V{0, -1, 0}.Mul(y)).Norm()
	r.Pos = c.Pos.Add(r.Dir.Mul(50)) // 焦距50
	return &r
}
