package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var a float64 = 1e-3 // a is max cost of one length by two coordinates

// Alpha is alpha
var Alpha float64 = 0.01

// Min is minimum value of x or y
var Min = -10

// Max is maximum value of x or y
var Max = 10

// Count is count of points
var Count = 50

// MaxIterationCount is a condition to stop counting
var MaxIterationCount = 10000000

func main() {
	fmt.Println("Start")

	var points = fillWithRandomPoints(Min, Max, Count)
	findLowerCost(points, Alpha)
}

func findLowerCost(points [][]float64, alpha float64) {
	var iterationsCount = 0
	var result float64
	fmt.Println("\nPradine kaina:\n", costFunction(points))

	// mazint zingsny
	for iterationsCount < MaxIterationCount && alpha > 1e-10 {
		iterationsCount++
		fValue := costFunction(points)
		count := 0
		gradientsVector := countGradient(points)
		normalizeGradientVector(gradientsVector)
		pointsCopy := makeCopy(points)
		for i, point := range pointsCopy {
			for j, xy := range point {
				pointsCopy[i][j] = float64(xy) - (alpha * gradientsVector[count])
				count++
			}
		}
		fValueNext := costFunction(pointsCopy)
		if fValueNext < fValue {
			points = pointsCopy
			result = fValueNext
		} else {
			alpha = alpha / 2
		}
	}
	fmt.Println("\nIteraciju skaicius: ", iterationsCount)
	fmt.Println("\nGalutinė kaina: \n", result)
	fmt.Println("\nGauti taskai:\n", points)
}

// Function that normalize vector of gradients
func normalizeGradientVector(vector []float64) {
	vectorsLength := countVectorsLength(vector)
	for i, element := range vector {
		vector[i] = element / vectorsLength
	}
}

// Creates slice of counted points gradient
func countGradient(points [][]float64) []float64 {
	var gradients []float64
	for i, point := range points {
		for j := range point {
			gradient := countPointGradient(i, j, 1e-6, points)
			gradients = append(gradients, gradient)
		}
	}
	return gradients
}

// Counts point x or y gradient - where function grows faster
func countPointGradient(xIndex, yIndex int, h float64, points [][]float64) float64 {
	var pointsCopy = makeCopy(points)

	pointsCopy[xIndex][yIndex] += h
	var gradient = (costFunction(pointsCopy) - costFunction(points)) / float64(h)

	return gradient
}

// Counts length of vector with n elements
// Needed for normalize gradient vector
func countVectorsLength(vector []float64) float64 {
	var sum float64 = 0
	for _, element := range vector {
		sum += math.Pow(element, 2)
	}
	return math.Sqrt(sum)
}

// Clones 2d slice
// Needed for counting point gradient
func makeCopy(points [][]float64) [][]float64 {
	pointsCopy := make([][]float64, len(points))

	for i, point := range points {
		pointCopy := make([]float64, len(point))
		copy(pointCopy, point)
		pointsCopy[i] = pointCopy

	}
	return pointsCopy
}

// C(l)= SUM((l-a)^2)
// Counts cost of all lines between points
func costFunction(points [][]float64) float64 {
	var cost float64 = 0
	counter := 0

	for i, point := range points {
		for _, pointNext := range points[i+1:] {
			counter++
			var length = countLength(point, pointNext)
			var temp = length - a
			cost += math.Pow(temp, 2)
		}
	}
	return cost
}

// Counts length by two pints in 2d dimension
func countLength(x1 []float64, x2 []float64) float64 {
	var xDif = x1[0] - x2[0]
	var yDif = x1[1] - x2[1]
	var xDifRaisedToTwo = math.Pow(float64(xDif), 2)
	var yDifRaisedToTwo = math.Pow(float64(yDif), 2)
	var sum = xDifRaisedToTwo + yDifRaisedToTwo
	var length = math.Sqrt(sum)
	return length
}

// Creates 2d slice with random points
func fillWithRandomPoints(min, max, count int) [][]float64 {
	var primaryPoint = []float64{0, 0}
	points := make([][]float64, count)

	for i := range points {
		rand.Seed(time.Now().UTC().UnixNano())
		// rand.Seed(1)
		if i == 0 {
			points[i] = primaryPoint
			continue
		}
		var x = rand.Intn(max-min) + min
		var y = rand.Intn(max-min) + min
		var row = []float64{float64(x), float64(y)}
		points[i] = row
	}
	fmt.Println("\nPradiniai taskai:\n", points)
	return points
}
