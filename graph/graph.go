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

func (g *Graph[V, W]) HasCycle() bool {
	parent := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	for _, v := range g.Vertices {
		if _, visited := parent[v]; visited {
			continue
		}
		parent[v] = nil
		if g.hasCycle(v, parent) {
			return true
		}
	}
	return false
}

func (g *Graph[V, W]) hasCycle(v *vertex.Vertex[V, W], parent map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W]) bool {
	if v.InProcess {
		return true
	}
	if v.HasEmptyEdges() {
		return false
	}

	v.InProcess = true
	for e := range v.Edges {
		if e.InProcess {
			return true
		}
		if _, visited := parent[e]; visited {
			continue
		}
		parent[e] = v
		if g.hasCycle(e, parent) {
			return true
		}
	}
	v.InProcess = false
	return false
}

// TODO
func (g *Graph[V, W]) TopologicalSort() *vertex.Path[V, W] {
	parent := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	for _, v := range g.Vertices {
		if _, visited := parent[v]; visited {
			continue
		}
		parent[v] = nil
		g.depthFirstSearch(v, parent)
	}
	return nil
}

func (g *Graph[V, W]) DepthFirstSearchAll() {
	parent := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	for _, v := range g.Vertices {
		if _, visited := parent[v]; visited {
			continue
		}
		parent[v] = nil
		g.depthFirstSearch(v, parent)
	}
}

func (g *Graph[V, W]) DepthFirstSearch(start *vertex.Vertex[V, W]) {
	parent := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	parent[start] = nil
	g.depthFirstSearch(start, parent)
}

func (g *Graph[V, W]) depthFirstSearch(v *vertex.Vertex[V, W], parent map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W]) {
	if v.HasEmptyEdges() {
		return
	}

	for edge := range v.Edges {
		if _, visited := parent[edge]; visited {
			continue
		}
		parent[edge] = v
		g.depthFirstSearch(edge, parent)
	}
}

func (g *Graph[V, W]) BreadthFirstSearch(start *vertex.Vertex[V, W]) error {
	visitedNodes := make(map[*vertex.Vertex[V, W]]bool)
	parent := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	vertices := queue.NewQueue[*vertex.Vertex[V, W]]()

	vertices.Enqueue(start)
	visitedNodes[start] = true
	parent[start] = nil

	for v, err := vertices.Dequeue(); err == nil; v, err = vertices.Dequeue() {
		for u := range v.Edges {
			if _, visited := visitedNodes[u]; visited {
				continue
			}

			parent[u] = v
			visitedNodes[u] = true
			err := vertices.Enqueue(u)
			if err != nil {
				return err
			}
		}
		fmt.Println(v.GetState())
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
		if parent == *new(V) {
			g.Vertices = append(g.Vertices, vertex.NewVertex[V, W](state))
			return nil
		} else {
			return err
		}
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
