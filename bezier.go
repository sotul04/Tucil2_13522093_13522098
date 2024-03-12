package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

type Point struct {
	x float64
	y float64
}

func midPoint(p1, p2 Point) Point {
	mid_x := (p1.x + p2.x) / 2
	mid_y := (p1.y + p2.y) / 2
	return Point{mid_x, mid_y}
}

type BezierPoints struct {
	list []Point
	neff int
}

func (bp *BezierPoints) insertBefore(other BezierPoints) {
	bp.list = append(other.list, bp.list...)
	bp.neff = bp.neff + other.neff
}

func (bp *BezierPoints) insertAfter(other BezierPoints) {
	bp.list = append(bp.list, other.list...)
	bp.neff = bp.neff + other.neff
}

func (bp *BezierPoints) insertLast(other_point ...Point) {
	bp.list = append(bp.list, other_point...)
	bp.neff = bp.neff + len(other_point)
}

func (bp *BezierPoints) insertFirst(other_point ...Point) {
	bp.list = append(other_point, bp.list...)
	bp.neff = bp.neff + len(other_point)
}

func (bp BezierPoints) findCurve(maxIter int) BezierPoints {
	result := findPoints(bp, 0, maxIter)
	result.insertFirst(bp.list[0])
	result.insertLast(bp.list[len(bp.list)-1])

	return result
}

func findPoints(bp BezierPoints, iter, maxIter int) BezierPoints {
	if iter >= maxIter {
		return BezierPoints{}
	}
	left, mid, right := findMidPoint(bp)
	leftPoints := findPoints(left, iter+1, maxIter)
	rightPoints := findPoints(right, iter+1, maxIter)

	mid.insertBefore(leftPoints)
	mid.insertAfter(rightPoints)
	return mid
}

func findMidPoint(bp BezierPoints) (BezierPoints, BezierPoints, BezierPoints) {
	tmpBP := []BezierPoints{}

	left := BezierPoints{}
	left.insertLast(bp.list[0])

	iter := bp.neff - 1

	right := BezierPoints{}
	right.insertFirst(bp.list[iter])

	for i := 0; i < iter; i++ {
		if i == 0 {
			tmpBP = append(tmpBP, BezierPoints{})
			traversal := bp.neff - 1
			for j := 0; j < traversal; j++ {
				tmpBP[0].insertLast(midPoint(bp.list[j], bp.list[j+1]))
			}
		} else {
			tmpBP = append(tmpBP, BezierPoints{})
			traversal := tmpBP[i-1].neff - 1
			for j := 0; j < traversal; j++ {
				tmpBP[i].insertLast(midPoint(tmpBP[i-1].list[j], tmpBP[i-1].list[j+1]))
			}
		}
		left.insertLast(tmpBP[i].list[0])
		right.insertFirst(tmpBP[i].list[iter-i-1])
	}
	return left, tmpBP[iter-1], right

}

func drawBezierCurve(dc *gg.Context, points BezierPoints, corner BezierPoints) {
	dc.MoveTo(points.list[0].x, points.list[0].y)
	for i := 1; i < points.neff; i++ {
		dc.LineTo(points.list[i].x, points.list[i].y)
	}
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(2)
	dc.MoveTo(corner.list[0].x, corner.list[0].y)
	for i := 1; i < corner.neff; i++ {
		dc.LineTo(corner.list[i].x, corner.list[i].y)
	}
	dc.SetRGB(1, 1, 1)

	dc.Stroke()
}

func main() {
	point1 := Point{30, 30}
	point2 := Point{220, 380}
	point3 := Point{400, 20}
	point4 := Point{450, 375}
	// point5 := Point{400, 10}
	// point6 := Point{540, 120}
	// point7 := Point{570, 350}
	// point8 := Point{415, 380}
	points := BezierPoints{}
	points.insertLast(point1, point2, point3, point4)
	// points.insertLast(point1, point2, point3, point4, point5, point6, point7, point8)
	// points.insertLast(point1, point2, point3)
	curve := points.findCurve(7)

	const W = 600
	const H = 400

	dc := gg.NewContext(W, H)
	dc.InvertY()

	drawBezierCurve(dc, curve, points)

	if err := dc.SavePNG("bezier_curve_2.png"); err != nil {
		fmt.Println("Error saving PNG:", err)
	}
}
