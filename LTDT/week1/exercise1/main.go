package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var mapGraph map[string]int
var isVisited []bool

var numberVertex int
var numberEdge int

var start int
var target int

// ========================================================================================

func main() {
	lines, _ := readLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/exercise1/input.txt")
	mapGraph = makeMapCombinedVertex(lines)

	isVisited = make([]bool, numberVertex)
	fmt.Println("Duyet do thi DFS:")
	DFS(start, target)

	fmt.Println()

	isVisited = make([]bool, numberVertex)
	fmt.Println("Duyet do thi BFS:")
	BFS(start, target)
}

// ========================================================================================

type process struct {
	parent int
	vertex int
}

// DFS ...
func DFS(start int, target int) []int {
	// begin = 1 => 0
	start--
	target--

	prosAll := []process{}

	var stack [100]int
	var top int

	stack[top] = start
	top++

	for top != 0 {
		top--
		vertex := stack[top]

		if isVisited[vertex] == false {
			fmt.Print(viewV(vertex), " ")
			isVisited[vertex] = true

			for v := numberVertex - 1; v >= 0; v-- {
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
					// if v == target || vertex == target {
					// 	goto OUT_LOOP
					// }
				}
			}
		}

	}
	// OUT_LOOP:

	ways := []int{}
	tmp := target
	for i := len(prosAll) - 1; i >= 0; i-- {
		p := prosAll[i]

		if tmp == p.vertex {
			tmp = p.parent
			ways = append(ways, viewV(tmp))
		}
	}
	fmt.Println()
	fmt.Println("Way: ", ways)

	return ways
}

// ========================================================================================

// BFS ...
func BFS(start int, target int) []int {
	// begin = 1 => 0
	start--
	target--

	prosAll := []process{}

	var queue [100]int
	top := 0
	bottom := 0

	queue[bottom] = start
	isVisited[start] = true
	fmt.Print(viewV(start), " ")

	for top >= bottom {
		vertex := queue[bottom]
		bottom++
		for v := 0; v < numberVertex; v++ {
			key := fmt.Sprintf("%v-%v", vertex, v)
			gTmp := mapGraph[key]
			// gTmp := graph[vertex][v]

			if isVisited[v] == false && gTmp != 0 {
				top++
				queue[top] = v
				isVisited[v] = true

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

	ways := []int{}
	tmp := target
	for i := len(prosAll) - 1; i >= 0; i-- {
		p := prosAll[i]

		if tmp == p.vertex {
			tmp = p.parent
			ways = append(ways, viewV(tmp))
		}
	}
	fmt.Println()
	fmt.Println("Way: ", ways)

	return ways
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

	for i := 1; i < len(lines)-1; i++ {
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
