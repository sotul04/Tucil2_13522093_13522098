package main

import (
	"fmt"
	"math"
	"strconv"

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

// func drawBezierCurve(dc *gg.Context, points BezierPoints, corner BezierPoints) {

// 	dc.SetRGB(1, 0, 0)
// 	dc.MoveTo(points.list[0].x, points.list[0].y)
// 	for i := 1; i < points.neff; i++ {
// 		dc.LineTo(points.list[i].x, points.list[i].y)
// 	}
// 	dc.SetLineWidth(2.5)
// 	dc.Stroke()

// 	dc.SetRGB(0.2, 0.4, 1)
// 	dc.MoveTo(corner.list[0].x, corner.list[0].y)
// 	for i := 1; i < corner.neff; i++ {
// 		dc.LineTo(corner.list[i].x, corner.list[i].y)
// 	}
// 	dc.SetLineWidth(1)
// 	dc.Stroke()

// 	dc.SetRGB(0, 1, 0)
// 	for i := 0; i < points.neff; i++ {
// 		dc.DrawPoint(points.list[i].x, points.list[i].y, 1)
// 		dc.Stroke()
// 	}

// 	dc.SetRGB(1, 1, 0.5)
// 	for i := 0; i < corner.neff; i++ {
// 		dc.DrawPoint(corner.list[i].x, corner.list[i].y, 1)
// 		dc.Stroke()
// 	}
// }

func getPreferredDimension(corner BezierPoints) (int, int, float64, float64, float64, float64) {
	min_x := corner.list[0].x
	min_y := corner.list[0].y

	max_x := corner.list[0].x
	max_y := corner.list[0].y

	pref_x := 600
	pref_y := 600

	r_x := float64(1)
	r_y := float64(1)

	add_x := float64(0)
	add_y := float64(0)

	for i := 1; i < corner.neff; i++ {
		if min_x > corner.list[i].x {
			min_x = corner.list[i].x
		} else if max_x < corner.list[i].x {
			max_x = corner.list[i].x
		}
		if min_y > corner.list[i].y {
			min_y = corner.list[i].y
		} else if max_y < corner.list[i].y {
			max_y = corner.list[i].y
		}
	}

	dx := max_x - min_x
	dy := max_y - min_y

	if dx < 400 {
		r_x = 400.0 / dx
		pref_x = 530
	} else if dx > 535 {
		pref_x = int(dx) + 130
	}

	if dy < 400 {
		r_y = 400.0 / dy
		pref_y = 530
	} else if dy > 535 {
		pref_y = int(dy) + 130
	}

	add_x = 65 - min_x*r_x
	add_y = 65 - min_y*r_y

	return pref_x, pref_y, r_x, r_y, add_x, add_y
}

func drawSketch(points, corner BezierPoints, iter int) {
	pref_x, pref_y, r_x, r_y, add_x, add_y := getPreferredDimension(corner)

	for i := 1; i <= iter; i++ {
		step := int(math.Pow(2, float64(iter-i)))

		newSketch := gg.NewContext(pref_x, pref_y)
		newSketch.SetRGB(1, 0, 0)
		newSketch.InvertY()

		newSketch.MoveTo(r_x*points.list[0].x+add_x, r_y*points.list[0].y+add_y)
		for i := 0; i < points.neff; i += step {
			newSketch.LineTo(r_x*points.list[i].x+add_x, r_y*points.list[i].y+add_y)
		}
		newSketch.SetLineWidth(2.5)
		newSketch.Stroke()

		newSketch.SetRGB(0.2, 0.4, 1)
		newSketch.MoveTo(r_x*corner.list[0].x+add_x, r_y*corner.list[0].y+add_y)
		for i := 1; i < corner.neff; i++ {
			newSketch.LineTo(r_x*corner.list[i].x+add_x, r_y*corner.list[i].y+add_y)
		}
		newSketch.SetLineWidth(1)
		newSketch.Stroke()

		newSketch.SetRGB(0, 1, 0)
		for i := 0; i < points.neff; i += step {
			newSketch.DrawPoint(r_x*points.list[i].x+add_x, r_y*points.list[i].y+add_y, 1.2)
			newSketch.Stroke()
		}

		newSketch.SetRGB(1, 1, 0.5)
		for i := 0; i < corner.neff; i++ {
			newSketch.DrawPoint(r_x*corner.list[i].x+add_x, r_y*corner.list[i].y+add_y, 1.2)
			newSketch.Stroke()
		}

		if err := newSketch.SavePNG("bezier_curve_" + strconv.Itoa(i) + ".png"); err != nil {
			fmt.Println("Error saving PNG:", err)
		}
	}
}

func main() {
	point1 := Point{-50, 350}
	point2 := Point{30, 30}
	point3 := Point{260, 340}
	point4 := Point{360, 189}
	point5 := Point{230, 110}
	point6 := Point{430, 50}
	point7 := Point{550, 190}
	point8 := Point{445, 300}
	points := BezierPoints{}
	// points.insertLast(point1, point2, point3, point4)
	points.insertLast(point1, point2, point3, point4, point5, point6, point7, point8)
	curve := points.findCurve(7)

	drawSketch(curve, points, 7)
}
