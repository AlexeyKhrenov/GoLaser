package models

type Shot struct {
	Health int

	IsSuccess bool
}

func newShot(isSuccess bool) Shot {
	s := shotPrototype
	s.IsSuccess = isSuccess
	return s
}
