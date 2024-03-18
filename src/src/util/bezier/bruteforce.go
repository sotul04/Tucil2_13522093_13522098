package bezier

func GetRatioPoint4(points BezierPoints, ratio float64) Point {
	if points.Neff == 2 {
		tmpPoint := Point{(1-ratio)*points.List[0].X + ratio*points.List[1].X, (1-ratio)*points.List[0].Y + ratio*points.List[1].Y}
		return tmpPoint
	}
	
	newpoints := BezierPoints{}
	for i := 1; i < points.Neff; i++ {
		tmpPoint := Point{(1-ratio)*points.List[i-1].X + ratio*points.List[i].X, (1-ratio)*points.List[i-1].Y + ratio*points.List[i].Y}
		newpoints.InsertLast(tmpPoint)
	}
	return GetRatioPoint4(newpoints, ratio)
}

func (bp BezierPoints) DrawCurveBruteForce(n_point int) BezierPoints {
	points := BezierPoints{}

	add := 1 / float64(n_point-1)

	for i := 0; i < n_point; i++ {
		points.InsertLast(GetRatioPoint4(bp, float64(i)*add))
	}

	return points
}
