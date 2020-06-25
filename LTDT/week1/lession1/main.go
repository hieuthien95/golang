package main

import "fmt"

var mapGraph map[string]int

// var graph [7][7]int

var isVisited []bool
var n int

// ========================================================================================

func main() {
	n = 7

	mapGraph = make(map[string]int)
	mapGraph["0-1"] = 1
	mapGraph["1-2"] = 1
	mapGraph["2-6"] = 1
	mapGraph["6-5"] = 1
	mapGraph["5-4"] = 1
	mapGraph["4-3"] = 1
	mapGraph["0-6"] = 1

	mapGraph["1-0"] = 1
	mapGraph["2-1"] = 1
	mapGraph["6-2"] = 1
	mapGraph["5-6"] = 1
	mapGraph["4-5"] = 1
	mapGraph["3-4"] = 1
	mapGraph["6-0"] = 1

	// graph = make([][]int, n)
	// graph = [][]int{
	// 	{0, 1, 0, 0, 0, 0, 1},
	// 	{1, 0, 1, 0, 0, 0, 0},
	// 	{0, 1, 0, 0, 0, 0, 1},
	// 	{0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 0, 1, 0, 1, 0},
	// 	{0, 0, 0, 0, 1, 0, 1},
	// 	{1, 0, 1, 0, 0, 1, 0},
	// }

	isVisited = make([]bool, n)
	fmt.Println("Duyet do thi DFS:")
	DFS(0)

	fmt.Println()

	isVisited = make([]bool, n)
	fmt.Println("Duyet do thi BFS:")
	BFS(0)
}

// ========================================================================================

// DFS ...
func DFS(s int) {
	var stack [100]int
	var top int

	stack[top] = 2
	top++

	for top != 0 {
		top--
		vertex := stack[top]

		if isVisited[vertex] == false {
			fmt.Print(vertex, " ")
			isVisited[vertex] = true

			for i := n - 1; i >= 1; i-- {
				// cạnh vô hướng
				key1 := fmt.Sprintf("%v-%v", vertex, i)
				key2 := fmt.Sprintf("%v-%v", i, vertex)
				gTmp1 := mapGraph[key1]
				gTmp2 := mapGraph[key2]
				// gTmp := graph[vertex][i]

				if isVisited[i] == false && (gTmp1 != 0 || gTmp2 != 0) {
					stack[top] = i
					top++
				}
			}
		}

	}
}

// ========================================================================================

// BFS ...
func BFS(u int) {
	var queue [100]int

	top := 0
	bottom := 0
	for i := 0; i < n; i++ {
		queue[i] = 0
	}

	queue[bottom] = u
	isVisited[u] = true
	fmt.Print(u, " ")

	for top >= bottom {
		p := queue[bottom]
		bottom++
		for v := 0; v < n; v++ {
			// cạnh vô hướng
			key1 := fmt.Sprintf("%v-%v", p, v)
			key2 := fmt.Sprintf("%v-%v", v, p)
			gTmp1 := mapGraph[key1]
			gTmp2 := mapGraph[key2]
			// gTmp := graph[p][v]

			if isVisited[v] == false && (gTmp1 != 0 || gTmp2 != 0) {
				top++
				queue[top] = v
				isVisited[v] = true
				fmt.Print(v, " ")
			}

		}
	}
}
