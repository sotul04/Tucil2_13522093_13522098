package bezier

import (
	"fmt"

	"github.com/fogleman/gg"
)

func GetRatioPoint4(points BezierPoints, sketch *gg.Context, ratio float64, pref_Y int, r_X, r_Y, add_X, add_Y float64) Point {
	if points.Neff == 2 {
		tmpPoint := Point{(1-ratio)*points.List[0].X + ratio*points.List[1].X, (1-ratio)*points.List[0].Y + ratio*points.List[1].Y}
		sketch.SetRGB(0, 1, 0)
		sketch.MoveTo(r_X*points.List[0].X+add_X, -(r_Y*points.List[0].Y+add_Y)+float64(pref_Y))
		sketch.LineTo(r_X*points.List[1].X+add_X, -1*(r_Y*points.List[1].Y+add_Y)+float64(pref_Y))
		sketch.SetLineWidth(0.3)
		sketch.Stroke()
		sketch.SetRGB(1, 1, 0)
		sketch.DrawPoint(r_X*tmpPoint.X+add_X, -1*(r_Y*tmpPoint.Y+add_Y)+float64(pref_Y), 1)
		sketch.Stroke()
		return tmpPoint
	}
	newpoints := BezierPoints{}
	for i := 1; i < points.Neff; i++ {
		sketch.SetRGB(0, 1, 0)
		sketch.MoveTo(r_X*points.List[i-1].X+add_X, -(r_Y*points.List[i-1].Y+add_Y)+float64(pref_Y))
		sketch.LineTo(r_X*points.List[i].X+add_X, -1*(r_Y*points.List[i].Y+add_Y)+float64(pref_Y))
		sketch.SetLineWidth(0.3)
		sketch.Stroke()
		sketch.SetRGB(1, 1, 0)
		tmpPoint := Point{(1-ratio)*points.List[i-1].X + ratio*points.List[i].X, (1-ratio)*points.List[i-1].Y + ratio*points.List[i].Y}
		sketch.DrawPoint(r_X*tmpPoint.X+add_X, -1*(r_Y*tmpPoint.Y+add_Y)+float64(pref_Y), 1)
		sketch.Stroke()
		newpoints.InsertLast(tmpPoint)
	}
	return GetRatioPoint4(newpoints, sketch, ratio, pref_Y, r_X, r_Y, add_X, add_Y)
}

func (bp BezierPoints) DrawCurveBruteForce(n_point int) BezierPoints {
	pref_X, pref_Y, r_X, r_Y, add_X, add_Y := GetPreferredDimension(bp)

	sketch := gg.NewContext(pref_X, pref_Y)

	points := BezierPoints{}

	add := 1 / float64(n_point)
	iterate := n_point+1

	for i := 0; i < iterate; i++ {
		points.InsertLast(GetRatioPoint4(bp, sketch, float64(i)*add, pref_Y, r_X, r_Y, add_X, add_Y))
	}

	sketch.SetRGB(1, 0, 0)
	sketch.MoveTo(r_X*points.List[0].X+add_X, -1*(r_Y*points.List[0].Y+add_Y)+float64(pref_Y))
	for i := 0; i < n_point; i++ {
		sketch.LineTo(r_X*points.List[i].X+add_X, -1*(r_Y*points.List[i].Y+add_Y)+float64(pref_Y))
	}
	sketch.SetLineWidth(2.5)
	sketch.Stroke()

	sketch.SetRGB(0.2, 0.5, 1)
	for i := 0; i < 41; i++ {
		sketch.DrawPoint(r_X*points.List[i].X+add_X, -1*(r_Y*points.List[i].Y+add_Y)+float64(pref_Y), 1.2)
		sketch.Stroke()
	}

	sketch.SetRGB(0.2, 0.4, 1)
	sketch.MoveTo(r_X*bp.List[0].X+add_X, -1*(r_Y*bp.List[0].Y+add_Y)+float64(pref_Y))
	for i := 1; i < bp.Neff; i++ {
		sketch.LineTo(r_X*bp.List[i].X+add_X, -1*(r_Y*bp.List[i].Y+add_Y)+float64(pref_Y))
	}
	sketch.SetLineWidth(1.8)
	sketch.Stroke()

	sketch.SetRGB(1, 1, 0.5)
	for i := 0; i < bp.Neff; i++ {
		sketch.DrawPoint(r_X*bp.List[i].X+add_X, -1*(r_Y*bp.List[i].Y+add_Y)+float64(pref_Y), 1.2)
		sketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, bp.List[i].X, bp.List[i].Y), r_X*bp.List[i].X+add_X, -1*(r_Y*bp.List[i].Y+add_Y)+float64(pref_Y), 0.5, -0.5)
		sketch.Stroke()
	}

	if err := sketch.SavePNG("bezier/bezier_curve_brute_force.png"); err != nil {
		fmt.Println("Error saving PNG:", err)
	}

	return points
}
