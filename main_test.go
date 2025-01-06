package main

import "testing"

func Test1(t *testing.T) {
	IntersectBezierLine([]float64{0, 0, 1, 2, 2, 2, 3, 3}, []float64{4, 0, 4, 4})     // never crosses
	IntersectBezierLine([]float64{0, 0, 1, 2, 2, 2, 3, 3}, []float64{0, 0, 3, 3})     // 1 intersection
	IntersectBezierLine([]float64{0, 0, 1, 2, 2, 3, 4, 3}, []float64{0, 3, 4, 3})     // 2 intersections
	IntersectBezierLine([]float64{0, 0, 1, 3, 3, 1, 4, 4}, []float64{2, 2, 4, 2})     // 3 intersections
	IntersectBezierLine([]float64{0, 0, 1, 1, 2, 2, 3, 3}, []float64{0, 0, 3, 3})     // lies on line
	IntersectBezierLine([]float64{0, 0, 1, 2, 2, 1, 3, 3}, []float64{3, 3, 4, 4})     // intersects on endpoint
	IntersectBezierLine([]float64{0, 0, 1, 1, 3, 3, 4, 4}, []float64{0, 1, 4, 5})     // parallel line
	IntersectBezierLine([]float64{0, 0, 1, 2, 3, 2, 4, 4}, []float64{2, 0, 2, 5})     // vertical line
	IntersectBezierLine([]float64{0, 0, 1, 3, 3, 1, 4, 4}, []float64{0, 2, 4, 2})     // horizontal line
	IntersectBezierLine([]float64{0, 0, 3, 0, 0, 3, 3, 3}, []float64{1.5, 0, 1.5, 3}) // curve loops onto itself

	IntersectBezierLine([]float64{0, 0, 1, 2, 2, 2, 3, 0}, []float64{1, 6, 0, 5})  // testing quadratic bezier (0 intersections)
	IntersectBezierLine([]float64{0, 0, 1, 2, 2, 2, 3, 0}, []float64{1, 6, 0, -5}) // testing quadratic bezier (1 intersection)

}
