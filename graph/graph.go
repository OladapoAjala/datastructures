package graph

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/queues/queue"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sequences/node"
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

func (g *Graph[V, W]) TopologicalSort() (*linkedlist.LinkedList[*vertex.Vertex[V, W]], error) {
	if g.HasCycle() {
		return nil, fmt.Errorf("cannot topologically sort cyclic graph")
	}

	visited := make(map[*vertex.Vertex[V, W]]bool)
	sorted := linkedlist.NewList[*vertex.Vertex[V, W]]()
	for _, v := range g.Vertices {
		if _, visited := visited[v]; visited {
			continue
		}
		g.topologicalSort(v, visited, sorted)
	}
	return sorted, nil
}

func (g *Graph[V, W]) topologicalSort(v *vertex.Vertex[V, W],
	visited map[*vertex.Vertex[V, W]]bool,
	path *linkedlist.LinkedList[*vertex.Vertex[V, W]]) error {
	visited[v] = true
	for e := range v.Edges {
		if visited[e] {
			continue
		}

		err := g.topologicalSort(e, visited, path)
		if err != nil {
			return err
		}
	}
	return path.InsertFirst(v)
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

func (g *Graph[V, W]) ShortestPathTopologicalSort(start, stop V) (*linkedlist.LinkedList[*vertex.Vertex[V, W]], error) {
	startVertex, err := g.Search(start)
	if err != nil {
		return nil, err
	}
	stopVertex, err := g.Search(stop)
	if err != nil {
		return nil, err
	}
	return g.shortestPathTopologicalSort(startVertex, stopVertex)
}

func (g *Graph[V, W]) shortestPathTopologicalSort(start, stop *vertex.Vertex[V, W]) (*linkedlist.LinkedList[*vertex.Vertex[V, W]], error) {
	topologicalOrder, err := g.TopologicalSort()
	if err != nil {
		return nil, err
	}

	pi := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	delta := make(map[*vertex.Vertex[V, W]]W)
	first := topologicalOrder.Head.Data
	pi[first] = nil
	delta[first] = *new(W)

	_ = topologicalOrder.ForEach(func(n *node.Node[*vertex.Vertex[V, W]]) error {
		for e, w := range n.Data.Edges {
			if _, visited := pi[e]; visited {
				currWeight := delta[e]
				calcWeight := delta[n.Data] + w
				if currWeight <= calcWeight {
					continue
				}
			}

			pi[e] = n.Data
			delta[e] = delta[n.Data] + w
		}
		return nil
	})

	return path(pi, start, stop)
}

func path[V comparable, W constraints.Ordered](pi map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W], start, end *vertex.Vertex[V, W]) (*linkedlist.LinkedList[*vertex.Vertex[V, W]], error) {
	if end == start {
		list := linkedlist.NewList[*vertex.Vertex[V, W]]()
		return list, list.InsertLast(end)
	}

	list, err := path(pi, start, pi[end])
	if err != nil {
		return nil, err
	}
	return list, list.InsertLast(end)
}

func (g *Graph[V, W]) ShortestPath(start, stop V) (*linkedlist.LinkedList[*vertex.Vertex[V, W]], error) {
	startVertex, err := g.Search(start)
	if err != nil {
		return nil, err
	}
	stopVertex, err := g.Search(stop)
	if err != nil {
		return nil, err
	}
	return g.shortestPath(startVertex, stopVertex)
}

func (g *Graph[V, W]) shortestPath(start, stop *vertex.Vertex[V, W]) (*linkedlist.LinkedList[*vertex.Vertex[V, W]], error) {
	delta := make(map[V]W)
	pi := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	delta[start.GetState()] = *new(W)
	pi[start] = nil

	vertices := queue.NewQueue[*vertex.Vertex[V, W]]()
	vertices.Enqueue(start)
	for v, err := vertices.Dequeue(); err == nil; v, err = vertices.Dequeue() {
		for u, w := range v.Edges {
			if _, visited := pi[u]; visited {
				currDistance := delta[u.GetState()]
				calcDistance := delta[v.GetState()] + w
				if calcDistance < currDistance {
					delta[u.GetState()] = calcDistance
					pi[u] = v
				}
				continue
			}

			pi[u] = v
			delta[u.GetState()] = delta[v.GetState()] + w
			err := vertices.Enqueue(u)
			if err != nil {
				return nil, err
			}
		}
	}

	path := linkedlist.NewList[*vertex.Vertex[V, W]]()
	end := stop
	for end != nil {
		err := path.InsertFirst(end)
		if err != nil {
			return nil, err
		}
		end = pi[end]
	}

	isValidPath := false
	path.ForEach(func(n *node.Node[*vertex.Vertex[V, W]]) error {
		if n.Data.GetState() == start.GetState() {
			isValidPath = true
		}
		return nil
	})
	if isValidPath {
		return path, nil
	}
	return nil, fmt.Errorf("no path from %v -> %v", start.GetState(), stop.GetState())
}

func (g *Graph[V, W]) Search(data V) (*vertex.Vertex[V, W], error) {
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
		g.Vertices = append(g.Vertices, vertex.NewVertex[V, W](parent))
	}

	parentVertex, err := g.Search(parent)
	if err != nil {
		if err.Error() == fmt.Sprintf("data %v not found in graph", parent) {
			parentVertex = vertex.NewVertex[V, W](parent)
			g.Vertices = append(g.Vertices, parentVertex)
		} else {
			return err
		}
	}
	if parentVertex.HasEdge(state) {
		return fmt.Errorf("edge %v -> %v is already present in graph", parent, state)
	}

	stateVertex, _ := g.Search(state)
	if stateVertex != nil {
		parentVertex.AddEdge(stateVertex, weight)
		return nil
	}

	v := vertex.NewVertex[V, W](state)
	parentVertex.AddEdge(v, weight)
	g.Vertices = append(g.Vertices, v)
	return nil
}
