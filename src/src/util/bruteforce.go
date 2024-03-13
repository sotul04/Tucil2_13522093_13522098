package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

// func getRatioPoint(p0, p1, p2 Point, ratio float64) Point {
// 	one_min_r := 1 - ratio
// 	end_X := one_min_r*one_min_r*p0.X + 2*one_min_r*ratio*p1.X + ratio*ratio*p2.X
// 	end_Y := one_min_r*one_min_r*p0.Y + 2*one_min_r*ratio*p1.Y + ratio*ratio*p2.Y
// 	return Point{end_X, end_Y}
// }

func getRatioPoint4(points BezierPoints, sketch *gg.Context, ratio float64, pref_Y int, r_X, r_Y, add_X, add_Y float64) Point {
	if points.neff == 2 {
		tmpPoint := Point{(1-ratio)*points.list[0].X + ratio*points.list[1].X, (1-ratio)*points.list[0].Y + ratio*points.list[1].Y}
		sketch.SetRGB(0, 1, 0)
		sketch.MoveTo(r_X*points.list[0].X+add_X, -(r_Y*points.list[0].Y+add_Y)+float64(pref_Y))
		sketch.LineTo(r_X*points.list[1].X+add_X, -1*(r_Y*points.list[1].Y+add_Y)+float64(pref_Y))
		sketch.SetLineWidth(0.3)
		sketch.Stroke()
		sketch.SetRGB(1, 1, 0)
		sketch.DrawPoint(r_X*tmpPoint.X+add_X, -1*(r_Y*tmpPoint.Y+add_Y)+float64(pref_Y), 1)
		sketch.Stroke()
		return tmpPoint
	}
	newpoints := BezierPoints{}
	for i := 1; i < points.neff; i++ {
		sketch.SetRGB(0, 1, 0)
		sketch.MoveTo(r_X*points.list[i-1].X+add_X, -(r_Y*points.list[i-1].Y+add_Y)+float64(pref_Y))
		sketch.LineTo(r_X*points.list[i].X+add_X, -1*(r_Y*points.list[i].Y+add_Y)+float64(pref_Y))
		sketch.SetLineWidth(0.3)
		sketch.Stroke()
		sketch.SetRGB(1, 1, 0)
		tmpPoint := Point{(1-ratio)*points.list[i-1].X + ratio*points.list[i].X, (1-ratio)*points.list[i-1].Y + ratio*points.list[i].Y}
		sketch.DrawPoint(r_X*tmpPoint.X+add_X, -1*(r_Y*tmpPoint.Y+add_Y)+float64(pref_Y), 1)
		sketch.Stroke()
		newpoints.insertLast(tmpPoint)
	}
	return getRatioPoint4(newpoints, sketch, ratio, pref_Y, r_X, r_Y, add_X, add_Y)
}

func (bp BezierPoints) drawCurveBruteForce() BezierPoints {
	pref_X, pref_Y, r_X, r_Y, add_X, add_Y := getPreferredDimension(bp)

	sketch := gg.NewContext(pref_X, pref_Y)

	points := BezierPoints{}

	for i := 0; i < 41; i++ {
		points.insertLast(getRatioPoint4(bp, sketch, float64(i)*0.025, pref_Y, r_X, r_Y, add_X, add_Y))
	}

	sketch.SetRGB(1, 0, 0)
	sketch.MoveTo(r_X*points.list[0].X+add_X, -1*(r_Y*points.list[0].Y+add_Y)+float64(pref_Y))
	for i := 0; i < 40; i++ {
		sketch.LineTo(r_X*points.list[i].X+add_X, -1*(r_Y*points.list[i].Y+add_Y)+float64(pref_Y))
	}
	sketch.SetLineWidth(2.5)
	sketch.Stroke()

	sketch.SetRGB(0.2, 0.5, 1)
	for i := 0; i < 41; i++ {
		sketch.DrawPoint(r_X*points.list[i].X+add_X, -1*(r_Y*points.list[i].Y+add_Y)+float64(pref_Y), 1.2)
		sketch.Stroke()
	}

	sketch.SetRGB(0.2, 0.4, 1)
	sketch.MoveTo(r_X*bp.list[0].X+add_X, -1*(r_Y*bp.list[0].Y+add_Y)+float64(pref_Y))
	for i := 1; i < bp.neff; i++ {
		sketch.LineTo(r_X*bp.list[i].X+add_X, -1*(r_Y*bp.list[i].Y+add_Y)+float64(pref_Y))
	}
	sketch.SetLineWidth(1.8)
	sketch.Stroke()

	sketch.SetRGB(1, 1, 0.5)
	for i := 0; i < bp.neff; i++ {
		sketch.DrawPoint(r_X*bp.list[i].X+add_X, -1*(r_Y*bp.list[i].Y+add_Y)+float64(pref_Y), 1.2)
		sketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, bp.list[i].X, bp.list[i].Y), r_X*bp.list[i].X+add_X, -1*(r_Y*bp.list[i].Y+add_Y)+float64(pref_Y), 0.5, -0.5)
		sketch.Stroke()
	}

	if err := sketch.SavePNG("bezier/bezier_curve_brute_force.png"); err != nil {
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
