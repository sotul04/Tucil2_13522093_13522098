package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

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

// func getRatioPoint(p0, p1, p2 Point, ratio float64) Point {
// 	one_min_r := 1 - ratio
// 	end_x := one_min_r*one_min_r*p0.x + 2*one_min_r*ratio*p1.x + ratio*ratio*p2.x
// 	end_y := one_min_r*one_min_r*p0.y + 2*one_min_r*ratio*p1.y + ratio*ratio*p2.y
// 	return Point{end_x, end_y}
// }

func getRatioPoint4(points BezierPoints, sketch *gg.Context, ratio float64, pref_y int, r_x, r_y, add_x, add_y float64) Point {
	if points.neff == 2 {
		tmpPoint := Point{(1-ratio)*points.list[0].x + ratio*points.list[1].x, (1-ratio)*points.list[0].y + ratio*points.list[1].y}
		sketch.SetRGB(0, 1, 0)
		sketch.MoveTo(r_x*points.list[0].x+add_x, -(r_y*points.list[0].y+add_y)+float64(pref_y))
		sketch.LineTo(r_x*points.list[1].x+add_x, -1*(r_y*points.list[1].y+add_y)+float64(pref_y))
		sketch.SetLineWidth(0.3)
		sketch.Stroke()
		sketch.SetRGB(1, 1, 0)
		sketch.DrawPoint(r_x*tmpPoint.x+add_x, -1*(r_y*tmpPoint.y+add_y)+float64(pref_y), 1)
		sketch.Stroke()
		return tmpPoint
	}
	newpoints := BezierPoints{}
	for i := 1; i < points.neff; i++ {
		sketch.SetRGB(0, 1, 0)
		sketch.MoveTo(r_x*points.list[i-1].x+add_x, -(r_y*points.list[i-1].y+add_y)+float64(pref_y))
		sketch.LineTo(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y))
		sketch.SetLineWidth(0.3)
		sketch.Stroke()
		sketch.SetRGB(1, 1, 0)
		tmpPoint := Point{(1-ratio)*points.list[i-1].x + ratio*points.list[i].x, (1-ratio)*points.list[i-1].y + ratio*points.list[i].y}
		sketch.DrawPoint(r_x*tmpPoint.x+add_x, -1*(r_y*tmpPoint.y+add_y)+float64(pref_y), 1)
		sketch.Stroke()
		newpoints.insertLast(tmpPoint)
	}
	return getRatioPoint4(newpoints, sketch, ratio, pref_y, r_x, r_y, add_x, add_y)
}

func (bp BezierPoints) drawCurveBruteForce() BezierPoints {
	pref_x, pref_y, r_x, r_y, add_x, add_y := getPreferredDimension(bp)

	sketch := gg.NewContext(pref_x, pref_y)

	points := BezierPoints{}

	for i := 0; i < 41; i++ {
		points.insertLast(getRatioPoint4(bp, sketch, float64(i)*0.025, pref_y, r_x, r_y, add_x, add_y))
	}

	sketch.SetRGB(1, 0, 0)
	sketch.MoveTo(r_x*points.list[0].x+add_x, -1*(r_y*points.list[0].y+add_y)+float64(pref_y))
	for i := 0; i < 40; i++ {
		sketch.LineTo(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y))
	}
	sketch.SetLineWidth(2.5)
	sketch.Stroke()

	sketch.SetRGB(0.2, 0.5, 1)
	for i := 0; i < 41; i++ {
		sketch.DrawPoint(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y), 1.2)
		sketch.Stroke()
	}

	sketch.SetRGB(0.2, 0.4, 1)
	sketch.MoveTo(r_x*bp.list[0].x+add_x, -1*(r_y*bp.list[0].y+add_y)+float64(pref_y))
	for i := 1; i < bp.neff; i++ {
		sketch.LineTo(r_x*bp.list[i].x+add_x, -1*(r_y*bp.list[i].y+add_y)+float64(pref_y))
	}
	sketch.SetLineWidth(1.8)
	sketch.Stroke()

	sketch.SetRGB(1, 1, 0.5)
	for i := 0; i < bp.neff; i++ {
		sketch.DrawPoint(r_x*bp.list[i].x+add_x, -1*(r_y*bp.list[i].y+add_y)+float64(pref_y), 1.2)
		sketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, bp.list[i].x, bp.list[i].y), r_x*bp.list[i].x+add_x, -1*(r_y*bp.list[i].y+add_y)+float64(pref_y), 0.5, -0.5)
		sketch.Stroke()
	}

	if err := sketch.SavePNG("bezier_curve_brute_force.png"); err != nil {
		fmt.Println("Error saving PNG:", err)
	}

	return points
}

