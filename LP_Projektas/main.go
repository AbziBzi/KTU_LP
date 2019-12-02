package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var c int64 = 1

var a float64 = 1e-3 // a is max cost of one length by two coordinates

// Alpha is alpha
var Alpha float64 = 0.01

var threadsCounts = []int{1, 5, 10, 15, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200}

var h float64 = 1e-10

// Min is minimum value of x or y
var Min = -10

// Max is maximum value of x or y
var Max = 10

// Counts is count of points
var Counts = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

// MaxIterationCount is a condition to stop counting
var MaxIterationCount = 10000

var grad = make([]float64, 1000)

func main() {

	// Fro whole program run
	// for _, count := range Counts {
	// 	var points = fillWithRandomPoints(Min, Max, count)
	// 	fmt.Println("Points Count:\t", count)

	// 	for _, threadsCount := range threadsCounts {
	// 		fmt.Print(" |", threadsCount, "\t|")
	// 		start := time.Now()

	// 		findLowerCost(threadsCount, points, Alpha)

	// 		elapsed := time.Since(start)
	// 		fmt.Println(elapsed, "\t|")
	// 	}
	// 	fmt.Println()
	// }

	// For one part of program run
	points := fillWithRandomPoints(Min, Max, Counts[3])
	fmt.Println("Points Count:\t", Counts[3])
	start := time.Now()
	findLowerCost(threadsCounts[4], points, Alpha)
	elapsed := time.Since(start)
	fmt.Println(elapsed, "\t|")

}

func findLowerCost(threadsCount int, points [][]float64, alpha float64) {
	var iterationsCount = 0
	var result float64
	fmt.Print(costFunction(points), "\t|")

	// mazint zingsny
	for iterationsCount < MaxIterationCount && alpha > 1e-5 {
		iterationsCount++
		fValue := costFunction(points)
		count := 0
		executeGradient(threadsCount, points)
		normalizeGradientVector(grad)
		pointsCopy := makeCopy(points)
		for i, point := range pointsCopy {
			for j, xy := range point {
				pointsCopy[i][j] = float64(xy) - (alpha * grad[count])
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
	fmt.Print(result, "\t|")
	fmt.Print(iterationsCount, "\t|")
}

// Function that normalize vector of gradients
func normalizeGradientVector(vector []float64) {
	vectorsLength := countVectorsLength(vector)
	for i, element := range vector {
		vector[i] = element / vectorsLength
	}
}

func executeGradient(threadsCount int, points [][]float64) {
	wg := sync.WaitGroup{}

	var index int = 0

	// create channels
	send := make(chan []int)

	wg.Add(threadsCount)
	for i := 0; i < threadsCount; i++ {
		go countPointGradient(points, send, &wg)
	}

	for i, point := range points {
		for j := range point {
			send <- []int{index, i, j}
			index++
		}
	}
	close(send)
	wg.Wait()
}

func countPointGradient(points [][]float64, pointsChanel <-chan []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for xy := range pointsChanel {
		if xy != nil && xy[0] < 2 {
			grad[xy[0]] = 0
		} else if xy != nil {
			var pointsCopy = makeCopy(points)
			pointsCopy[xy[1]][xy[2]] += h
			grad[xy[0]] = (costFunction(pointsCopy) - costFunction(points)) / float64(h)
		}
	}
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
		// rand.Seed(time.Now().UTC().UnixNano())
		rand.Seed(c)
		c++
		if i == 0 {
			points[i] = primaryPoint
			continue
		}
		var x = rand.Intn(max-min) + min
		var y = rand.Intn(max-min) + min
		var row = []float64{float64(x), float64(y)}
		points[i] = row
	}
	return points
}
