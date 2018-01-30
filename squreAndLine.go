package SLGFov

type Point struct {
	X, Y float64
}

type Chart struct {
	Start Point
	End   Point
}

func CheckIntersectSquareAndBeeline(square Chart, beeline Chart) bool {
	linea := Chart{Start: Point{square.Start.X, square.Start.Y}, End: Point{square.End.X, square.Start.Y}}
	lineb := Chart{Start: Point{square.End.X, square.Start.Y}, End: Point{square.End.X, square.End.Y}}
	linec := Chart{Start: Point{square.End.X, square.End.Y}, End: Point{square.Start.X, square.End.Y}}
	lined := Chart{Start: Point{square.Start.X, square.End.Y}, End: Point{square.Start.X, square.Start.Y}}
	intersectPointa := getIntersectPoint(linea.Start, linea.End, beeline.Start, beeline.End)
	intersectPointb := getIntersectPoint(lineb.Start, lineb.End, beeline.Start, beeline.End)
	intersectPointc := getIntersectPoint(linec.Start, linec.End, beeline.Start, beeline.End)
	intersectPointd := getIntersectPoint(lined.Start, lined.End, beeline.Start, beeline.End)
	if !checkPointInSquare(square, intersectPointa) && !checkPointInSquare(square, intersectPointb) &&
		!checkPointInSquare(square, intersectPointc) && !checkPointInSquare(square, intersectPointd) {
		return false
	}
	return true
}
func checkPointInSquare(square Chart, point *Point) bool {
	if point == nil {
		return false
	}
	if point.X < square.Start.X || point.X > square.End.X || point.Y < square.Start.Y || point.Y > square.End.Y {
		return false
	}
	return true
}

func getIntersectPoint(a, b, c, d Point) *Point {
	denominator := (b.Y-a.Y)*(d.X-c.X) - (a.X-b.X)*(c.Y-d.Y)
	if denominator == 0 {
		return nil
	}

	x := ((b.X-a.X)*(d.X-c.X)*(c.Y-a.Y) + (b.Y-a.Y)*(d.X-c.X)*a.X - (d.Y-c.Y)*(b.X-a.X)*c.X) / denominator
	y := -((b.Y-a.Y)*(d.Y-c.Y)*(c.X-a.X) + (b.X-a.X)*(d.Y-c.Y)*a.Y - (d.X-c.X)*(b.Y-a.Y)*c.Y) / denominator
	return &Point{x, y}
}
