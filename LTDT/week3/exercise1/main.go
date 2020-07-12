package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	graph := Graph{}
	graph.ReadLineFile("/Users/thienbui/Documents/Learn/git-hieuthien95/golang/LTDT/week3/exercise1/input.txt")

	fmt.Println(graph.Dijkstra())
}

// Graph ...
type Graph struct {
	input []string

	MapEdges map[string]int
	MapNodes map[string]int

	startNode string
	endNode   string
}

// AddEdge adds an Edge to the Graph
func (g *Graph) AddEdge(from, to string, cost int) {
	// edge
	g.MapEdges[from+"-"+to] = cost
}

// Dijkstra implements THE Dijkstra algorithm
// Returns the shortest path from startNode to all the other Nodes
func (g *Graph) Dijkstra() (shortestPathTable string, mapTracking map[string][]string) {
	mapTracking = make(map[string][]string)

	// First, we instantiate a "Cost Table", it will hold the information:
	// "From startNode, what's is the cost to all the other Nodes?"
	// When initialized, It looks like this:
	// NODE  COST
	//  A     0    // The startNode has always the lowest cost to itself, in this case, 0
	//  B    Inf   // the distance to all the other Nodes are unknown, so we mark as Infinity
	//  C    Inf
	// ...
	costTable := g.MapNodes

	// An empty list of "visited" Nodes. Everytime the algorithm runs on a string, we add it here
	var visited []string

	// A loop to visit all Nodes
	for len(visited) != len(g.MapNodes) {

		// Get closest non visited string (lower cost) from the costTable
		node := getNodeClosestNotVisited(costTable, visited)

		// Mark string as visited
		visited = append(visited, node)

		// Get string's Edges (its neighbors)
		nodeEdges := g.GetEdgesOfNode(node)

		for _, edge := range nodeEdges {

			// The distance to that neighbor, let's say B is the cost from the costTable + the cost to get there (Edge cost)
			// In the first run, the costTable says it's "Infinity"
			// Plus the actual cost, let's say "5"
			// The distance becomes "5"
			distanceToNeighbor := costTable[node] + g.MapEdges[edge]

			// If the distance above is lesser than the distance currently in the costTable for that neighbor
			arrNodeOfEdge := strings.Split(edge, "-")
			to := arrNodeOfEdge[1]
			if distanceToNeighbor < costTable[to] {
				// Update the costTable for that neighbor
				costTable[to] = distanceToNeighbor

				// tracking
				mapTracking[to] = append(mapTracking[node], node)
			}
		}
	}

	// Make the costTable nice to read :)
	for node, cost := range costTable {
		shortestPathTable += fmt.Sprintf("Distance from %s to %s = %d\n", g.startNode, node, cost)
	}

	return shortestPathTable, mapTracking
}

// GetEdgesOfNode returns all the Edges that start with the specified string
// In other terms, returns all the Edges connecting to the string's neighbors
func (g *Graph) GetEdgesOfNode(node string) (edges []string) {
	for edge := range g.MapEdges {
		arrNodeOfEdge := strings.Split(edge, "-")
		from := arrNodeOfEdge[0]
		if from == node {
			edges = append(edges, edge)
		}
	}

	return edges
}

// getNodeClosestNotVisited returns the closest string (with the lower cost) from the costTable
// **if the node hasn't been visited yet**
func getNodeClosestNotVisited(costTable map[string]int, visited []string) string {
	type CostTableToSort struct {
		Node string
		Cost int
	}
	var sorted []CostTableToSort

	// Verify if the string has been visited already
	for node, cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			sorted = append(sorted, CostTableToSort{node, cost})
		}
	}

	// We need the string with the lower cost from the costTable
	// So it's important to sort it
	// Here I'm using an anonymous struct to make it easier to sort a map
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}

// ReadLineFile ...
func (g *Graph) ReadLineFile(path string) ([]string, error) {
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

	g.MakeGraph()
	return g.input, scanner.Err()
}

// MakeGraph ...
func (g *Graph) MakeGraph() {
	g.MapEdges = make(map[string]int)
	g.MapNodes = make(map[string]int)

	n, m, _ := cutNumber(g.input[0])
	s, e, _ := cutNumber(g.input[(len(g.input) - 1)])

	// add Node
	for i := 1; i <= n; i++ {
		g.MapNodes[fmt.Sprint(i)] = 999999999999
	}

	// start/end
	g.startNode = fmt.Sprint(s)
	g.MapNodes[fmt.Sprint(s)] = 0
	g.endNode = fmt.Sprint(e)

	// add edge
	for i := 1; i < m+1; i++ {
		n1, n2, n3 := cutNumber(g.input[i])
		g.AddEdge(fmt.Sprint(n1), fmt.Sprint(n2), n3)
	}
}

func cutNumber(line string) (int, int, int) {
	line = strings.Replace(line, "\r", "", 1)
	arrStr := strings.Split(line, " ")

	first, _ := strconv.Atoi(arrStr[0])
	second, _ := strconv.Atoi(arrStr[1])
	var third int
	if len(arrStr) == 3 {
		third, _ = strconv.Atoi(arrStr[2])
	}

	return first, second, third
}
