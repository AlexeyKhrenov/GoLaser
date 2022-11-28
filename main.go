package main

import (
	"sim/models"
	"sim/settings"
)

func main() {
	s := settings.Settings{}
	settings.ReadSettings(&s)

	models.SetPrototypes(s.Field, s.Round, s.OverallScore, s.Shot, s.Target)

	feed := make(chan bool)
	go runDetector(feed)
	showField(s, feed)
}
