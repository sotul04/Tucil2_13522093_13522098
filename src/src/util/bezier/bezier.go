package bezier

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fogleman/gg"
)

func (bp BezierPoints) FindCurve(maXIter int) BezierPoints {
	result := FindPoints(bp, 0, maXIter)
	result.InsertFirst(bp.List[0])
	result.InsertLast(bp.List[len(bp.List)-1])

	return result
}

func FindPoints(bp BezierPoints, iter, maXIter int) BezierPoints {
	if iter >= maXIter {
		return BezierPoints{}
	}
	left, mid, right := FindMidPoint(bp)
	leftPoints := FindPoints(left, iter+1, maXIter)
	rightPoints := FindPoints(right, iter+1, maXIter)

	mid.InsertBefore(leftPoints)
	mid.InsertAfter(rightPoints)
	return mid
}

func FindMidPoint(bp BezierPoints) (BezierPoints, BezierPoints, BezierPoints) {
	tmpBP := []BezierPoints{}

	left := BezierPoints{}
	left.InsertLast(bp.List[0])

	iter := bp.Neff - 1

	right := BezierPoints{}
	right.InsertFirst(bp.List[iter])

	for i := 0; i < iter; i++ {
		if i == 0 {
			tmpBP = append(tmpBP, BezierPoints{})
			traversal := bp.Neff - 1
			for j := 0; j < traversal; j++ {
				tmpBP[0].InsertLast(MidPoint(bp.List[j], bp.List[j+1]))
			}
		} else {
			tmpBP = append(tmpBP, BezierPoints{})
			traversal := tmpBP[i-1].Neff - 1
			for j := 0; j < traversal; j++ {
				tmpBP[i].InsertLast(MidPoint(tmpBP[i-1].List[j], tmpBP[i-1].List[j+1]))
			}
		}
		left.InsertLast(tmpBP[i].List[0])
		right.InsertFirst(tmpBP[i].List[iter-i-1])
	}
	return left, tmpBP[iter-1], right

}

func DrawSketch(points, corner BezierPoints, iter int) {
	pref_X, pref_Y, r_X, r_Y, add_X, add_Y := GetPreferredDimension(corner)

	for i := 1; i <= iter; i++ {
		step := int(math.Pow(2, float64(iter-i)))

		newSketch := gg.NewContext(pref_X, pref_Y)
		newSketch.SetRGB(1, 0, 0)

		newSketch.MoveTo(r_X*points.List[0].X+add_X, -1*(r_Y*points.List[0].Y+add_Y)+float64(pref_Y))
		for i := 0; i < points.Neff; i += step {
			newSketch.LineTo(r_X*points.List[i].X+add_X, -1*(r_Y*points.List[i].Y+add_Y)+float64(pref_Y))
		}
		newSketch.SetLineWidth(2.5)
		newSketch.Stroke()

		newSketch.SetRGB(1, 0.7, 0.4)
		newSketch.MoveTo(r_X*corner.List[0].X+add_X, -1*(r_Y*corner.List[0].Y+add_Y)+float64(pref_Y))
		for i := 1; i < corner.Neff; i++ {
			newSketch.LineTo(r_X*corner.List[i].X+add_X, -1*(r_Y*corner.List[i].Y+add_Y)+float64(pref_Y))
		}
		newSketch.SetLineWidth(1)
		newSketch.Stroke()

		newSketch.SetRGB(0, 1, 0)
		for i := 0; i < points.Neff; i += step {
			newSketch.DrawPoint(r_X*points.List[i].X+add_X, -1*(r_Y*points.List[i].Y+add_Y)+float64(pref_Y), 1.2)
			newSketch.Stroke()
		}

		newSketch.SetRGB(1, 1, 0.5)
		for i := 0; i < corner.Neff; i++ {
			newSketch.DrawPoint(r_X*corner.List[i].X+add_X, -1*(r_Y*corner.List[i].Y+add_Y)+float64(pref_Y), 1.2)
			newSketch.DrawStringAnchored(fmt.Sprintf("P%d(%0.1f, %0.1f)", i, corner.List[i].X, corner.List[i].Y), r_X*corner.List[i].X+add_X, -1*(r_Y*corner.List[i].Y+add_Y)+float64(pref_Y), 0.5, -0.5)
			newSketch.Stroke()
		}

		if err := newSketch.SavePNG("dummy/dnc/bezier_curve_" + strconv.Itoa(i) + ".png"); err != nil {
			fmt.Println("Error saving PNG:", err)
		}
	}
}
