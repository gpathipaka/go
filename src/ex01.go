package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"
	"time"
)

//Point Type
type Point struct {
	X, Y float64
}

//ColoredPoint is
type ColoredPoint struct {
	Point
	Color color.RGBA
}

//Distance function
func Distance(p, fq Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance Method associated with Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Path type Slice1
type Path []Point

//Distance ss
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//ScaleBy dsd
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	const day = 24 * time.Hour
	//fmt.Println(day)
	fmt.Println(day.Seconds())

	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	r := &Point{1, 2}
	r.ScaleBy(5)
	fmt.Println("Point r ", *r)

	var r1 *Point
	r1 = &Point{1, 2}
	r1.ScaleBy(2)
	fmt.Println("Point r1 ", *r1)

	r2 := Point{1, 2}
	//ptr := &r2
	//ptr.ScaleBy(2)
	(&r2).ScaleBy(3)
	fmt.Println("Point r1 ", r2)

	fmt.Println(strings.Repeat("*", 25))
	p1 := Point{1, 2}
	q1 := Point{4, 2}

	dis := p1.Distance
	fmt.Println(dis(q1))

	var origin Point
	fmt.Println("Distrance From Origin : ", dis(origin))

	scaleP1 := p1.ScaleBy
	scaleP1(2)
	fmt.Println("scaleP1(2)", p1)
	scaleP1(3)
	fmt.Println("scaleP1(2)", p1)
	scaleP1(5)
	fmt.Println("scaleP1(2)", p1)
	fmt.Printf("Type %T\n", p1)

	dis1 := Point.Distance

	fmt.Println(dis1(p, q))
	fmt.Printf("Type %T\n", dis1)
	scale := (*Point).ScaleBy
	scale(&p, 5)
	fmt.Println(p)
	fmt.Printf("Type %T\n", scale)
}
