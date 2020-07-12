package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Graph struct {
	Edges []*Edge
	Nodes []*Node
}

type Edge struct {
	Parent *Node
	Child  *Node
	Cost   int
}

type Node struct {
	Name string
}

const Infinity = int(^uint(0) >> 1)

// AddEdge adds an Edge to the Graph
func (g *Graph) AddEdge(parent, child *Node, cost int) {
	edge := &Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

// AddNode adds a Node to the Graph list of Nodes, if the the node wasn't already added
func (g *Graph) AddNode(node *Node) {
	var isPresent bool
	for _, n := range g.Nodes {
		if n == node {
			isPresent = true
		}
	}

	if !isPresent {
		g.Nodes = append(g.Nodes, node)
	}
}

// String returns a string representation of the Graph
func (g *Graph) String() string {
	var s string

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += edge.Parent.Name + " -> " + edge.Child.Name + " = " + strconv.Itoa(edge.Cost)
		s += "\n"
	}
	s += "\n"

	s += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
	}
	s += "\n"

	return s
}

// Dijkstra implements THE Dijkstra algorithm
// Returns the shortest path from startNode to all the other Nodes
func (g *Graph) Dijkstra(startNode *Node) (shortestPathTable string) {

	// First, we instantiate a "Cost Table", it will hold the information:
	// "From startNode, what's is the cost to all the other Nodes?"
	// When initialized, It looks like this:
	// NODE  COST
	//  A     0    // The startNode has always the lowest cost to itself, in this case, 0
	//  B    Inf   // the distance to all the other Nodes are unknown, so we mark as Infinity
	//  C    Inf
	// ...
	costTable := g.NewCostTable(startNode)

	// An empty list of "visited" Nodes. Everytime the algorithm runs on a Node, we add it here
	var visited []*Node

	// A loop to visit all Nodes
	for len(visited) != len(g.Nodes) {

		// Get closest non visited Node (lower cost) from the costTable
		node := getClosestNonVisitedNode(costTable, visited)

		// Mark Node as visited
		visited = append(visited, node)

		// Get Node's Edges (its neighbors)
		nodeEdges := g.GetNodeEdges(node)

		for _, edge := range nodeEdges {

			// The distance to that neighbor, let's say B is the cost from the costTable + the cost to get there (Edge cost)
			// In the first run, the costTable says it's "Infinity"
			// Plus the actual cost, let's say "5"
			// The distance becomes "5"
			distanceToNeighbor := costTable[node] + edge.Cost

			// If the distance above is lesser than the distance currently in the costTable for that neighbor
			if distanceToNeighbor < costTable[edge.Child] {

				// Update the costTable for that neighbor
				costTable[edge.Child] = distanceToNeighbor
			}
		}
	}

	// Make the costTable nice to read :)
	for node, cost := range costTable {
		shortestPathTable += fmt.Sprintf("Distance from %s to %s = %d\n", startNode.Name, node.Name, cost)
	}

	return shortestPathTable
}

// NewCostTable returns an initialized cost table for the Dijkstra algorithm work with
// by default, the lowest cost is assigned to the startNode – so the algorithm starts from there
// all the other Nodes in the Graph receives the Infinity value
func (g *Graph) NewCostTable(startNode *Node) map[*Node]int {
	costTable := make(map[*Node]int)
	costTable[startNode] = 0

	for _, node := range g.Nodes {
		if node != startNode {
			costTable[node] = Infinity
		}
	}

	return costTable
}

// GetNodeEdges returns all the Edges that start with the specified Node
// In other terms, returns all the Edges connecting to the Node's neighbors
func (g *Graph) GetNodeEdges(node *Node) (edges []*Edge) {
	for _, edge := range g.Edges {
		if edge.Parent == node {
			edges = append(edges, edge)
		}
	}

	return edges
}

// getClosestNonVisitedNode returns the closest Node (with the lower cost) from the costTable
// **if the node hasn't been visited yet**
func getClosestNonVisitedNode(costTable map[*Node]int, visited []*Node) *Node {
	type CostTableToSort struct {
		Node *Node
		Cost int
	}
	var sorted []CostTableToSort

	// Verify if the Node has been visited already
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

	// We need the Node with the lower cost from the costTable
	// So it's important to sort it
	// Here I'm using an anonymous struct to make it easier to sort a map
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}

func main() {
	p1 := &Node{Name: "1"}
	p2 := &Node{Name: "2"}
	p3 := &Node{Name: "3"}
	p4 := &Node{Name: "4"}
	p5 := &Node{Name: "5"}
	p6 := &Node{Name: "6"}

	graph := Graph{}
	// 1
	graph.AddEdge(p1, p2, 7)
	graph.AddEdge(p1, p3, 9)
	graph.AddEdge(p1, p6, 14)
	// 2
	graph.AddEdge(p2, p1, 7)
	graph.AddEdge(p2, p3, 10)
	graph.AddEdge(p2, p4, 15)
	// 3
	graph.AddEdge(p3, p1, 9)
	graph.AddEdge(p3, p2, 10)
	graph.AddEdge(p3, p4, 11)
	graph.AddEdge(p3, p6, 2)
	// 4
	graph.AddEdge(p4, p2, 15)
	graph.AddEdge(p4, p3, 11)
	graph.AddEdge(p4, p5, 6)
	// 5
	graph.AddEdge(p5, p4, 6)
	graph.AddEdge(p5, p6, 9)
	// 6
	graph.AddEdge(p6, p1, 14)
	graph.AddEdge(p6, p3, 2)
	graph.AddEdge(p6, p5, 9)

	fmt.Println(graph.Dijkstra(p1))
	fmt.Println(graph.String())
}
