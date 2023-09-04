package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	seconds_radians := secondsInRadians(t)
	secondhand_length := float64(90)

	return Point{
		X: 150 + math.Sin(seconds_radians)*secondhand_length,
		Y: 150 - math.Cos(seconds_radians)*secondhand_length,
	}
}

func secondsInRadians(t time.Time) float64 {
	seconds_fraction := float64(t.Second()) / 60
	seconds_radians := seconds_fraction * 2 * math.Pi
	return seconds_radians
}
