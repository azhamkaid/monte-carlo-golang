package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("pilih 1untuk menhitung luas lingkran  untuk menhitung luas elip 3 untk luas kurva")
	var x int
	fmt.Scanln(&x)
	if x == 1 {
		Circle()
		var r float64
		fmt.Printf("insert value of r: ")
		fmt.Scanln(&r)
		area := r * r * Circle()
		fmt.Println(area)
	} else if x == 2 {

		var a, b float64
		fmt.Printf("insert value of a: ")
		fmt.Scanln(&a)

		fmt.Printf("insert value of b: ")
		fmt.Scanln(&b)
		e := elips(a, b)
		fmt.Println(e)

	} else if x == 3 {

		nonlinear()
	}

}

func Circle() float64 {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())
	// Number of random points to generate
	numPoints := 100000
	// Count of points inside the quarter circle
	insideCircle := 0
	for i := 0; i < numPoints; i++ {
		// Generate a random coordinate in float64 type between 0.0 and 1.0
		x := rand.Float64()
		y := rand.Float64()
		// Calculate the distance from the origin
		distance := math.Sqrt(x*x + y*y)

		// Check if the point is inside the quarter circle
		if distance < 1 {
			insideCircle++
		}
	}

	// Estimate the value of Pi
	piEstimation := 4 * float64(insideCircle) / float64(numPoints)

	// Print the result
	fmt.Printf("Estimated value of Pi: %f\n", piEstimation)
	return piEstimation
}

func elips(a, b float64) float64 {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())
	// Number of random points to generate
	numPoints := 100000
	// Count of points inside the quarter elips
	insideElips := 0
	for i := 0; i < numPoints; i++ {
		// Generate a random coordinate in float64 type between 0.0 and 1.0
		x := rand.Float64()
		y := rand.Float64()
		// Calculate the distance from the origin
		distance := math.Sqrt(x*x + y*y)
		// Check if the point is inside the quarter circle
		if distance < 1 {
			insideElips++
		}
	}

	// Estimate the value of Pi
	piEstimation := 4 * a * b * float64(insideElips) / float64(numPoints)

	// Print the result
	fmt.Printf("Estimated area of elips: %f\n", piEstimation)
	return piEstimation
}

func nonlinear() {
	fmt.Println("Y=aX^2+bX+c")
	var a, b, c float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	//y := x*x - 10*x + 24
	x1 := (-b - math.Sqrt(b*b-4*a*c)) / 2
	x2 := (-b + math.Sqrt(b*b-4*a*c)) / 2
	fmt.Println(x1, x2)
	yp := a*(x1+x2)*(x1+x2)/4 + b*(x1+x2)/2 + c
	fmt.Println(yp)

	numPoint := 10000
	undercurve := 0
	for i := 0; i < numPoint; i++ {
		x := rand.Float64()*math.Abs(x2-x1) + math.Min(x1, x2)
		y := rand.Float64() * math.Abs(yp)

		if y < math.Abs(a*x*x+b*x+c) && y > 0 {
			undercurve++
		}
	}
	fmt.Println(math.Min(x1, x2), math.Max(x1, x2))
	fmt.Println(undercurve)
	areaesetimation := float64(undercurve) / float64(numPoint) * math.Abs(x1-x2) * yp
	fmt.Println(areaesetimation)
}
