package models

import (
	"math/rand"
	"time"
)

type Round struct {
	ShouldBeOverOnMiss bool
	DelayMs            int
	DelayDev           int

	IsOver    bool
	Start     time.Time
	Target    Target
	IsSuccess bool
	Elapsed   time.Duration
}

func newRound() Round {
	r := roundPrototype
	target := newTarget()
	r.Target = target
	return r
}

func (r *Round) start() {
	delay := r.DelayMs
	if r.DelayDev != 0 {
		delay += rand.Intn(r.DelayDev*2) - r.DelayDev
	}

	frame := time.Millisecond * time.Duration(delay)
	time.Sleep(frame)

	r.Start = time.Now().UTC()
}

func (r *Round) shot(s Shot) {
	if s.IsSuccess {

		r.Target.shot(s)
		if r.Target.IsKilled {
			r.IsOver = true
			r.Elapsed = time.Since(r.Start)
			r.IsSuccess = true
		}
	} else {
		if r.ShouldBeOverOnMiss {
			r.IsOver = true
			r.Elapsed = time.Since(r.Start)
			r.IsSuccess = false
		}
	}
}
