package main

import (
	"fmt"
	"time"
)

func format(t time.Duration) string {

	h := t / time.Hour
	t -= h * time.Hour

	m := t / time.Minute
	t -= m * time.Minute

	s := t / time.Second
	mil := t.Milliseconds() % 1000

	text := fmt.Sprintf("%02d.%03d", s, mil)
	return text
}
