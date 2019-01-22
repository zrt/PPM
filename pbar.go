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

var tttt = 0

func (p *Pbar) Step(now int, newline bool) {
	p.now = now
	if float64(now-p.lastPrint) >= outputEvery*float64(p.total) {
		if p.lastPrint == 0 {
			newline = true
		}
		p.lastPrint = now
		if newline {
			tttt++
			fmt.Printf("\n%c progress: %3.1f%% (%3d/%3d), used %4.2fs, eta %4.2fs, speed %4.2f/s",
				"|/-\\"[tttt&3],
				float64(now)/float64(p.total)*100, now, p.total,
				time.Now().Sub(p.startTime).Seconds(),
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)),
				float64(now)/time.Now().Sub(p.startTime).Seconds())
		} else {
			tttt++
			fmt.Printf("\r%c progress: %3.1f%% (%3d/%3d), used %4.2fs, eta %4.2fs, speed %4.2f/s",
				"|/-\\"[tttt&3],
				float64(now)/float64(p.total)*100, now, p.total,
				time.Now().Sub(p.startTime).Seconds(),
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)),
				float64(now)/time.Now().Sub(p.startTime).Seconds())
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
			tttt++
			fmt.Printf("\n%c progress: %3.1f%% (%3d/%3d), used %4.2fs, eta %4.2fs, speed %4.2f/s",
				"|/-\\"[tttt&3],
				float64(now)/float64(p.total)*100, now, p.total,
				time.Now().Sub(p.startTime).Seconds(),
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)),
				float64(now)/time.Now().Sub(p.startTime).Seconds())
		} else {
			tttt++
			fmt.Printf("\r%c progress: %3.1f%% (%3d/%3d), used %4.2fs, eta %4.2fs, speed %4.2f/s",
				"|/-\\"[tttt&3],
				float64(now)/float64(p.total)*100, now, p.total,
				time.Now().Sub(p.startTime).Seconds(),
				float64(p.total-now)*(time.Now().Sub(p.startTime).Seconds()/float64(now)),
				float64(now)/time.Now().Sub(p.startTime).Seconds())
		}
	}
}
