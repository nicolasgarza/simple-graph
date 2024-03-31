package main

import (
	"fmt"
	"strings"
)

type NodeID int

type Graph struct {
	Nodes         map[NodeID]Node
	Edges         []Edge
	IncomingEdges map[NodeID][]Edge
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
		Nodes:         make(map[NodeID]Node),
		Edges:         []Edge{},
		IncomingEdges: make(map[NodeID][]Edge),
	}
}

func (g *Graph) AddNode(id NodeID, value interface{}) {
	g.Nodes[id] = Node{ID: id, Value: value}
}

func (g *Graph) AddEdge(from, to NodeID, weight float64) {
	_, ok := g.Nodes[from]
	_, ok2 := g.Nodes[to]
	if !ok {
		fmt.Println(fmt.Errorf("'From' node does not exist: %d", from))
		return
	} else if !ok2 {
		fmt.Println(fmt.Errorf("'To' node does not exist: %d", to))
		return
	}
	g.Edges = append(g.Edges, Edge{To: to, From: from, Weight: weight})
	g.IncomingEdges[to] = append(g.IncomingEdges[to], Edge{To: to, From: from, Weight: weight})
}

func (g *Graph) RemoveNode(id NodeID) *Graph {
	delete(g.Nodes, id)

	g.Edges = filterEdges(g.Edges, func(edge Edge) bool {
		return edge.From != id
	})

	incomingEdges := g.IncomingEdges[id]
	for _, edge := range incomingEdges {
		g.Edges = filterEdges(g.Edges, func(e Edge) bool {
			return !(e.From == edge.From && e.To == edge.To)
		})
		g.IncomingEdges[edge.From] = filterEdges(g.IncomingEdges[edge.From], func(e Edge) bool {
			return !(e.To == edge.From && e.From == edge.To)
		})
	}
	delete(g.IncomingEdges, id)

	return g
}

func filterEdges(edges []Edge, predicate func(Edge) bool) []Edge {
	filtered := make([]Edge, 0)
	for _, edge := range edges {
		if predicate(edge) {
			filtered = append(filtered, edge)
		}
	}

	return filtered
}

// func (g *Graph) RemoveEdge(from, to NodeID) *Graph

// func (g *Graph) GetNode(id NodeID) (Node, bool)

// func (g *Graph) GetEdges(id NodeID) []Edge

// func (g *Graph) GetAdjacentNodes(id NodeID) []Node

// func (g *Graph) HasNode(id NodeID) bool

// func (g *Graph) HasEdge(from, to NodeID) bool

func (g *Graph) GetNodeCount() int {
	return len(g.Nodes)
}

func (g *Graph) GetEdgeCount() int {
	numEdges := 0
	for _, edges := range g.IncomingEdges {
		numEdges += len(edges)
	}
	return numEdges
}

// func (g *Graph) DFS(startID NodeID, visit func(Node))

// func (g *Graph) BFS(startID NodeID, visit func(Node))

func (g *Graph) String() string {
	var sb strings.Builder
	sb.WriteString("Graph:\n")

	for nodeID := range g.Nodes {
		sb.WriteString(fmt.Sprintf("Node %d:\n", nodeID))

		outgoingEdges := g.GetOutgoingEdges(nodeID)
		if len(outgoingEdges) == 0 {
			sb.WriteString(" - Outgoing Edges: None\n")
		} else {
			sb.WriteString(" - Outgoing Edges:\n")
			for _, edge := range outgoingEdges {
				sb.WriteString(fmt.Sprintf("   - %d -> %d, Weight: %.2f\n", edge.From, edge.To, edge.Weight))
			}
		}

		incomingEdges := g.GetIncomingEdges(nodeID)
		if len(incomingEdges) == 0 {
			sb.WriteString(" - Incoming Edges: None\n")
		} else {
			sb.WriteString(" - Incoming Edges:\n")
			for _, edge := range incomingEdges {
				sb.WriteString(fmt.Sprintf("   - %d -> %d, Weight: %.2f\n", edge.From, edge.To, edge.Weight))
			}
		}
	}
	return sb.String()
}

func (g *Graph) GetIncomingEdges(id NodeID) []Edge {
	var incomingEdges []Edge
	for _, edge := range g.Edges {
		if edge.To == id {
			incomingEdges = append(incomingEdges, edge)
		}
	}
	return incomingEdges
}

func (g *Graph) GetOutgoingEdges(id NodeID) []Edge {
	var outgoingEdges []Edge
	for _, edge := range g.Edges {
		if edge.From == id {
			outgoingEdges = append(outgoingEdges, edge)
		}
	}

	return outgoingEdges
}

func main() {
	my_graph := NewGraph()
	fmt.Println("Populating graph with nodes: ")
	my_graph.AddNode(1, 4)
	my_graph.AddNode(2, 5)
	my_graph.AddNode(3, 6)
	fmt.Println("\n", my_graph.String())

	fmt.Println("Adding edges: ")
	my_graph.AddEdge(1, 2, 3)
	my_graph.AddEdge(2, 3, 4)
	my_graph.AddEdge(1, 3, 5)
	my_graph.AddEdge(2, 1, 7)
	my_graph.AddEdge(3, 2, 6)
	fmt.Println("\n", my_graph.String())

	fmt.Println("Removing node: ")
	my_graph.RemoveNode(1)
	fmt.Println("\n", my_graph.String())

	fmt.Println("Program done")
}
