package bezier

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
