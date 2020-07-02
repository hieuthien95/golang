package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := graph{}
	g.readFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/graph-bfs-dfs/input.txt")
	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}
	n, err := strconv.Atoi(strings.Replace(g.input[0], "\r", "", 1))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	g.n = n

	g.makeMapCombinedVertex()
	g.makeListEdge()
	g.makeAdjacencyMatrix()
	g.makeAdjacencyEdge()

	// g.isVisited = make([]bool, g.n)
	fmt.Println("Duyet do thi DFS:")
	g.DFS(0)

	fmt.Println()
	fmt.Println()

	fmt.Println("Duyet do thi BFS:")
	g.BFS(0)
}

// ========================================================================================

type graph struct {
	input []string

	mapCombinedVertex map[string]int
	listEdge          []edge
	adjacencyMatrix   [][]int

	adjacencyEdgeArrayV []int
	adjacencyEdgeArrayE []int

	n int
}

type edge struct {
	from int
	to   int
	// w    int
}

func (g *graph) makeMapCombinedVertex() {
	fmt.Println("makeMapCombinedVertex")

	g.mapCombinedVertex = make(map[string]int)
	// g.adjacencyMatrix = make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	g.adjacencyMatrix[i] = make([]int, n)
	// }

	for i := 1; i <= g.n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		for j, s := range strArr {
			w, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if w != 0 {
				key := fmt.Sprintf("%v-%v", i-1, j)
				g.mapCombinedVertex[key] = w
			}

			// // fill [][]array
			// w := g.adjacencyMatrix[i][j]
			// if w != 0 {
			// 	key := fmt.Sprintf("%v-%v", i, j)
			// 	g.mapCombinedVertex[key] = w
			// }
		}
	}

	// // convert [][]array to map
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		w := g.adjacencyMatrix[i][j]
	// 		if w != 0 {
	// 			key := fmt.Sprintf("%v-%v", i, j)
	// 			g.mapCombinedVertex[key] = w
	// 		}

	// 		w = g.adjacencyMatrix[j][i]
	// 		if w != 0 {
	// 			key := fmt.Sprintf("%v-%v", j, i)
	// 			g.mapCombinedVertex[key] = w
	// 		}
	// 	}
	// }

	fmt.Println(g.mapCombinedVertex)
	fmt.Println()
}

func (g *graph) makeListEdge() {
	fmt.Println("makeListEdge")

	g.listEdge = []edge{}

	for i := 1; i <= g.n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		for j, s := range strArr {
			w, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if w != 0 {
				g.listEdge = append(g.listEdge, edge{i - 1, j})
			}
		}
	}

	fmt.Println(g.listEdge)
	fmt.Println()
}

func (g *graph) makeAdjacencyMatrix() {
	fmt.Println("makeAdjacencyMatrix")

	g.adjacencyMatrix = make([][]int, g.n)
	for i := 0; i < g.n; i++ {
		g.adjacencyMatrix[i] = make([]int, g.n)
	}

	for i := 1; i <= g.n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		for j, s := range strArr {
			w, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			g.adjacencyMatrix[i-1][j] = w
		}
	}

	for i := 0; i < g.n; i++ {
		fmt.Println(g.adjacencyMatrix[i])
	}
	fmt.Println()
}

func (g *graph) makeAdjacencyEdge() {
	fmt.Println("makeAdjacencyEdge")

	index := 0
	for i := 1; i <= g.n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		g.adjacencyEdgeArrayV = append(g.adjacencyEdgeArrayV, index)
		for j, s := range strArr {
			w, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if w != 0 {
				g.adjacencyEdgeArrayE = append(g.adjacencyEdgeArrayE, j)
				index++
			}
		}
	}

	fmt.Println(g.adjacencyEdgeArrayV)
	fmt.Println(g.adjacencyEdgeArrayE)
	fmt.Println()
}

// ========================================================================================

type stackQueueItem struct {
	vertex      int
	pathVisited string
}

// DFS ...
func (g graph) DFS(start int) {
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

			for v := g.n - 1; v >= 0; v-- {
				key := fmt.Sprintf("%v-%v", vertex, v)
				wTmp := g.mapCombinedVertex[key]

				if strings.Contains(pathVisited, fmt.Sprint(v)) == false && wTmp != 0 {
					stack[top] = stackQueueItem{
						vertex:      v,
						pathVisited: pathVisited,
					}
					top++

					fmt.Print(pathVisited)
					fmt.Println(v)
				}
			}
		}

	}
}

// BFS ...
func (g graph) BFS(start int) {
	isQVisited := make([]bool, g.n)
	var queue [100]int
	top := 0
	bottom := 0

	queue[bottom] = start
	isQVisited[start] = true
	fmt.Print(start, " ")

	for top >= bottom {
		vertex := queue[bottom]
		bottom++

		for v := 0; v < g.n; v++ {
			key := fmt.Sprintf("%v-%v", vertex, v)
			wTmp := g.mapCombinedVertex[key]

			if isQVisited[v] == false && wTmp != 0 {
				top++
				queue[top] = v
				isQVisited[v] = true

				fmt.Print(v, " ")
			}

		}
	}
}

// ========================================================================================

func (g *graph) readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	lines := strings.Split(string(b), "\n")
	g.input = lines

	return lines, nil
}

func (g *graph) readLineFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	g.input = lines
	return lines, scanner.Err()
}