// func (points BezierPoints) findCurveBruteForce4() BezierPoints {
// 	add := 0.01
// 	curve := BezierPoints{}
// 	curve.insertLast(points.list[0])
// 	for i := 1; i < 100; i++ {
// 		curve.insertLast(getRatioPoint4(points.list[0], points.list[1], points.list[2], points.list[3], float64(i)*(add)))
// 	}
// 	curve.insertLast(points.list[2])
// 	return curve
// }

// func (points BezierPoints) findCurveBruteForce() BezierPoints {
// 	add := 0.01
// 	curve := BezierPoints{}
// 	curve.insertLast(points.list[0])
// 	for i := 1; i < 100; i++ {
// 		curve.insertLast(getRatioPoint(points.list[0], points.list[1], points.list[2], float64(i)*(add)))
// 	}
// 	curve.insertLast(points.list[2])
// 	return curve
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
		pref_x = 550
	} else if dx > 515 {
		pref_x = int(dx) + 170
	}

	if dy < 400 {
		r_y = 400.0 / dy
		pref_y = 550
	} else if dy > 515 {
		pref_y = int(dy) + 170
	}

	if r_x > r_y {
		pref_y = int(r_x * float64(pref_y) / r_y)
		r_y = r_x
	} else if r_y > r_x {
		pref_x = int(r_y * float64(pref_x) / r_x)
		r_x = r_y
	}

	add_x = 85 - min_x*r_x
	add_y = 85 - min_y*r_y

	return pref_x, pref_y, r_x, r_y, add_x, add_y
}

func drawSketch(points, corner BezierPoints, iter int) {
	pref_x, pref_y, r_x, r_y, add_x, add_y := getPreferredDimension(corner)

	for i := 1; i <= iter; i++ {
		step := int(math.Pow(2, float64(iter-i)))

		newSketch := gg.NewContext(pref_x, pref_y)
		newSketch.SetRGB(1, 0, 0)

		newSketch.MoveTo(r_x*points.list[0].x+add_x, -1*(r_y*points.list[0].y+add_y)+float64(pref_y))
		for i := 0; i < points.neff; i += step {
			newSketch.LineTo(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y))
		}
		newSketch.SetLineWidth(2.5)
		newSketch.Stroke()

		newSketch.SetRGB(0.2, 0.4, 1)
		newSketch.MoveTo(r_x*corner.list[0].x+add_x, -1*(r_y*corner.list[0].y+add_y)+float64(pref_y))
		for i := 1; i < corner.neff; i++ {
			newSketch.LineTo(r_x*corner.list[i].x+add_x, -1*(r_y*corner.list[i].y+add_y)+float64(pref_y))
		}
		newSketch.SetLineWidth(1)
		newSketch.Stroke()

		newSketch.SetRGB(0, 1, 0)
		for i := 0; i < points.neff; i += step {
			newSketch.DrawPoint(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y), 1.2)
			newSketch.Stroke()
		}

		newSketch.SetRGB(1, 1, 0.5)
		for i := 0; i < corner.neff; i++ {
			newSketch.DrawPoint(r_x*corner.list[i].x+add_x, -1*(r_y*corner.list[i].y+add_y)+float64(pref_y), 1.2)
			newSketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, corner.list[i].x, corner.list[i].y), r_x*corner.list[i].x+add_x, -1*(r_y*corner.list[i].y+add_y)+float64(pref_y), 0.5, -0.5)
			newSketch.Stroke()
		}

		if err := newSketch.SavePNG("bezier_curve_" + strconv.Itoa(i) + ".png"); err != nil {
			fmt.Println("Error saving PNG:", err)
		}
	}
}

