package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)
type Trees struct {
	Trees []Tree `json:"trees"`
}

type Tree struct {
	Type string `json:"type"`
	Age int `json:"age"`
	Height float64 `json:"height_m"`
	Result int
}

var wg = sync.WaitGroup{}
var FilterValue int = 92941
var ThreadsCount int = 1
var FilePath string = "data/IFF72_ZubowiczE_L1_dat_1.json"
var ResultPath string = "data/IFF72_ZubowiczE_L1_rez.txt"

func main() {
	// Reads file and creates slice of tree struct
	var trees = ReadJsonFile(FilePath)
	var resultTrees []Tree

	// Creates channels in which calculations will be done
	worker := make(chan Tree)
	receive := make(chan Tree)

	// return all trees from chanel
	for i := 0; i < ThreadsCount; i++{
		wg.Add(1)
		go Execute(worker, receive)
	}

	wg.Add(1)
	// result printing function
	go func(chanel <-chan Tree) {
		for element := range chanel{
			resultTrees = append(resultTrees, element)
		}
		wg.Done()
	}(receive)

	// Add trees to chanel
	for _, tree := range trees{
		worker <- tree
	}
	wg.Wait()
	WriteResultToFile(ResultPath, resultTrees, trees)
}

func Execute(chanel <-chan Tree, chanel2 chan<- Tree) {
	defer wg.Done()
	for element := range chanel {
		tree := element
		var value = FindPrimeNumber(tree)
		if value <= FilterValue {
			tree.Result = value
			chanel2 <- tree
		}
	}
}

// Calculations needed for filter trees
func FindPrimeNumber(tree Tree) int{
	var strLen = len(tree.Type)
	var n = (int(tree.Height) * tree.Age * strLen)/2
	var count int = 0
	var a int = 2

	for count < n {
		var b int = 2
		var prime int = 1
		for b * b <= a {
			if a % b == 0 {
				prime = 0
				break
			}
			b++
		}
		if prime > 0 {
			count++
		}
		a++
	}
	return a-1
}

// Write results to result File
func WriteResultToFile(resultPath string, resTrees []Tree, primTrees []Tree){
	file, err := os.Create(resultPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString("Pradiniai duomenys:\n")
	for _, tree := range primTrees{
		s := fmt.Sprintf("| %-10s | %10d | %10.2f |", tree.Type, tree.Age, tree.Height)
		_, err := fmt.Fprintln(file, s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	_, err = file.WriteString("\nSurušiuoti duomenys:\n")
	for _, tree := range resTrees{
		s := fmt.Sprintf("| %-10s | %10d | %10.2f | %10d |", tree.Type, tree.Age, tree.Height, tree.Result)
		_, err := fmt.Fprintln(file, s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

/// Reading from Json file, and adding data to array
func ReadJsonFile(filePath string) []Tree {
	file, _ := ioutil.ReadFile(filePath)
	trees := Trees{}
	_ = json.Unmarshal([]byte(file), &trees)
	return trees.Trees
}
