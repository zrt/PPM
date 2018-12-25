package PPM

type Sampler interface {
	Shape() (int, int) // output shape
	GenPixelList(chan PixelPos)
}

type SamplerBalanced struct {
	W, H int
}

func (s SamplerBalanced) Shape() (int, int) {
	return s.W, s.H
}

func (s SamplerBalanced) GenPixelList(c chan PixelPos) {
	// for i := 0; i < s.H; i++ {
	// 	for j := 0; j < s.W; j++ {
	// 		if i == 250 && j == 430 {
	// 			p := PixelPos{i, j, float32(i) / float32(s.H-1), float32(j) / float32(s.W-1), float32(1)}
	// 			c <- p
	// 		} else {
	// 			// p := PixelPos{i, j, float32(i) / float32(s.H-1), float32(j) / float32(s.W-1), float32(0.05)}
	// 			// c <- p
	// 		}

	// 	}
	// }
	// close(c)
	// return

	for i := 0; i < s.H; i++ {
		for j := 0; j < s.W; j++ {
			p := PixelPos{i, j, float64(i) / float64(s.H-1), float64(j) / float64(s.W-1), 1.}
			c <- p
		}
	}
}