// func drawSketchBruteForce(points, corner BezierPoints) {
// 	pref_x, pref_y, r_x, r_y, add_x, add_y := getPreferredDimension(corner)

// 	newSketch := gg.NewContext(pref_x, pref_y)
// 	newSketch.SetRGB(1, 0, 0)
// 	newSketch.MoveTo(r_x*points.list[0].x+add_x, -1*(r_y*points.list[0].y+add_y)+float64(pref_y))
// 	for i := 0; i < points.neff; i++ {
// 		newSketch.LineTo(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y))
// 	}
// 	newSketch.SetLineWidth(2.5)
// 	newSketch.Stroke()
// 	newSketch.SetRGB(0.2, 0.4, 1)
// 	newSketch.MoveTo(r_x*corner.list[0].x+add_x, -1*(r_y*corner.list[0].y+add_y)+float64(pref_y))
// 	for i := 1; i < corner.neff; i++ {
// 		newSketch.LineTo(r_x*corner.list[i].x+add_x, -1*(r_y*corner.list[i].y+add_y)+float64(pref_y))
// 	}
// 	newSketch.SetLineWidth(1)
// 	newSketch.Stroke()
// 	newSketch.SetRGB(0, 1, 0)
// 	for i := 0; i < points.neff; i++ {
// 		newSketch.DrawPoint(r_x*points.list[i].x+add_x, -1*(r_y*points.list[i].y+add_y)+float64(pref_y), 1.2)
// 		newSketch.Stroke()
// 	}
// 	newSketch.SetRGB(1, 1, 0.5)
// 	for i := 0; i < corner.neff; i++ {
// 		newSketch.DrawPoint(r_x*corner.list[i].x+add_x, -1*(r_y*corner.list[i].y+add_y)+float64(pref_y), 1.2)
// 		newSketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, corner.list[i].x, corner.list[i].y), r_x*corner.list[i].x+add_x, -1*(r_y*corner.list[i].y+add_y)+float64(pref_y), 0.5, -0.5)
// 		newSketch.Stroke()
// 	}
// 	if err := newSketch.SavePNG("bezier_curve_brute_force.png"); err != nil {
// 		fmt.Println("Error saving PNG:", err)
// 	}
// }

func main() {
	// point1 := Point{-50, 350}
	// point2 := Point{30, 30}
	// point3 := Point{260, 340}
	// point4 := Point{360, 189}
	// point5 := Point{230, 110}
	// point6 := Point{430, 50}
	// point7 := Point{550, 190}
	// point8 := Point{445, 300}
	start := time.Now()
	// points := BezierPoints{[]Point{{165, 615}, {265, 115}, {765, 115}, {865, 615}}, 4}
	// points.insertLast(point1, point2, point3)
	// points.insertLast(point1, point2, point3, point4, point5, point6, point7, point8)

	points2 := BezierPoints{[]Point{{1, 13}, {6, 13}, {1, 7}, {5, 5}, {8, 7}, {10, 12}, {14, 6}, {12, 2}, {8, 1}}, 9}
	// fmt.Println(points2)

	curveDnC := points2.findCurve(9)
	points2.drawCurveBruteForce()
	elapsedTime := time.Since(start)

	fmt.Printf("Time Elapsed: %0.3fms", float64(elapsedTime.Milliseconds()))
	drawSketch(curveDnC, points2, 12)
}
