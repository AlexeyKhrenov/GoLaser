package models

var fieldPrototype Field
var roundPrototype Round
var overallScore Score
var shotPrototype Shot
var targetPrototype Target

func SetPrototypes(f Field, r Round, os Score, s Shot, t Target) {
	fieldPrototype = f
	roundPrototype = r
	overallScore = os
	shotPrototype = s
	targetPrototype = t
}
