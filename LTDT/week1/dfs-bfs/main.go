package main

import (
	"fmt"
	"strings"
)

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

	start := 0
	target := 1

	// ------------------------------------
	isVisited = make([]bool, n)
	fmt.Println("Duyet do thi DFS:")
	DFS(start, target)

	fmt.Println()

	// ------------------------------------
	isVisited = make([]bool, n)
	fmt.Println("Duyet do thi BFS:")
	printWays(
		BFS(start, target),
		target,
	)
}

// ========================================================================================

type stackQueueItem struct {
	vertex      int
	pathVisited string
}

// DFS ...
func DFS(start int, target int) {
	var stack [100]stackQueueItem
	var top int

	stack[top] = stackQueueItem{
		vertex:      start,
		pathVisited: "",
	}
	top++

	for top != 0 {
		top--
		vertex := stack[top].vertex
		pathVisited := stack[top].pathVisited

		if strings.Contains(pathVisited, fmt.Sprint(vertex)) == false {
			// fmt.Print(vertex, " ")
			pathVisited += fmt.Sprint(vertex) + "-"

			for v := n - 1; v >= 0; v-- {
				key := fmt.Sprintf("%v-%v", vertex, v)
				gTmp := mapGraph[key]
				// gTmp := graph[vertex][v]

				if strings.Contains(pathVisited, fmt.Sprint(v)) == false && gTmp != 0 {
					sqItem := stackQueueItem{
						vertex:      v,
						pathVisited: pathVisited,
					}
					stack[top] = sqItem
					top++

					if v == target {
						fmt.Print(pathVisited)
						fmt.Println(v)
					}
				}
			}
		}

	}
}

// ========================================================================================

type process struct {
	parent int
	vertex int
}

// BFS ...
func BFS(start int, target int) []process {
	prosAll := []process{}

	var queue [100]int
	top := 0
	bottom := 0

	queue[bottom] = start
	isVisited[start] = true
	fmt.Print(start, " ")

	for top >= bottom {
		vertex := queue[bottom]
		bottom++
		for v := 0; v < n; v++ {
			key := fmt.Sprintf("%v-%v", vertex, v)
			gTmp := mapGraph[key]
			// gTmp := graph[vertex][v]

			if isVisited[v] == false && gTmp != 0 {
				top++
				queue[top] = v
				isVisited[v] = true

				fmt.Print(v, " ")

				//
				prosAll = append(prosAll, process{
					parent: vertex,
					vertex: v,
				})
				if v == target || vertex == target {
					return prosAll
				}
			}

		}
	}

	return prosAll
}

func printWays(prosAll []process, target int) {
	ways := []int{}
	tmp := target
	for i := len(prosAll) - 1; i >= 0; i-- {
		p := prosAll[i]

		if tmp == p.vertex && p.vertex != 0 {
			tmp = p.parent
			ways = append(ways, tmp)
		}
	}
	fmt.Println()
	fmt.Println("Way: ", ways)
}
