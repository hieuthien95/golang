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
	g.readFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/lession2/input.txt")

	g.makeMapCombinedVertex()
	g.makeListEdge()
	g.makeAdjacencyMatrix()
	g.makeAdjacencyEdge()

	fmt.Println()

	g.isVisited = make([]bool, g.n)
	fmt.Println("Duyet do thi DFS:")
	g.DFS(0)

	fmt.Println()

	g.isVisited = make([]bool, g.n)
	fmt.Println("Duyet do thi BFS:")
	g.BFS(0)
}

// ========================================================================================

type graph struct {
	input []string

	mapCombinedVertex map[string]int
	listEdge          []edge
	adjacencyMatrix   [][]int
	adjacencyEdge     []string

	n         int
	isVisited []bool
}

type edge struct {
	from int
	to   int
	w    int
}

func (g *graph) makeMapCombinedVertex() {
	fmt.Println("makeMapCombinedVertex")

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

	g.mapCombinedVertex = make(map[string]int)
	// g.adjacencyMatrix = make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	g.adjacencyMatrix[i] = make([]int, n)
	// }

	for i := 1; i <= n; i++ {
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

	g.listEdge = []edge{}

	for i := 1; i <= n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		for j, s := range strArr {
			w, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if w != 0 {
				g.listEdge = append(g.listEdge, edge{i - 1, j, w})
			}
		}
	}

	fmt.Println(g.listEdge)
	fmt.Println()
}

func (g *graph) makeAdjacencyMatrix() {
	fmt.Println("makeAdjacencyMatrix")

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

	g.adjacencyMatrix = make([][]int, n)
	for i := 0; i < n; i++ {
		g.adjacencyMatrix[i] = make([]int, n)
	}

	for i := 1; i <= n; i++ {
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

	for i := 0; i < n; i++ {
		fmt.Println(g.adjacencyMatrix[i])
	}
	fmt.Println()
}

func (g *graph) makeAdjacencyEdge() {
	fmt.Println("makeAdjacencyEdge")
	fmt.Println()
}

// ========================================================================================

// DFS ...
func (g graph) DFS(start int) {
	var stack [100]int
	var top int

	stack[top] = start
	top++

	for top != 0 {
		top--
		vertex := stack[top]

		if g.isVisited[vertex] == false {
			fmt.Print(vertex, " ")
			g.isVisited[vertex] = true

			for i := g.n - 1; i >= 0; i-- {
				key := fmt.Sprintf("%v-%v", vertex, i)
				gTmp := g.mapCombinedVertex[key]
				// gTmp := graph[vertex][i]

				if g.isVisited[i] == false && gTmp != 0 {
					stack[top] = i
					top++
				}
			}
		}

	}
}

// BFS ...
func (g graph) BFS(start int) {
	var queue [100]int

	top := 0
	bottom := 0
	for i := 0; i < g.n; i++ {
		queue[i] = 0
	}

	queue[bottom] = start
	g.isVisited[start] = true
	fmt.Print(start, " ")

	for top >= bottom {
		p := queue[bottom]
		bottom++
		for v := 0; v < g.n; v++ {
			key := fmt.Sprintf("%v-%v", p, v)
			gTmp := g.mapCombinedVertex[key]
			// gTmp := graph[p][v]

			if g.isVisited[v] == false && gTmp != 0 {
				top++
				queue[top] = v
				g.isVisited[v] = true
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
