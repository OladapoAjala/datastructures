package graph

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/queues/queue"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sequences/node"
)

type Graph[V comparable] struct {
	Vertices *linkedlist.LinkedList[*vertex.Vertex[V]]
}

func NewGraph[V comparable]() *Graph[V] {
	return &Graph[V]{
		Vertices: linkedlist.NewList[*vertex.Vertex[V]](),
	}
}

func (g *Graph[V]) BreadthFirstSearch(start *vertex.Vertex[V]) {
	visitedNodes := make(map[*vertex.Vertex[V]]bool)
	parents := make(map[*vertex.Vertex[V]]*vertex.Vertex[V])
	vertices := queue.NewQueue[*vertex.Vertex[V]]()
	vertices.Enqueue(start)

	for v, err := vertices.Dequeue(); err == nil; {
		v.Neighbours.ForEach(func(n *node.Node[*vertex.Vertex[V]]) error {
			if _, visited := visitedNodes[n.Data]; !visited {
				parents[n.Data] = v
				visitedNodes[n.Data] = true
				return vertices.Enqueue(n.Data)
			}
			return nil
		})
		fmt.Println(v)
	}
}

func (g *Graph[V]) search(data V) (*vertex.Vertex[V], error) {
	for it := g.Vertices.Head; it != nil; it = it.Next {
		if it.Data.GetVertexData() == data {
			return it.Data, nil
		}
	}
	return nil, fmt.Errorf("data %v not found in graph", data)
}

func (g *Graph[V]) contains(data V) bool {
	for it := g.Vertices.Head; it != nil; it = it.Next {
		if it.Data.GetVertexData() == data {
			return true
		}
	}
	return false
}

func (g *Graph[V]) Add(data, parent V) error {
	if g.Vertices.IsEmpty() {
		return g.Vertices.InsertLast(vertex.NewVertex(data))
	}

	if g.contains(data) {
		return fmt.Errorf("data %v already present in graph", data)
	}
	parentVertex, err := g.search(parent)
	if err != nil {
		return err
	}

	v := vertex.NewVertex(data)
	err = g.Vertices.InsertLast(v)
	if err != nil {
		return err
	}
	return parentVertex.Neighbours.InsertLast(v)
}
