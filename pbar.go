package PPM

// 进度条
import (
	"fmt"
	"time"
)

const outputEvery = 0.005

type Pbar struct {
	startTime time.Time
	total     int
	lastPrint int
}

func (p *Pbar) Init(num int) error {
	p.total = num
	p.startTime = time.Now()
	p.lastPrint = 0
	return nil
}

func (p *Pbar) Step(now int) {
	if float64(now-p.lastPrint) > outputEvery*float64(p.total) {
		p.lastPrint = now
		fmt.Printf("progress: %.1f%% (%d/%d), eta %.2fs\n",
			float64(now)/float64(p.total)*100, now, p.total,
			float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)))
	}
}
