package SLGFov

import "testing"

func TestBresenham(t *testing.T) {
	ret := Bresenham(1, 2, 99, 99)
	if len(ret) <= 10 {
		t.FailNow()
	}
}
func TestCheckIntersectSquareAndBeeline(t *testing.T) {
	squre := Chart{Point{1, 1}, Point{10, 10}}
	beeline := Chart{Point{2, 2}, Point{5, 2}}
	if !CheckIntersectSquareAndBeeline(squre, beeline) {
		t.FailNow()
	}
}
