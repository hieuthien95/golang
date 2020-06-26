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
	DFS(0, 99)

	fmt.Println()

	isVisited = make([]bool, n)
	fmt.Println("Duyet do thi BFS:")
	BFS(0, 99)
}

// ========================================================================================

type process struct {
	parent int
	vertex int
}

// DFS ...
func DFS(start int, target int) []int {
	prosAll := []process{}

	var stack [100]int
	var top int

	stack[top] = start
	top++

	for top != 0 {
		top--
		vertex := stack[top]

		if isVisited[vertex] == false {
			fmt.Print(vertex, " ")
			isVisited[vertex] = true

			for v := n - 1; v >= 0; v-- {
				key := fmt.Sprintf("%v-%v", vertex, v)
				gTmp := mapGraph[key]
				// gTmp := graph[vertex][v]

				if isVisited[v] == false && gTmp != 0 {
					stack[top] = v
					top++

					//
					prosAll = append(prosAll, process{
						parent: vertex,
						vertex: v,
					})
					if v == target || vertex == target {
						goto OUT_LOOP
					}
				}
			}
		}

	}
OUT_LOOP:

	var ways []int
	tmp := target
	for i := len(prosAll) - 1; i >= 0; i-- {
		p := prosAll[i]

		if tmp == p.vertex {
			tmp = p.parent
			ways = append(ways, tmp)
		}
	}
	fmt.Println()
	fmt.Println("Way: ", ways)

	return ways
}

// ========================================================================================

// BFS ...
func BFS(start int, target int) []int {
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
					goto OUT_LOOP
				}
			}

		}
	}
OUT_LOOP:

	var ways []int
	tmp := target
	for i := len(prosAll) - 1; i >= 0; i-- {
		p := prosAll[i]

		if tmp == p.vertex {
			tmp = p.parent
			ways = append(ways, tmp)
		}
	}
	fmt.Println()
	fmt.Println("Way: ", ways)

	return ways
}
