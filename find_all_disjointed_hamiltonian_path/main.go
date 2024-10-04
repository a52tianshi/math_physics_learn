package main

import "fmt"

var numOfSolution = 0
var overLapMap = make(map[[2]int]bool)

func main() {
	n := 9
	var initPath = make([]int, n+1)
	for i := 0; i <= n; i++ {
		initPath[i] = i % n
	}
	//initPath := []int{0, 1, 2, 3, 4 ... n-1,0}

	allPath := genAllHamiltonianPath(n, []int{0}, [][]int{})
	fmt.Println("allPath: ", len(allPath))
	filteredPath := make([][]int, 0)
	filteredPath = append(filteredPath, initPath)
	for _, path := range allPath {
		if path[1] > path[len(path)-2] { //remove duplicate path
			continue
		}
		if !overLap([][]int{initPath}, path) {
			filteredPath = append(filteredPath, path)
		}
	}
	fmt.Println("filteredPath: ", len(filteredPath))
	//gen overLapMap to speed up check overlap
	for i := range filteredPath {
		for j := i; j < len(filteredPath); j++ {
			if overLap([][]int{filteredPath[i]}, filteredPath[j]) {
				overLapMap[[2]int{i, j}] = true
				overLapMap[[2]int{j, i}] = true
			}
		}
	}
	fmt.Println("init overLapMap: ", len(overLapMap))

	DFS(n, []int{0}, filteredPath)
	fmt.Println("numOfSolution: ", numOfSolution)
}

func DFS(n int, existPathsIndex []int, allPaths [][]int) {
	if len(existPathsIndex) == (n-1)/2 {
		//solution
		fmt.Println("solution")
		for _, index := range existPathsIndex {
			fmt.Println(allPaths[index])
		}
		//fmt.Println("solution index: ", existPathsIndex)
		numOfSolution++
		return
	}
	existPaths := make([][]int, len(existPathsIndex))
	for i, index := range existPathsIndex {
		existPaths[i] = allPaths[index]
	}
	lastPathIndex := existPathsIndex[len(existPathsIndex)-1]
	for i := lastPathIndex + 1; i < len(allPaths); i++ {
		var isOverLap bool
		for _, index := range existPathsIndex {
			if overLapMap[[2]int{index, i}] {
				isOverLap = true
				break
			}
		}
		if isOverLap {
			continue
		}
		if !contains(existPathsIndex, i) && !overLap(existPaths, allPaths[i]) {
			var allPathsIndex = make([]int, len(existPathsIndex))
			copy(allPathsIndex, existPathsIndex)
			DFS(n, append(allPathsIndex, i), allPaths)
		}
	}
}

func genAllHamiltonianPath(n int, trail []int, result [][]int) [][]int {
	if len(trail) == n {
		trail = append(trail, 0) //add 0 to the end of trail
		result = append(result, trail)
		return result
	}
	for i := 0; i < n; i++ {
		if !contains(trail, i) {
			result = genAllHamiltonianPath(n, append(trail, i), result)
		}
	}
	return result
}

func contains(arr []int, ele int) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func overLap(paths [][]int, add []int) bool {
	if len(paths) == 0 {
		return false
	}
	var visited = make(map[[2]int]bool)
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			if path[i] < path[i+1] {
				visited[[2]int{path[i], path[i+1]}] = true
			} else {
				visited[[2]int{path[i+1], path[i]}] = true
			}
		}
	}

	for i := 0; i < len(add)-1; i++ {
		if add[i] < add[i+1] {
			if visited[[2]int{add[i], add[i+1]}] {
				return true
			}
		} else {
			if visited[[2]int{add[i+1], add[i]}] {
				return true
			}
		}
	}
	return false
}
