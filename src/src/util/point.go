package main

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"Y"`
}

func midPoint(p1, p2 Point) Point {
	mid_X := (p1.X + p2.X) / 2
	mid_Y := (p1.Y + p2.Y) / 2
	return Point{mid_X, mid_Y}
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

func getPreferredDimension(corner BezierPoints) (int, int, float64, float64, float64, float64) {
	min_X := corner.list[0].X
	min_Y := corner.list[0].Y

	maX_X := corner.list[0].X
	maX_Y := corner.list[0].Y

	pref_X := 600
	pref_Y := 600

	r_X := float64(1)
	r_Y := float64(1)

	add_X := float64(0)
	add_Y := float64(0)

	for i := 1; i < corner.neff; i++ {
		if min_X > corner.list[i].X {
			min_X = corner.list[i].X
		} else if maX_X < corner.list[i].X {
			maX_X = corner.list[i].X
		}
		if min_Y > corner.list[i].Y {
			min_Y = corner.list[i].Y
		} else if maX_Y < corner.list[i].Y {
			maX_Y = corner.list[i].Y
		}
	}

	dX := maX_X - min_X
	dY := maX_Y - min_Y

	if dX < 400 {
		r_X = 400.0 / dX
		pref_X = 550
	} else if dX > 515 {
		pref_X = int(dX) + 170
	}

	if dY < 400 {
		r_Y = 400.0 / dY
		pref_Y = 550
	} else if dY > 515 {
		pref_Y = int(dY) + 170
	}

	if r_X > r_Y {
		pref_Y = int(r_X * float64(pref_Y) / r_Y)
		r_Y = r_X
	} else if r_Y > r_X {
		pref_X = int(r_Y * float64(pref_X) / r_X)
		r_X = r_Y
	}

	add_X = 85 - min_X*r_X
	add_Y = 85 - min_Y*r_Y

	return pref_X, pref_Y, r_X, r_Y, add_X, add_Y
}
