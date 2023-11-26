package graph

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/queues/queue"
)

type Graph[V comparable] struct {
	Vertices []*vertex.Vertex[V]
}

func NewGraph[V comparable]() *Graph[V] {
	return &Graph[V]{
		Vertices: make([]*vertex.Vertex[V], 0),
	}
}

func (g *Graph[V]) DepthFirstSearch(start *vertex.Vertex[V]) {
	parent := make(map[*vertex.Vertex[V]]*vertex.Vertex[V])
	parent[start] = nil
	g.depthFirstSearch(start, parent)

	fmt.Println("---back tracking---")
	for p := parent[g.Vertices[4]]; p != nil; {
		fmt.Printf("%v -> ", p.VertexData)
		p = parent[p]
	}
}

func (g *Graph[V]) depthFirstSearch(v *vertex.Vertex[V], parent map[*vertex.Vertex[V]]*vertex.Vertex[V]) {
	if v.Neighbours == nil {
		return
	}
	fmt.Printf("%v\n", v.VertexData)
	for _, u := range v.Neighbours {
		if _, visited := parent[u]; visited {
			continue
		}
		parent[u] = v
		g.depthFirstSearch(u, parent)
	}
}

func (g *Graph[V]) BreadthFirstSearch(start *vertex.Vertex[V]) error {
	visitedNodes := make(map[*vertex.Vertex[V]]bool)
	parents := make(map[*vertex.Vertex[V]]*vertex.Vertex[V])
	vertices := queue.NewQueue[*vertex.Vertex[V]]()
	vertices.Enqueue(start)
	parents[start] = nil

	for v, err := vertices.Dequeue(); err == nil; v, err = vertices.Dequeue() {
		for _, u := range v.Neighbours {
			if _, visited := visitedNodes[u]; visited {
				continue
			}

			parents[u] = v
			visitedNodes[u] = true
			err := vertices.Enqueue(u)
			if err != nil {
				return err
			}
		}
		fmt.Println(v)
	}
	return nil
}

func (g *Graph[V]) search(data V) (int, error) {
	for i, d := range g.Vertices {
		if d.VertexData == data {
			return i, nil
		}
	}
	return -1, fmt.Errorf("data %v not found in graph", data)
}

func (g *Graph[V]) contains(data V) bool {
	for _, d := range g.Vertices {
		if d.VertexData == data {
			return true
		}
	}
	return false
}

func (g *Graph[V]) Add(data, parent V) error {
	if len(g.Vertices) == 0 {
		g.Vertices = append(g.Vertices, vertex.NewVertex(data))
		return nil
	}

	if g.contains(data) {
		return fmt.Errorf("data %v already present in graph", data)
	}
	parentIndex, err := g.search(parent)
	if err != nil {
		return err
	}
	parentVertex := g.Vertices[parentIndex]
	v := vertex.NewVertex(data)
	g.Vertices = append(g.Vertices, v)
	parentVertex.Neighbours = append(parentVertex.Neighbours, v)
	return nil
}
