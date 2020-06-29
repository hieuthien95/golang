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
var isVisited []bool

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

type process struct {
	parent int
	vertex int
}

// DFS ...
func DFS(start int, target int) []process {
	// begin = 1 => 0
	start--
	target--

	isVisited = make([]bool, numberVertex)
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

	printWays(prosAll, start+1, target+1, "DFS")
	return prosAll
}

// ========================================================================================

// BFS ...
func BFS(start int, target int) []process {
	// begin = 1 => 0
	start--
	target--

	isVisited = make([]bool, numberVertex)
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

	printWays(prosAll, start+1, target+1, "BFS")
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

func write(mapOutput map[int][]int, ftype string) {
	f, err := os.Create("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week1/exercise1/output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	if len(mapOutput) == 0 {
		w.WriteString(fmt.Sprint(0))
		w.Flush()
		return
	}

	// DFS min
	if ftype == "DFS" {
		minWays := []int{}
		for _, ways := range mapOutput {
			minWays = ways
			break
		}
		for _, ways := range mapOutput {
			if len(minWays) >= len(ways) {
				minWays = ways
			}
		}

		for i := len(minWays) - 1; i >= 0; i-- {
			w.WriteString(fmt.Sprint(minWays[i]) + " ")
		}
		w.WriteString("\n")
	}

	// DFS / BFS
	for _, ways := range mapOutput {
		for i := len(ways) - 1; i >= 0; i-- {
			w.WriteString(fmt.Sprint(ways[i]) + " ")
		}
		w.WriteString("\n")
	}

	w.Flush()
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

func printWays(prosAll []process, start int, target int, ftype string) {
	// begin = 1 => 0
	start--
	target--

	mapOutput := make(map[int][]int)

	prosCurrent := []process{}
	for _, p := range prosAll {
		prosCurrent = append(prosCurrent, p)

		if p.vertex == target {
			// add target
			ways := []int{viewV(target)}

			tmp := target
			for i := len(prosCurrent) - 1; i >= 0; i-- {
				pp := prosCurrent[i]

				if tmp == pp.vertex && pp.parent != 0 {
					tmp = pp.parent
					ways = append(ways, viewV(tmp))
				}
			}

			// add start
			ways = append(ways, viewV(start))
			fmt.Println()
			fmt.Print(ways)

			time.Sleep(time.Nanosecond)
			mapOutput[time.Now().Nanosecond()] = ways
		}
	}

	fmt.Println()
	fmt.Println(mapOutput)

	// write
	write(mapOutput, ftype)
}
