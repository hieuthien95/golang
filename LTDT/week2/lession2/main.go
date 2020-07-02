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

	mapEdge   map[string]int
	mapVertex map[string]int
	arrVertex []string

	dong int
	cot  int

	startX int
	startY int

	endX int
	endY int

	limitClimb int
}

func main() {
	g := graph{}
	g.readLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week2/lession2/input.txt")
	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}
	g.dong, g.cot, g.startX, g.startY, g.endX, g.endY, g.limitClimb = cutNumber(g.input[0])
	g.makeMapVertex()
	g.makeMapEdge()

	g.DFS("0.0", "3.3")
}

func (g *graph) makeMapVertex() {
	g.mapVertex = make(map[string]int)

	for i := 0; i < g.dong; i++ {
		line := strings.Replace(g.input[i+1], "\r", "", 1)
		arrStr := strings.Split(line, " ")
		if len(arrStr) < g.cot {
			return
		}
		for j := 0; j < g.cot; j++ {
			key := fmt.Sprintf("%v.%v", i, j)
			g.mapVertex[key], _ = strconv.Atoi(arrStr[j])
			g.arrVertex = append(g.arrVertex, key)
		}
	}

	// fmt.Println(g.mapVertex)
}

func (g *graph) makeMapEdge() {
	g.mapEdge = make(map[string]int)

	for v := range g.mapVertex {
		arr := strings.Split(v, ".")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])

		keyV := fmt.Sprintf("%v.%v", x, y)
		keyV1 := fmt.Sprintf("%v.%v", x+1, y)
		keyV2 := fmt.Sprintf("%v.%v", x-1, y)
		keyV3 := fmt.Sprintf("%v.%v", x, y+1)
		keyV4 := fmt.Sprintf("%v.%v", x, y-1)

		weightV, okV := g.mapVertex[keyV]
		weightV1, okV1 := g.mapVertex[keyV1]
		weightV2, okV2 := g.mapVertex[keyV2]
		weightV3, okV3 := g.mapVertex[keyV3]
		weightV4, okV4 := g.mapVertex[keyV4]
		if okV == false {
			return
		}

		keyE01 := keyV + "-" + keyV1
		keyE02 := keyV + "-" + keyV2
		keyE03 := keyV + "-" + keyV3
		keyE04 := keyV + "-" + keyV4

		if okV1 == true {
			w := weightV1 - weightV
			if weightV <= weightV1 && w <= g.limitClimb {
				g.mapEdge[keyE01] = w
			}
		}
		if okV2 == true {
			w := weightV2 - weightV
			if weightV <= weightV2 && w <= g.limitClimb {
				g.mapEdge[keyE02] = w
			}
		}
		if okV3 == true {
			w := weightV3 - weightV
			if weightV <= weightV3 && w <= g.limitClimb {
				g.mapEdge[keyE03] = w
			}
		}
		if okV4 == true {
			w := weightV4 - weightV
			if weightV <= weightV4 && w <= g.limitClimb {
				g.mapEdge[keyE04] = w
			}
		}
	}

	// fmt.Println(g.mapEdge)
}

// ==================================================================================================

func (g *graph) calWeight(str string) int {
	weight := 0
	arrPoint := strings.Split(strings.Trim(str, " "), " ")
	for i, s := range arrPoint {
		if i > 0 {
			key := arrPoint[i-1] + "-" + s
			weight += g.mapEdge[key]
		}
	}

	return weight
}

// ==================================================================================================

type stackQueueItem struct {
	vertex      string
	pathVisited string
}

// DFS ...
func (g *graph) DFS(start string, target string) map[int]stackQueueItem {
	// // begin = 1 => 0
	// start--
	// target--

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

		if strings.Contains(pathVisited, vertex) == false {
			// fmt.Print(viewV(vertex), " ")
			pathVisited += vertex + " - "

			for i := len(g.arrVertex) - 1; i >= 0; i-- {
				v := g.arrVertex[i]
				key := fmt.Sprintf("%v-%v", vertex, v)
				_, okTmp := g.mapEdge[key]
				// wTmp := graph[vertex][v]

				if strings.Contains(pathVisited, v) == false && okTmp == true {
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

// ==================================================================================================

func cutNumber(line string) (int, int, int, int, int, int, int) {
	line = strings.Replace(line, "\r", "", 1)
	arrStr := strings.Split(line, " ")

	first, _ := strconv.Atoi(arrStr[0])
	second, _ := strconv.Atoi(arrStr[1])
	third, _ := strconv.Atoi(arrStr[2])
	forth, _ := strconv.Atoi(arrStr[3])

	var fifth, sixth, seventh int
	if len(arrStr) == 7 {
		fifth, _ = strconv.Atoi(arrStr[4])
		sixth, _ = strconv.Atoi(arrStr[5])
		seventh, _ = strconv.Atoi(arrStr[6])
	}

	return first, second, third, forth, fifth, sixth, seventh
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

func write(str string) {
	f, err := os.Create("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week2/lession1/output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	w.WriteString(str)
	w.WriteString("\n")

	w.Flush()
}
