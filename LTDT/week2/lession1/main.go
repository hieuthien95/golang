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

	isVisited []bool
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

	find2ways(g)
}

// ==================================================================================================

func find2ways(g graph) string {

	processStart := g.DFS(g.start, g.end)
	waysStart := printWays(processStart, g.start, g.end)

	processEnd := g.DFS(g.end, g.start)
	waysEnd := printWays(processEnd, g.end, g.start)

	fmt.Println()
	fmt.Println(waysStart, " - ", waysEnd)

	if len(waysStart) == 0 || len(waysEnd) == 0 {
		return "0"
	}

	str := ""
	weight := 9999999
	for _, arrS := range waysStart {
		strS := fmt.Sprint(g.start) + " "
		for i := len(arrS) - 1; i >= 0; i-- {
			s := arrS[i]
			strS += fmt.Sprint(s) + " "
		}

		for _, arrE := range waysEnd {
			strE := ""
			for i := len(arrE) - 2; i >= 0; i-- {
				e := arrE[i]
				strE += fmt.Sprint(e) + " "
			}

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

	str = fmt.Sprint(len(waysStart)*len(waysEnd)) + " " + fmt.Sprint(weight) + "\n" + str
	write(str)

	fmt.Println(str)
	return str
}

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

func printWays(prosAll []process, start int, target int) map[int][]int {
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
			// ways = append(ways, viewV(start))
			// fmt.Println()
			// fmt.Print(ways)

			time.Sleep(time.Nanosecond)
			mapOutput[time.Now().Nanosecond()] = ways
		}
	}

	// fmt.Println()
	// fmt.Println(mapOutput)
	return mapOutput
}

// ==================================================================================================

type process struct {
	parent int
	vertex int
	w      int
}

// DFS ...
func (g *graph) DFS(start int, target int) []process {
	// begin = 1 => 0
	start--
	target--

	g.isVisited = make([]bool, g.numberVertex)
	prosAll := []process{}

	var stack [100]int
	var top int

	stack[top] = start
	top++

	// fmt.Println(viewV(start))

	for top != 0 {
		top--
		vertex := stack[top]

		if g.isVisited[vertex] == false {
			// fmt.Print(viewV(vertex), " ")
			g.isVisited[vertex] = true

			for v := g.numberVertex - 1; v >= 0; v-- {
				key := fmt.Sprintf("%v-%v", viewV(vertex), viewV(v))
				gTmp := g.mapCombinedVertex[key]
				// gTmp := graph[vertex][v]

				if g.isVisited[v] == false && gTmp != 0 {
					stack[top] = v

					// fmt.Println(viewV(v))
					top++

					//
					prosAll = append(prosAll, process{
						parent: vertex,
						vertex: v,
						w:      gTmp,
					})
					// if v == target || vertex == target {
					// 	goto OUT_LOOP
					// }
				}
			}
		}

	}
	// OUT_LOOP:

	return prosAll
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
