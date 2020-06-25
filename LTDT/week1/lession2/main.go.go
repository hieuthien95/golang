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
	g := Graph{}
	g.readFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/lession2/input.txt")

	g.printMapCombinedVertex()
	g.printListEdge()
	g.printAdjacencyMatrix()
	g.printAdjacencyEdge()

	fmt.Println()
}

// ========================================================================================

// Graph ...
type Graph struct {
	input []string

	mapCombinedVertex map[string]int
	listEdge          []string
	adjMatrix         [][]int
	adjEdge           []string
}

func (g Graph) printMapCombinedVertex() {

	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}

	n, err := strconv.Atoi(strings.Replace(g.input[0], "\r", "", 1))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	graph := make(map[string]int)

	for i := 1; i < n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		for j, s := range strArr {
			m, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if m != 0 {
				key := fmt.Sprintf("%v-%v", i, j+1)
				graph[key] = m
			}
		}
	}

	fmt.Println(graph)
	fmt.Println()
}

func (g Graph) printListEdge() {
	fmt.Println("printListEdge")
	fmt.Println()
}

func (g Graph) printAdjacencyMatrix() {
	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}

	n, err := strconv.Atoi(strings.Replace(g.input[0], "\r", "", 1))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for i := 1; i < n; i++ {
		strLine := strings.Replace(g.input[i], "\r", "", 1)
		strArr := strings.Split(strLine, " ")

		for j, s := range strArr {
			m, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			graph[i-1][j] = m
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(graph[i])
	}
	fmt.Println()
}

func (g Graph) printAdjacencyEdge() {
	fmt.Println("printAdjacencyEdge")
	fmt.Println()
}

// ========================================================================================

func (g *Graph) readFile(fileName string) ([]string, error) {
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

func (g *Graph) readLineFile(path string) ([]string, error) {
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
