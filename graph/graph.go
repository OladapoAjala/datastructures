package graph

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/queues/queue"
	"golang.org/x/exp/constraints"
)

type Graph[V comparable, W constraints.Ordered] struct {
	Vertices []*vertex.Vertex[V, W]
}

func NewGraph[V comparable, W constraints.Ordered]() *Graph[V, W] {
	return &Graph[V, W]{
		Vertices: make([]*vertex.Vertex[V, W], 0),
	}
}

var count = 0

func (g *Graph[V, W]) DepthFirstSearch(start *vertex.Vertex[V, W]) {
	parent := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	parent[start] = nil
	g.depthFirstSearch(start, parent)

	fmt.Println("---back tracking---")
	fmt.Printf("%v", g.Vertices[len(g.Vertices)-1].State)
	for p := parent[g.Vertices[len(g.Vertices)-1]]; p != nil; {
		fmt.Printf(" -> %v", p.State)
		p = parent[p]
	}

	fmt.Printf("\nCount %d\n", count)
}

func (g *Graph[V, W]) depthFirstSearch(v *vertex.Vertex[V, W], parent map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W]) {
	if len(v.Edges) == 0 {
		return
	}

	for d := range v.Edges {
		count++
		if _, visited := parent[d]; visited {
			continue
		}
		fmt.Printf("%v ", d.State)
		parent[d] = v
		g.depthFirstSearch(d, parent)
	}
}

func (g *Graph[V, W]) BreadthFirstSearch(start *vertex.Vertex[V, W]) error {
	visitedNodes := make(map[*vertex.Vertex[V, W]]bool)
	parents := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	vertices := queue.NewQueue[*vertex.Vertex[V, W]]()
	vertices.Enqueue(start)
	parents[start] = nil

	for v, err := vertices.Dequeue(); err == nil; v, err = vertices.Dequeue() {
		for u := range v.Edges {
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

func (g *Graph[V, W]) ShortestPath(start, stop *vertex.Vertex[V, W]) error {
	// delta := make(map[*vertex.Vertex[V, W]]int)
	// pi := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	// delta[start] = 0
	// pi[start] = nil

	// vertices := queue.NewQueue[*vertex.Vertex[V, W]]()
	// vertices.Enqueue(start)
	// for v, err := vertices.Dequeue(); err == nil; v, err = vertices.Dequeue() {
	// 	for _, u := range v.Neighbours {
	// 		if _, visited := visitedNodes[u]; visited {
	// 			continue
	// 		}

	// 		parents[u] = v
	// 		visitedNodes[u] = true
	// 		err := vertices.Enqueue(u)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// 	fmt.Println(v)
	// }

	return nil
}

func (g *Graph[V, W]) search(data V) (*vertex.Vertex[V, W], error) {
	for i, d := range g.Vertices {
		if d.State == data {
			return g.Vertices[i], nil
		}
	}
	return nil, fmt.Errorf("data %v not found in graph", data)
}

func (g *Graph[V, W]) contains(data V) bool {
	for _, d := range g.Vertices {
		if d.State == data {
			return true
		}
	}
	return false
}

func (g *Graph[V, W]) Add(weight W, parent, state V) error {
	if len(g.Vertices) == 0 {
		g.Vertices = append(g.Vertices, vertex.NewVertex[V, W](state))
		return nil
	}

	parentVertex, err := g.search(parent)
	if err != nil {
		return err
	}
	if parentVertex.HasEdge(state) {
		return fmt.Errorf("edge %v -> %v is already present in graph", parent, state)
	}

	stateVertex, err := g.search(state)
	if stateVertex != nil {
		parentVertex.AddEdge(stateVertex, weight)
		return nil
	}

	v := vertex.NewVertex[V, W](state)
	parentVertex.AddEdge(v, weight)
		g.Vertices = append(g.Vertices, v)
	return nil
}
