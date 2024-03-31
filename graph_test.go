package main

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")
	g.AddNode(2, "Node 2")

	if g.GetNodeCount() != 2 {
		t.Errorf("Expected node count to be 2, got %d", g.GetNodeCount())
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")
	g.AddNode(2, "Node 2")
	g.AddEdge(1, 2, 1.0)

	if g.GetEdgeCount() != 1 {
		t.Errorf("Expected edge count to be 1, got %d", g.GetEdgeCount())
	}
}

func TestRemoveNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")
	g.AddNode(2, "Node 2")
	g.AddEdge(1, 2, 1.0)
	g.RemoveNode(1)

	if g.GetNodeCount() != 1 {
		t.Errorf("Expected node count to be 1, got %d", g.GetNodeCount())
	}

	if g.GetEdgeCount() != 0 {
		t.Errorf("Expected edge count to be 0, got %d", g.GetEdgeCount())
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")
	g.AddNode(2, "Node 2")
	g.AddEdge(1, 2, 1.0)
	g.RemoveEdge(1, 2)

	if g.GetEdgeCount() != 0 {
		t.Errorf("Expected edge count to be 0, got %d", g.GetEdgeCount())
	}
}

func TestGetNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")

	node, ok := g.GetNode(1)
	if !ok {
		t.Errorf("Expected node to exist")
	}

	if node.Value != "Node 1" {
		t.Errorf("Expected node value to be 'Node 1', got '%v'", node.Value)
	}
}

func TestGetAdjacentNodes(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")
	g.AddNode(2, "Node 2")
	g.AddNode(3, "Node 3")
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 1.0)

	adjacentNodes := g.GetAdjacentNodes(1)
	if len(adjacentNodes) != 2 {
		t.Errorf("Expected 2 adjacent nodes, got %d", len(adjacentNodes))
	}
}

func TestHasNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")

	if !g.HasNode(1) {
		t.Errorf("Expected node 1 to exist")
	}

	if g.HasNode(2) {
		t.Errorf("Expected node 2 to not exist")
	}
}

func TestHasEdge(t *testing.T) {
	g := NewGraph()
	g.AddNode(1, "Node 1")
	g.AddNode(2, "Node 2")
	g.AddEdge(1, 2, 1.0)

	if !g.HasEdge(1, 2) {
		t.Errorf("Expected edge from node 1 to node 2 to exist")
	}

	if g.HasEdge(2, 1) {
		t.Errorf("Expected edge from node 2 to node 1 to not exist")
	}
}
