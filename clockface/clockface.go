package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)

	p = makeHand(p, secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func secondsInRadians(t time.Time) float64 {
	seconds_fraction := float64(t.Second()) / 60
	seconds_radians := seconds_fraction * 2 * math.Pi
	return seconds_radians
}

func minutesInRadians(t time.Time) float64 {
	return float64(t.Minute())/30*math.Pi + (secondsInRadians(t) / 60)
}

func angleToPoint(radian float64) Point {
	return Point{
		X: math.Sin(radian),
		Y: math.Cos(radian),
	}
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func secondHandPoint(t time.Time) Point {
	seconds_radians := secondsInRadians(t)
	return angleToPoint(seconds_radians)
}

func minuteHandPoint(t time.Time) Point {
	minute_radians := minutesInRadians(t)
	return angleToPoint(minute_radians)
}

func minuteHand(w io.Writer, t time.Time) {
	p := minuteHandPoint(t)
	p = makeHand(p, minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
