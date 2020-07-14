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

	g.MatrixA = [][]int{
		{0, 3, 0, 0, 0, 0, 1},
		{3, 0, 4, 0, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 5, 0, 0},
		{0, 0, 0, 5, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 1},
		{1, 0, 2, 0, 0, 1, 0},
	}

	g.bellmanFord(0)

	fmt.Println(g.L)
	fmt.Println(g.P)
}

// Graph ...
type Graph struct {
	n       int
	MatrixA [][]int
	MapEdge map[string]int

	P []int
	L []int
}

func (g *Graph) bellmanFord(start int) {
	//Bước khởi tạo giống mã giả
	g.P = make([]int, g.n)
	g.L = make([]int, g.n)

	for i := 0; i < len(g.L); i++ {
		g.L[i] = 999999999999
	}
	g.L[start] = 0

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
				if g.L[j] > g.L[i]+g.MapEdge[key] {

					g.L[j] = g.L[i] + g.MapEdge[key]
					// g.P[j] = i

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
}

// func (g *Graph) bellmanFord(source int) {

// 	// Vì mãng của Phong nhận vào load từ file nên chô nào vô cùng là Phong để số 0.Nên khi vô code ta phải chuyển nó về giá trị max value
// 	for i := 0; i < len(g.MatrixA); i++ {
// 		for j := 0; j < len(g.MatrixA); j++ {
// 			if g.MatrixA[i][j] == 0 {
// 				g.MatrixA[i][j] = 999999999999
// 			}
// 		}
// 	}
// 	//Bước khởi tạo giống mã giả
// 	g.P = make([]int, g.n)
// 	g.L = make([]int, g.n)
// 	for i := 0; i < len(g.L); i++ {
// 		g.L[i] = 999999999999
// 	}
// 	g.L[source] = 0
// 	//Biến xác định lặp n lần không biết trước
// 	stop := false
// 	k := 0
// 	for !stop {
// 		stop = true
// 		k = k + 1
// 		//Đây là phần quan trọng của thuật toán tập trung vào câu if
// 		for i := 0; i < len(g.L); i++ {
// 			for j := 0; j < len(g.L); j++ {
// 				if g.MatrixA[i][j] > 0 && g.MatrixA[i][j] < 999999999999 {
// 					//Gần như đây là công thức nên cũng không khó hiểu phải hông nè.Phần này buộc các bạn phải chạy tay hoặc debug mới hiểu nguyên tắc.Mình hông biết phải nói như thế nào khi viết text nữa,hehe
// 					if g.L[j] > g.L[i]+g.MatrixA[i][j] {
// 						// System.out.println(i + " " + j + " " + g.MatrixA[i][j]);
// 						g.L[j] = g.L[i] + g.MatrixA[i][j]
// 						g.P[j] = i
// 						stop = false
// 					}
// 				}
// 			}
// 		}
// 		//Đây là phần kiểm tra chu trình âm
// 		if k > g.n {
// 			if stop == false {
// 				fmt.Println("Graph have negative cycle")
// 			}
// 			stop = true
// 		}
// 	}
// }
