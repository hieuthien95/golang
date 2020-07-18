package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	g := Graph{}
	g.n = 7

	g.MapEdge = make(map[string]int)
	g.MapEdge["0-1"] = 3
	g.MapEdge["0-6"] = 1
	g.MapEdge["1-0"] = 3
	g.MapEdge["1-2"] = 4
	g.MapEdge["2-1"] = 4
	g.MapEdge["2-6"] = 2
	g.MapEdge["3-4"] = 5
	g.MapEdge["4-3"] = 5
	g.MapEdge["4-5"] = 1
	g.MapEdge["5-4"] = 1
	g.MapEdge["5-6"] = 1
	g.MapEdge["6-0"] = 1
	g.MapEdge["6-2"] = 2
	g.MapEdge["6-5"] = 1

	g.bellmanFord(0)
}

// Graph ...
type Graph struct {
	n int
	// MatrixA [][]int
	MapEdge map[string]int

	Tracking []int
	Weight   []int
}

func (g *Graph) bellmanFord(start int) {
	//Bước khởi tạo giống mã giả
	g.Tracking = make([]int, g.n)
	g.Weight = make([]int, g.n)

	for i := 0; i < len(g.Weight); i++ {
		g.Weight[i] = 999999999999
	}
	g.Weight[start] = 0

	//Biến xác định lặp n lần không biết trước
	stop := false
	k := 0
	for !stop {
		stop = true
		k++

		//Đây là phần quan trọng của thuật toán tập trung vào câu if
		for key := range g.MapEdge {
			i, _ := strconv.Atoi(strings.Split(key, "-")[0])
			j, _ := strconv.Atoi(strings.Split(key, "-")[1])

			if g.MapEdge[key] > 0 && g.MapEdge[key] < 999999999999 {
				if g.Weight[j] > g.Weight[i]+g.MapEdge[key] {

					g.Weight[j] = g.Weight[i] + g.MapEdge[key]
					g.Tracking[j] = i

					stop = false
				}
			}
		}

		//Đây là phần kiểm tra chu trình âm
		if k > g.n {
			if stop == false {
				fmt.Println("Graph have negative cycle")
			}
			stop = true
		}
	}

	g.viewPathBellmanFord(start)
}

func (g *Graph) viewPathBellmanFord(start int) {

	for i := 0; i < g.n; i++ {
		way := ""

		//DÙng mãng P và đỉnh kết thúc để truy suất từng phần tử đã lưu trong P
		vertex := i
		for vertex != start {
			vertex = g.Tracking[vertex]
			way = fmt.Sprint(vertex, "-", way)
		}

		way += fmt.Sprint(i) + fmt.Sprint(" = ", g.Weight[i])
		fmt.Println(way)
	}
}
