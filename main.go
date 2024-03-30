package main

import (
	"fmt"
	"strings"
)

type NodeID int

type Graph struct {
	Nodes map[NodeID]Node
	Edges []Edge
}

type Node struct {
	ID    NodeID
	Value interface{}
}

type Edge struct {
	Weight float64
	To     NodeID
	From   NodeID
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[NodeID]Node),
		Edges: []Edge{},
	}
}

func (g *Graph) AddNode(id NodeID, value interface{}) {
	g.Nodes[id] = Node{ID: id, Value: value}
}

func (g *Graph) AddEdge(from, to NodeID, weight float64) {
	g.Edges = append(g.Edges, Edge{To: to, From: from, Weight: weight})
}

// func (g *Graph) RemoveNode(id NodeID) *Graph

// func (g *Graph) RemoveEdge(from, to NodeID) *Graph

// func (g *Graph) GetNode(id NodeID) (Node, bool)

// func (g *Graph) GetEdges(id NodeID) []Edge

// func (g *Graph) GetAdjacentNodes(id NodeID) []Node

// func (g *Graph) HasNode(id NodeID) bool

// func (g *Graph) HasEdge(from, to NodeID) bool

// func (g *Graph) GetNodeCount() int

// func (g *Graph) GetEdgeCount() int

// func (g *Graph) DFS(startID NodeID, visit func(Node))

// func (g *Graph) BFS(startID NodeID, visit func(Node))

func (g *Graph) String() string {

	type NodeAndWeight struct {
		Node   NodeID
		Weight float64
	}

	nodeMap := make(map[NodeID][]NodeAndWeight)

	for _, node := range g.Nodes {

		nodeMap[node.ID] = []NodeAndWeight{}
	}

	for _, edge := range g.Edges {

		nodeMap[edge.From] = append(nodeMap[edge.From], NodeAndWeight{Node: edge.To, Weight: edge.Weight})
	}

	var sb strings.Builder
	sb.WriteString("Graph:\n")
	for key, values := range nodeMap {

		if len(values) == 0 {
			sb.WriteString("    - Edges: None")
			continue
		}
		for _, node_and_weight := range values {
			sb.WriteString(fmt.Sprintf("    - Edge: %d -> %d, Weight: %.2f\n", int(key), node_and_weight.Node, node_and_weight.Weight))
		}
	}

	return sb.String()

}

func main() {
	my_graph := NewGraph()
	my_graph.AddNode(1, 4)
	my_graph.AddNode(2, 5)
	my_graph.AddNode(3, 6)

	my_graph.AddEdge(1, 2, 3)
	my_graph.AddEdge(2, 3, 4)
	my_graph.AddEdge(1, 3, 5)

	fmt.Println(my_graph.String())
	fmt.Println("Program done")
}
