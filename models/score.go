package models

import "fmt"

type Score struct {
	A int
	B int
}

func (s *Score) add(success bool) {
	if success {
		s.A++
	}
	s.B++
}

func (s *Score) ToString() string {
	return fmt.Sprintf("%d / %d", s.A, s.B)
}
