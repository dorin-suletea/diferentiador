package internal

import (
	"log"
	"time"
)

type Tick struct {
	tag     string
	timeSec int
}
type Ticker struct {
	ticks []Tick
}

func NewSingleUseTicker() *Ticker {
	return &Ticker{}
}

func (t *Ticker) Tick(tag string) *Ticker {
	t.ticks = append(t.ticks, Tick{tag: tag, timeSec: int(time.Now().Unix())})
	return t
}

func (t *Ticker) Print() {
	for i, tick := range t.ticks {
		prevTime := 0
		if i != 0 {
			prevTime = tick.timeSec
			log.Printf("-> %d tag=%s, duration=%d", i, tick.tag, tick.timeSec-prevTime)
		}
	}
}
