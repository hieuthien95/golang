package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type graph struct {
	input []string

	mapCombinedVertex map[string]int

	numberVertex int
	numberEdge   int

	start int
	end   int
}

type stackQueueItem struct {
	vertex      int
	pathVisited string
}

func main() {
	g := graph{}
	g.readLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week2/lession1/input.txt")
	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}
	g.numberVertex, g.numberEdge, _, _ = cutNumber(g.input[0])
	g.start, g.end, _, _ = cutNumber(g.input[len(g.input)-1])
	g.makeMapCombinedVertex()

	g.DFS(g.start, g.end)
	g.DFS(g.end, g.start)
}

// ==================================================================================================

// ==================================================================================================

// DFS ...
func (g *graph) DFS(start int, target int) map[int]stackQueueItem {
	// begin = 1 => 0
	start--
	target--

	mapTarget := make(map[int]stackQueueItem)

	var stack [100]stackQueueItem
	var top int

	stack[top] = stackQueueItem{
		vertex:      start,
		pathVisited: "",
	}
	top++

	// fmt.Println(viewV(start))

	for top != 0 {
		top--
		vertex := stack[top].vertex
		pathVisited := stack[top].pathVisited

		if strings.Contains(pathVisited, fmt.Sprint(vertex)) == false {
			// fmt.Print(viewV(vertex), " ")
			pathVisited += fmt.Sprint(vertex) + "-"

			for v := g.numberVertex - 1; v >= 0; v-- {
				key := fmt.Sprintf("%v-%v", viewV(vertex), viewV(v))
				gTmp := g.mapCombinedVertex[key]
				// gTmp := graph[vertex][v]

				if strings.Contains(pathVisited, fmt.Sprint(v)) == false && gTmp != 0 {
					sqItem := stackQueueItem{
						vertex:      v,
						pathVisited: pathVisited,
					}
					stack[top] = sqItem
					top++

					//
					if v == target {
						fmt.Print(pathVisited)
						fmt.Println(v)

						mapTarget[time.Now().Nanosecond()] = sqItem
					}
				}
			}
		}

	}
	// OUT_LOOP:

	return mapTarget
}

func (g *graph) makeMapCombinedVertex() {
	// fmt.Println("makeMapCombinedVertex")

	g.mapCombinedVertex = make(map[string]int)
	// len(g.input)-1
	for i := 1; i < g.numberEdge+1; i++ {
		x, y, p, q := cutNumber(g.input[i])

		key1 := fmt.Sprintf("%v-%v", x, y)
		key2 := fmt.Sprintf("%v-%v", y, x)
		g.mapCombinedVertex[key1] = p
		g.mapCombinedVertex[key2] = q
	}

	// // Println
	// fmt.Println(g.mapCombinedVertex)
	// fmt.Println()
}

// ==================================================================================================

func cutNumber(line string) (int, int, int, int) {
	line = strings.Replace(line, "\r", "", 1)
	arrStr := strings.Split(line, " ")

	first, _ := strconv.Atoi(arrStr[0])
	second, _ := strconv.Atoi(arrStr[1])

	var third, forth int
	if len(arrStr) == 4 {
		third, _ = strconv.Atoi(arrStr[2])
		forth, _ = strconv.Atoi(arrStr[3])
	}

	return first, second, third, forth
}

func viewV(i int) int {
	return i + 1
}

// ==================================================================================================

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
