package models

import (
	"math/rand"
)

type Target struct {
	Width  float32
	Height float32
	XDev   float32
	YDev   float32
	Health int

	X        float32
	Y        float32
	IsKilled bool
}

func newTarget() Target {
	target := targetPrototype
	x := (fieldPrototype.Width - target.Width) / 2
	y := (fieldPrototype.Height - target.Height) / 2

	x += rand.Float32()*target.XDev*2 + target.XDev
	y += rand.Float32()*target.YDev*2 + target.YDev

	target.X = x
	target.Y = y
	target.IsKilled = false
	return target
}

func (t *Target) shot(s Shot) {
	if s.IsSuccess {
		t.Health -= s.Health
	}

	if t.Health <= 0 {
		t.IsKilled = true
	}
}
