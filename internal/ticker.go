package internal

import (
	"log"
	"time"
)

type Tick struct {
	tag        string
	timeMillis int64
}
type Ticker struct {
	ticks []Tick
}

func NewSingleUseTicker() *Ticker {
	ret := &Ticker{}
	ret.Tick("-")
	return ret
}

func (t *Ticker) Tick(tag string) *Ticker {
	t.ticks = append(t.ticks, Tick{tag: tag, timeMillis: time.Now().UnixMilli()})
	return t
}

func (t *Ticker) Print() {
	for i, tick := range t.ticks {
		prevTime := int64(0)
		if i != 0 {
			prevTime = t.ticks[i-1].timeMillis
			log.Printf("-> %d tag=%s, duration=%d", i, tick.tag, tick.timeMillis-prevTime)
		}
	}
}
