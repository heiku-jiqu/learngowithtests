package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	return Point{
		X: clockCentreX + p.X*secondHandLength,
		Y: clockCentreY - p.Y*secondHandLength,
	}
}

func secondsInRadians(t time.Time) float64 {
	seconds_fraction := float64(t.Second()) / 60
	seconds_radians := seconds_fraction * 2 * math.Pi
	return seconds_radians
}

func secondHandPoint(t time.Time) Point {
	seconds_radians := secondsInRadians(t)

	return Point{
		X: math.Sin(seconds_radians),
		Y: math.Cos(seconds_radians),
	}
}
