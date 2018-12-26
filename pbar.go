package PPM

import (
	"fmt"
	"time"
)

// 进度条

const outputEvery = 0.005

type Pbar struct {
	startTime time.Time
	total     int
	lastPrint int
	now       int
}

func (p *Pbar) Init(num int) {
	p.total = num
	p.startTime = time.Now()
	p.lastPrint = 0
	p.now = 0
}

func (p *Pbar) Step(now int, newline bool) {
	p.now = now
	if float64(now-p.lastPrint) >= outputEvery*float64(p.total) {
		if p.lastPrint == 0 {
			newline = true
		}
		p.lastPrint = now
		if newline {
			fmt.Printf("\nprogress: %.1f%% (%d/%d), eta %.2fs",
				float64(now)/float64(p.total)*100, now, p.total,
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)))
		} else {
			fmt.Printf("\rprogress: %.1f%% (%d/%d), eta %.2fs",
				float64(now)/float64(p.total)*100, now, p.total,
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)))
		}
	}
}

func (p *Pbar) Tick() {
	p.now++
	now := p.now
	newline := false
	if float64(now-p.lastPrint) >= outputEvery*float64(p.total) {
		if p.lastPrint == 0 {
			newline = true
		}
		p.lastPrint = now
		if newline {
			fmt.Printf("\nprogress: %.1f%% (%d/%d), eta %.2fs",
				float64(now)/float64(p.total)*100, now, p.total,
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)))
		} else {
			fmt.Printf("\rprogress: %.1f%% (%d/%d), eta %.2fs",
				float64(now)/float64(p.total)*100, now, p.total,
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)))
		}
	}
}
