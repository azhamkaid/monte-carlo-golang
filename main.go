package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	Circle()
	var r float64
	fmt.Printf("masukan nilai r")
	fmt.Scanln(&r)
	area := r * r * Circle()
	fmt.Println(area)
	var a, b float64
	fmt.Printf("masukan nilai a: ")
	fmt.Scanln(&a)

	fmt.Printf("masukan nilai b: ")
	fmt.Scanln(&b)
	e := elips(a, b)
	fmt.Println(e)
}

func Circle() float64 {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())
	// Number of random points to generate
	numPoints := 1000000
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
	numPoints := 10000
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
