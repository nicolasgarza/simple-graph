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
	OutgoingEdges map[NodeID][]Edge
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
		OutgoingEdges: make(map[NodeID][]Edge),
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
	g.OutgoingEdges[from] = append(g.OutgoingEdges[from], Edge{To: to, From: from, Weight: weight})
}

func (g *Graph) RemoveNode(id NodeID) {
	delete(g.Nodes, id)

	g.Edges = filterEdges(g.Edges, func(edge Edge) bool {
		return edge.From != id
	})

	for _, edge := range g.IncomingEdges[id] {
		g.OutgoingEdges[edge.From] = filterEdges(g.OutgoingEdges[edge.From], func(e Edge) bool {
			return e.To != id
		})
	}

	delete(g.OutgoingEdges, id)
	delete(g.IncomingEdges, id)
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

func (g *Graph) RemoveEdge(from, to NodeID) {
	g.Edges = filterEdges(g.Edges, func(e Edge) bool {
		return !(e.From == from && e.To == to)
	})

	g.IncomingEdges[to] = filterEdges(g.IncomingEdges[to], func(e Edge) bool {
		return !(e.From == from && e.To == to)
	})

	g.OutgoingEdges[from] = filterEdges(g.OutgoingEdges[from], func(e Edge) bool {
		return !(e.From == from && e.To == to)
	})
}

func (g *Graph) GetNode(id NodeID) (Node, bool) {
	node, ok := g.Nodes[id]
	if !ok {
		return Node{}, false
	}
	return node, true
}

func (g *Graph) GetAdjacentNodes(id NodeID) []Node {
	n := []Node{}
	for _, edge := range g.GetOutgoingEdges(id) {
		n = append(n, g.Nodes[edge.To])
	}

	return n
}

func (g *Graph) GetNodeCount() int {
	return len(g.Nodes)
}

func (g *Graph) GetEdgeCount() int {
	return len(g.Edges)
}

func (g *Graph) HasNode(id NodeID) bool {
	_, hasNode := g.GetNode(id)
	return hasNode
}

func (g *Graph) HasEdge(from, to NodeID) bool {
	for _, edge := range g.Edges {
		if edge.From == from && edge.To == to {
			return true
		}
	}

	return false
}

func (g *Graph) DFS(startNodeID NodeID, visit func(node Node)) {
	visited := make(map[NodeID]bool)
	stack := []NodeID{startNodeID}

	for len(stack) > 0 {
		nodeID := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		_, ok := visited[nodeID]
		if !ok {
			visited[nodeID] = true
			node, _ := g.GetNode(nodeID)
			visit(node)
			for _, adjNode := range g.GetAdjacentNodes(nodeID) {
				_, ok := visited[adjNode.ID]
				if !ok {
					stack = append(stack, adjNode.ID)
				}
			}
		}
	}
}

func (g *Graph) BFS(startNodeID NodeID, visit func(node Node)) {
	visited := make(map[NodeID]bool)
	queue := []NodeID{startNodeID}

	for len(queue) > 0 {
		nodeID := queue[0]
		queue = queue[1:]

		_, ok := visited[nodeID]
		if !ok {
			visited[nodeID] = true
			node, _ := g.GetNode(nodeID)
			visit(node)
			for _, adjNode := range g.GetAdjacentNodes(nodeID) {
				_, ok := visited[adjNode.ID]
				if !ok {
					queue = append(queue, adjNode.ID)
				}
			}
		}
	}
}

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
	return g.IncomingEdges[id]
}

func (g *Graph) GetOutgoingEdges(id NodeID) []Edge {
	return g.OutgoingEdges[id]
}
