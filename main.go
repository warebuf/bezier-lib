package main

import "fmt"
import "math"

func main() {
	fmt.Println("Hello, Go!")
	IntersectBezierLine()
}

func IntersectBezierLine(bezier []float64, line []float64) bool {

	// Compute line coefficients A, B, C
	A := line[1] - line[3]                 // Sy - Ey
	B := line[2] - line[0]                 // Ex - Sx
	C := line[0]*line[3] - line[2]*line[1] // Sx*Ey - Ex*Sy

	// Compute Bézier coefficients
	x0, y0 := bezier[0], bezier[1]
	x1, y1 := bezier[2], bezier[3]
	x2, y2 := bezier[4], bezier[5]
	x3, y3 := bezier[6], bezier[7]

	a := A*(x3-3*x2+3*x1-x0) + B*(y3-3*y2+3*y1-y0)
	b := A*(3*x2-6*x1+3*x0) + B*(3*y2-6*y1+3*y0)
	c := A*(3*x1-3*x0) + B*(3*y1-3*y0)
	d := A*x0 + B*y0 + C

	// check if there are valid roots between [0, 1]
	// for cubic equation: at^3 + bt^2 + ct + d = 0
	return checkCubicRoots(a, b, c, d)

}

// at^3 + bt^2 + ct + d = 0, if the LHS is ever equal to 0, there is a root (so return true)
func checkCubicRoots(a float64, b float64, c float64, d float64) bool {

	// calculate t = 0, t = 1, check t(0) * t(1) is <= 0, it crosses the x-axis, so there is a root
	t0 := d
	t1 := a + b + c + d

	if t0*t1 <= 0 {
		fmt.Println("t(0)*t(1)<=0")
		return true
	}

	//
	// find the minima/maxima by finding the roots of the derivative
	//

	// take the derivative
	a2 := 3 * a
	b2 := 2 * b
	c2 := c

	// quadratic formula, get the roots
	if (b2*b2)-(4*a2*c2) < 0 { // no real roots
		fmt.Println("derivative has no real roots")
		return false
	}
	tp0_x := (-b2 + math.Sqrt((b2*b2)-(4*a2*c2))) / (2 * a2)
	tp1_x := (-b2 - math.Sqrt((b2*b2)-(4*a2*c2))) / (2 * a2)

	if tp0_x >= 0 && tp0_x <= 1 { // if the first root lies between [0,1]
		tp0_y := a*tp0_x*tp0_x*tp0_x + b*tp0_x*tp0_x + c*tp0_x + d // find the y-value of the first root

		if tp0_y*t0 <= 0 { // if the y-value is a different polarity, there is a root
			fmt.Println("tp0 has different polarity", t0, tp0_y)
			return true
		}
	}

	if tp1_x >= 0 && tp1_x <= 1 { // if the second root lies between [0,1]
		tp1_y := a*tp1_x*tp1_x*tp1_x + b*tp1_x*tp1_x + c*tp1_x + d // find the y-value of the second root

		if tp1_y*t0 <= 0 { // if the y-value is a different polarity, there is a root
			fmt.Println("tp1 has different polarity", t0, tp1_y)
			return true
		}
	}

	fmt.Println("never crosses x-axis")
	return false
}
