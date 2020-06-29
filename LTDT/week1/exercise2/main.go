package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := graph{}
	g.readLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/exercise2/input.txt")
	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}
	g.n, g.m = cutNumber(g.input[0])

	g.makeMapCombinedVertex()
	g.makeListEdge()
	g.makeAdjacencyMatrix()
	g.makeAdjacencyEdge()

	fmt.Println()
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
	m int
}

type edge struct {
	from int
	to   int
	// w    int
}

func (g *graph) makeMapCombinedVertex() {
	fmt.Println("makeMapCombinedVertex")

	g.mapCombinedVertex = make(map[string]int)
	// len(g.input)-1
	for i := 1; i < g.m+1; i++ {
		n1, n2 := cutNumber(g.input[i])
		key := fmt.Sprintf("%v-%v", n1, n2)
		g.mapCombinedVertex[key] = 1
	}

	// Println
	fmt.Println(g.mapCombinedVertex)
	fmt.Println()
}

func (g *graph) makeListEdge() {
	fmt.Println("makeListEdge")

	g.listEdge = []edge{}
	// len(g.input)-1
	for i := 1; i < g.m+1; i++ {
		n1, n2 := cutNumber(g.input[i])
		g.listEdge = append(g.listEdge, edge{n1, n2})
	}

	// Println
	fmt.Println(g.listEdge)
	fmt.Println()
}

func (g *graph) makeAdjacencyMatrix() {
	fmt.Println("makeAdjacencyMatrix")

	g.adjacencyMatrix = make([][]int, g.n)
	for i := 0; i < g.n; i++ {
		g.adjacencyMatrix[i] = make([]int, g.n)
	}

	// len(g.input)-1
	for i := 1; i < g.m+1; i++ {
		n1, n2 := cutNumber(g.input[i])
		g.adjacencyMatrix[n1-1][n2-1] = 1
	}

	// Println
	for i := 0; i < g.n; i++ {
		fmt.Println(g.adjacencyMatrix[i])
	}
	fmt.Println()
}

func (g *graph) makeAdjacencyEdge() {
	fmt.Println("makeAdjacencyEdge")

	index := 0
	for i := 1; i <= g.n; i++ {
		g.adjacencyEdgeArrayV = append(g.adjacencyEdgeArrayV, index)

		for j := 1; j <= g.n; j++ {

			// len(g.input)-1
			for l := 1; l < g.m+1; l++ {
				n1, n2 := cutNumber(g.input[l])
				if n1 == i && n2 == j {
					g.adjacencyEdgeArrayE = append(g.adjacencyEdgeArrayE, j)
					index++
					fmt.Print()
				}
			}
		}
	}

	fmt.Println("V", g.adjacencyEdgeArrayV)
	fmt.Println("E", g.adjacencyEdgeArrayE)
	fmt.Println()
}

// ========================================================================================

func (g *graph) readLineFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		g.input = append(g.input, scanner.Text())
	}

	return g.input, scanner.Err()
}

func cutNumber(line string) (int, int) {
	line = strings.Replace(line, "\r", "", 1)
	arrStr := strings.Split(line, " ")

	first, _ := strconv.Atoi(arrStr[0])
	second, _ := strconv.Atoi(arrStr[1])

	return first, second
}
