package main

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

func (g *Graph) AddNode(id NodeID, value interface{}) *Graph {
	newNodes := make(map[NodeID]Node)
	for k, v := range g.Nodes {
		newNodes[k] = v
	}

	newNodes[id] = Node{ID: id, Value: value}
	return &Graph{
		Nodes: newNodes,
		Edges: g.Edges,
	}
}

func (g *Graph) AddEdge(from, to NodeID, weight float64) *Graph {
	newEdges := append(g.Edges, Edge{To: to, From: from, Weight: weight})
	return &Graph{
		Nodes: g.Nodes,
		Edges: newEdges,
	}

}

func (g *Graph) RemoveNode(id NodeID) *Graph

func (g *Graph) RemoveEdge(from, to NodeID) *Graph

func (g *Graph) GetNode(id NodeID) (Node, bool)

func (g *Graph) GetEdges(id NodeID) []Edge

func (g *Graph) GetAdjacentNodes(id NodeID) []Node

func (g *Graph) HasNode(id NodeID) bool

func (g *Graph) HasEdge(from, to NodeID) bool

func (g *Graph) GetNodeCount() int

func (g *Graph) GetEdgeCount() int

func (g *Graph) DFS(startID NodeID, visit func(Node))

func (g *Graph) BFS(startID NodeID, visit func(Node))

func (g *Graph) String() string

func main() {
	my_graph := NewGraph()
	my_graph.AddNode(1, 4)
	my_graph.AddNode(2, 5)
	my_graph.AddNode(3, 6)

	my_graph.AddEdge(1, 2, 3)
	my_graph.AddEdge(2, 3, 4)
	my_graph.AddEdge(1, 3, 5)
}
