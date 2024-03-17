package bezier

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"Y"`
}

func MidPoint(p1, p2 Point) Point {
	mid_X := (p1.X + p2.X) / 2
	mid_Y := (p1.Y + p2.Y) / 2
	return Point{mid_X, mid_Y}
}

type BezierPoints struct {
	List []Point
	Neff int
}

func (bp *BezierPoints) InsertBefore(other BezierPoints) {
	bp.List = append(other.List, bp.List...)
	bp.Neff = bp.Neff + other.Neff
}

func (bp *BezierPoints) InsertAfter(other BezierPoints) {
	bp.List = append(bp.List, other.List...)
	bp.Neff = bp.Neff + other.Neff
}

func (bp *BezierPoints) InsertLast(other_point ...Point) {
	bp.List = append(bp.List, other_point...)
	bp.Neff = bp.Neff + len(other_point)
}

func (bp *BezierPoints) InsertFirst(other_point ...Point) {
	bp.List = append(other_point, bp.List...)
	bp.Neff = bp.Neff + len(other_point)
}