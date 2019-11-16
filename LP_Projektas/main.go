package main

import (
	"fmt"
	"math"
	"math/rand"
)

var a = 5 // a is max cost of one length by two coordinates

// MIN is minimum value of x or y
var MIN = -10

// MAX is maximum value of x or y
var MAX = 10

// COUNT is count of points
var COUNT = 100

func main() {
	fmt.Println("Start")

	var points = fillWithRandomPoints()
	fmt.Println(points)
}

// C(l)= SUM((l-a)^2)
func costFunction(points [][]int) float64 {
	var result float64 = 0

	return result
}

// Count length by two pints in 2d dimension
func countLength(x1 []int, x2 []int) float64 {
	var xDif = x1[0] - x2[0]
	var yDif = x1[1] - x2[1]
	var xDifRaisedToTwo = math.Pow(float64(xDif), 2)
	var yDifRaisedToTwo = math.Pow(float64(yDif), 2)
	var sum = xDifRaisedToTwo + yDifRaisedToTwo
	var length = math.Sqrt(sum)
	return length
}

func fillWithRandomPoints() [][]int {
	var primaryPoint = []int{0, 0}
	points := make([][]int, COUNT)

	for i := range points {
		if i == 0 {
			points[i] = primaryPoint
			continue
		}
		var x = rand.Intn(MAX-MIN) + MIN
		var y = rand.Intn(MAX-MIN) + MIN
		var row = []int{x, y}
		points[i] = row

	}
	return points
}
