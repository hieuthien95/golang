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

	aToB := g.DFS(g.start, g.end)
	bToA := g.DFS(g.end, g.start)

	g.printOutput(aToB, bToA)
}

func (g *graph) printOutput(aToB map[int]stackQueueItem, bToA map[int]stackQueueItem) {

	str := ""
	weight := 999999
	for _, sqItemS := range aToB {
		strS := ""
		arrS := strings.Split(strings.Trim(sqItemS.pathVisited, "-"), "-")
		for i := 0; i < len(arrS); i++ {
			strS += arrS[i] + "-"
		}

		for _, sqItemE := range bToA {
			strE := ""
			arrE := strings.Split(strings.Trim(sqItemE.pathVisited, "-"), "-")
			for i := 0; i < len(arrE); i++ {
				e := arrE[i]
				strE += fmt.Sprint(e) + "-"
			}
			strE += fmt.Sprint(g.start - 1)

			// new line
			str1Line := strS + strE
			str += str1Line + "\n"

			// point
			w := g.calWeight(str1Line)
			if weight > w {
				weight = w
			}
		}

	}

	str = fmt.Sprint(len(aToB)*len(bToA)) + " " + fmt.Sprint(weight) + "\n" + str
	// write(str)

	fmt.Println(str)
}

// ==================================================================================================

type stackQueueItem struct {
	vertex      int
	pathVisited string
}

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
				wTmp := g.mapCombinedVertex[key]

				if strings.Contains(pathVisited, fmt.Sprint(v)) == false && wTmp != 0 {
					sqItem := stackQueueItem{
						vertex:      v,
						pathVisited: pathVisited,
					}
					stack[top] = sqItem
					top++

					//
					if v == target {
						// fmt.Print(pathVisited)
						// fmt.Println(v)

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

func (g *graph) calWeight(str string) int {
	weight := 0
	arrPoint := strings.Split(strings.Trim(str, "-"), "-")
	for i, s := range arrPoint {
		sInt, _ := strconv.Atoi(s)
		if i > 0 {
			s2Int, _ := strconv.Atoi(arrPoint[i-1])
			key := fmt.Sprint(sInt+1) + "-" + fmt.Sprint(s2Int+1)
			weight += g.mapCombinedVertex[key]
		}
	}

	return weight
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
