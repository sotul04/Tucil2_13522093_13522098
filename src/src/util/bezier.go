package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fogleman/gg"
)

func (bp BezierPoints) findCurve(maXIter int) BezierPoints {
	result := findPoints(bp, 0, maXIter)
	result.insertFirst(bp.list[0])
	result.insertLast(bp.list[len(bp.list)-1])

	return result
}

func findPoints(bp BezierPoints, iter, maXIter int) BezierPoints {
	if iter >= maXIter {
		return BezierPoints{}
	}
	left, mid, right := findMidPoint(bp)
	leftPoints := findPoints(left, iter+1, maXIter)
	rightPoints := findPoints(right, iter+1, maXIter)

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

func drawSketch(points, corner BezierPoints, iter int) {
	pref_X, pref_Y, r_X, r_Y, add_X, add_Y := getPreferredDimension(corner)

	for i := 1; i <= iter; i++ {
		step := int(math.Pow(2, float64(iter-i)))

		newSketch := gg.NewContext(pref_X, pref_Y)
		newSketch.SetRGB(1, 0, 0)

		newSketch.MoveTo(r_X*points.list[0].X+add_X, -1*(r_Y*points.list[0].Y+add_Y)+float64(pref_Y))
		for i := 0; i < points.neff; i += step {
			newSketch.LineTo(r_X*points.list[i].X+add_X, -1*(r_Y*points.list[i].Y+add_Y)+float64(pref_Y))
		}
		newSketch.SetLineWidth(2.5)
		newSketch.Stroke()

		newSketch.SetRGB(0.2, 0.4, 1)
		newSketch.MoveTo(r_X*corner.list[0].X+add_X, -1*(r_Y*corner.list[0].Y+add_Y)+float64(pref_Y))
		for i := 1; i < corner.neff; i++ {
			newSketch.LineTo(r_X*corner.list[i].X+add_X, -1*(r_Y*corner.list[i].Y+add_Y)+float64(pref_Y))
		}
		newSketch.SetLineWidth(1)
		newSketch.Stroke()

		newSketch.SetRGB(0, 1, 0)
		for i := 0; i < points.neff; i += step {
			newSketch.DrawPoint(r_X*points.list[i].X+add_X, -1*(r_Y*points.list[i].Y+add_Y)+float64(pref_Y), 1.2)
			newSketch.Stroke()
		}

		newSketch.SetRGB(1, 1, 0.5)
		for i := 0; i < corner.neff; i++ {
			newSketch.DrawPoint(r_X*corner.list[i].X+add_X, -1*(r_Y*corner.list[i].Y+add_Y)+float64(pref_Y), 1.2)
			newSketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, corner.list[i].X, corner.list[i].Y), r_X*corner.list[i].X+add_X, -1*(r_Y*corner.list[i].Y+add_Y)+float64(pref_Y), 0.5, -0.5)
			newSketch.Stroke()
		}

		if err := newSketch.SavePNG("bezier/bezier_curve_" + strconv.Itoa(i) + ".png"); err != nil {
			fmt.Println("Error saving PNG:", err)
		}
	}
}
