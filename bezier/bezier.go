package bezier

import "fmt"
import "math"

func IntersectBezierLine(bezier []float64, line []float64) bool {

	fmt.Println("?", bezier, line)

	// Compute line coefficients A, B, C
	A := line[1] - line[3]                         // Y1-Y2
	B := line[2] - line[0]                         // X2-X1
	C := (line[3] * line[0]) - (line[2] * line[1]) // Y2*X1 - X2*Y1

	fmt.Println("A:", A, "B:", B, "C:", C)

	// Compute Bézier coefficients
	x0, y0 := bezier[0], bezier[1]
	x1, y1 := bezier[2], bezier[3]
	x2, y2 := bezier[4], bezier[5]
	x3, y3 := bezier[6], bezier[7]

	a := A*(x3-3*x2+3*x1-x0) + B*(y3-3*y2+3*y1-y0)
	b := A*(3*x2-6*x1+3*x0) + B*(3*y2-6*y1+3*y0)
	c := A*(3*x1-3*x0) + B*(3*y1-3*y0)
	d := C

	fmt.Println("a,b,c,d", a, b, c, d)

	// if a == 0, corner case, checkQuadraticRoots
	if a == 0 {
		if b == 0 {
			return checkLinearRoots(c, d)
		}
		return checkQuadraticRoots(b, c, d)
	}

	// check if there are valid roots between [0, 1]
	// for cubic equation: at^3 + bt^2 + ct + d = 0
	return checkCubicRoots(a, b, c, d, bezier, line)

}

// at^3 + bt^2 + ct + d = 0, if the LHS is ever equal to 0, there is a root (so return true)
func checkCubicRoots(a float64, b float64, c float64, d float64, bezier []float64, line []float64) bool {

	// calculate t = 0, t = 1, check t(0) * t(1) is <= 0, it crosses the x-axis, so there is a root
	t0 := d
	t1 := a + b + c + d

	if t0*t1 <= 0 {
		fmt.Println("t(0)*t(1)<=0")
		return true
	}

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
			fmt.Println("tp0 has different polarity", t0, tp0_x, tp0_y)
			calcXY(bezier, tp0_x, line)
			return true
		}
	}

	if tp1_x >= 0 && tp1_x <= 1 { // if the second root lies between [0,1]
		tp1_y := a*tp1_x*tp1_x*tp1_x + b*tp1_x*tp1_x + c*tp1_x + d // find the y-value of the second root

		if tp1_y*t0 <= 0 { // if the y-value is a different polarity, there is a root
			fmt.Println("tp1 has different polarity", t0, tp0_x, tp1_y)
			return true
		}
	}

	fmt.Println("never crosses x-axis")
	return false
}

// bt^2 + ct + d = 0
func checkQuadraticRoots(b float64, c float64, d float64) bool {

	t0 := d
	t1 := b + c + d

	if t0*t1 <= 0 {
		fmt.Println("quadratic root crosses 0 between [0,1]")
		return true
	}

	// take derivative to find turning point
	b2 := 2 * b
	c2 := c
	tp_x := -c2 / b2
	tp_y := b*tp_x*tp_x + c*tp_x + d

	if (tp_x >= 0) && (tp_x <= 1) { // there is a turning point between 0 and 1, otherwise it for sure doesn't cross the x-axis
		// if turning point is different polarity, return true, it crosses the root
		if tp_y*t0 <= 0 {
			fmt.Println("quadratic root check crosses 0")
			return true
		}
	}
	fmt.Println("quadratic root does not cross 0")
	return false

}

func checkLinearRoots(c float64, d float64) bool {
	t0 := d
	t1 := c + d

	if t0*t1 <= 0 {
		fmt.Println("linear root check crosses 0")
		return true
	}
	fmt.Println("linear root does not cross 0")
	return false
}

func calcXY(bezier []float64, t float64, line []float64) {

	X1 := (1 - t) * (1 - t) * (1 - t) * bezier[0]
	X2 := (1 - t) * (1 - t) * (t) * bezier[2] * 3
	X3 := (1 - t) * (t) * (t) * bezier[4] * 3
	X4 := (t) * (t) * (t) * bezier[6]

	X := X1 + X2 + X3 + X4

	Y1 := (1 - t) * (1 - t) * (1 - t) * bezier[1]
	Y2 := (1 - t) * (1 - t) * (t) * bezier[3] * 3
	Y3 := (1 - t) * (t) * (t) * bezier[5] * 3
	Y4 := (t) * (t) * (t) * bezier[7]

	Y := Y1 + Y2 + Y3 + Y4

	fmt.Println("X:", X, "Y:", Y)

	calcSlope(line, X)

}

func calcSlope(line []float64, X float64) {

	slope := (line[3] - line[1]) / (line[2] - line[0])

	fmt.Println("slope:", slope)

	y := line[3] + (slope * (X - line[2]))

	fmt.Println("line:", "x", X, "y:", y)

}
