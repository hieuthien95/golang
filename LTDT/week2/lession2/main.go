package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type graph struct {
	input []string

	mapCombinedVertex map[string]int

	dong int
	cot  int

	startX int
	startY int

	endX int
	endY int

	limitClimb int

	isVisited []bool
}

func main() {
	g := graph{}
	g.readLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week2/lession2/input.txt")
	if len(g.input) == 0 {
		fmt.Println("len=0")
		return
	}
	g.dong, g.cot, g.startX, g.startY, g.endX, g.endY, g.limitClimb = cutNumber(g.input[0])

}

// ==================================================================================================

func (g *graph) calWeight(str string) int {
	weight := 0
	arrPoint := strings.Split(strings.Trim(str, " "), " ")
	for i, s := range arrPoint {
		if i > 0 {
			key := arrPoint[i-1] + "-" + s
			weight += g.mapCombinedVertex[key]
		}
	}

	return weight
}

// ==================================================================================================

type process struct {
	parent int
	vertex int
	w      int
}

// DFS ...
// func (g *graph) DFS(start int, target int) []process {
// 	// begin = 1 => 0
// 	start--
// 	target--

// 	g.isVisited = make([]bool, g.numberVertex)
// 	prosAll := []process{}

// 	var stack [100]int
// 	var top int

// 	stack[top] = start
// 	top++

// 	// fmt.Println(viewV(start))

// 	for top != 0 {
// 		top--
// 		vertex := stack[top]

// 		if g.isVisited[vertex] == false {
// 			// fmt.Print(viewV(vertex), " ")
// 			g.isVisited[vertex] = true

// 			for v := g.numberVertex - 1; v >= 0; v-- {
// 				key := fmt.Sprintf("%v-%v", viewV(vertex), viewV(v))
// 				gTmp := g.mapCombinedVertex[key]
// 				// gTmp := graph[vertex][v]

// 				if g.isVisited[v] == false && gTmp != 0 {
// 					stack[top] = v

// 					// fmt.Println(viewV(v))
// 					top++

// 					//
// 					prosAll = append(prosAll, process{
// 						parent: vertex,
// 						vertex: v,
// 						w:      gTmp,
// 					})
// 					// if v == target || vertex == target {
// 					// 	goto OUT_LOOP
// 					// }
// 				}
// 			}
// 		}

// 	}
// 	// OUT_LOOP:

// 	return prosAll
// }

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
