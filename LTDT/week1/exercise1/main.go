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

var mapGraph map[string]int

var numberVertex int
var numberEdge int

var start int
var target int

// ========================================================================================

func main() {
	lines, _ := readLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/exercise1/input.txt")
	mapGraph = makeMapCombinedVertex(lines)

	// ------------------------------------
	fmt.Println("Duyet do thi DFS:")
	DFS(start, target)

	fmt.Println()

	// ------------------------------------
	fmt.Println("Duyet do thi BFS:")
	BFS(start, target)
}

// ========================================================================================

type stackQueueItem struct {
	vertex    int
	isVisited []*int
}

// DFS ...
func DFS(start int, target int) map[int]stackQueueItem {
	// begin = 1 => 0
	start--
	target--

	mapTarget := make(map[int]stackQueueItem)

	var stack [100]stackQueueItem
	var top int

	stack[top] = stackQueueItem{
		vertex:    start,
		isVisited: make([]*int, numberVertex),
	}
	top++

	for top != 0 {
		top--
		vertex := stack[top].vertex
		isVisited := stack[top].isVisited

		if isVisited[vertex] == nil {
			fmt.Print(viewV(vertex), " ")
			isVisited[vertex] = &vertex

			for v := numberVertex - 1; v >= 0; v-- {
				key := fmt.Sprintf("%v-%v", vertex, v)
				gTmp := mapGraph[key]
				// gTmp := graph[vertex][v]

				if isVisited[v] == nil && gTmp != 0 {
					sqItem := stackQueueItem{
						vertex:    v,
						isVisited: isVisited,
					}
					stack[top] = sqItem
					top++

					// fmt.Print(v, ": ")
					// for _, vv := range isVisited {
					// 	if vv == nil {
					// 		continue
					// 	}
					// 	fmt.Print(*vv, "-")
					// }
					// fmt.Println()

					//
					if v == target {
						mapTarget[time.Now().Nanosecond()] = sqItem
					}
				}
			}
		}

	}
	// OUT_LOOP:

	// printWays(prosAll, start+1, target+1, "DFS")
	return mapTarget
}

// ========================================================================================

type process struct {
	parent int
	vertex int
}

// BFS ...
func BFS(start int, target int) []process {
	isQVisited := make([]bool, numberVertex)

	// begin = 1 => 0
	start--
	target--

	prosAll := []process{}

	var queue [100]int
	top := 0
	bottom := 0

	queue[bottom] = start
	isQVisited[start] = true
	fmt.Print(viewV(start), " ")

	for top >= bottom {
		vertex := queue[bottom]
		bottom++
		for v := 0; v < numberVertex; v++ {
			key := fmt.Sprintf("%v-%v", vertex, v)
			gTmp := mapGraph[key]
			// gTmp := graph[vertex][v]

			if isQVisited[v] == false && gTmp != 0 {
				top++
				queue[top] = v
				isQVisited[v] = true

				fmt.Print(viewV(v), " ")

				//
				prosAll = append(prosAll, process{
					parent: vertex,
					vertex: v,
				})
				// if v == target || vertex == target {
				// 	goto OUT_LOOP
				// }
			}

		}
	}
	// OUT_LOOP:

	// printWays(prosAll, start+1, target+1, "BFS")
	return prosAll
}

// ========================================================================================

func readLineFile(path string) ([]string, error) {
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

	return lines, scanner.Err()
}

func makeMapCombinedVertex(lines []string) map[string]int {
	graph := make(map[string]int)

	if len(lines) == 0 {
		fmt.Println("len=0")
		return nil
	}

	numberVertex, numberEdge = cutNumber(lines[0])
	start, target = cutNumber(lines[len(lines)-1])

	// len(lines)-1
	for i := 1; i < numberEdge+1; i++ {
		n1, n2 := cutNumber(lines[i])
		key := fmt.Sprintf("%v-%v", n1-1, n2-1)
		graph[key] = 1
	}

	return graph
}

func cutNumber(line string) (int, int) {
	line = strings.Replace(line, "\r", "", 1)
	arrStr := strings.Split(line, " ")

	first, _ := strconv.Atoi(arrStr[0])
	second, _ := strconv.Atoi(arrStr[1])

	return first, second
}

func viewV(i int) int {
	return i + 1
}
