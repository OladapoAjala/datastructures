package vertex

import "golang.org/x/exp/constraints"

type Vertex[V comparable, W constraints.Ordered] struct {
	State V
	Edges map[*Vertex[V, W]]W
}

func NewVertex[V comparable, W constraints.Ordered](data V) *Vertex[V, W] {
	return &Vertex[V, W]{
		State: data,
		Edges: make(map[*Vertex[V, W]]W),
	}
}

func (v *Vertex[V, W]) GetState() V {
	return v.State
}

func (v *Vertex[V, W]) AddEdge(edge *Vertex[V, W], w W) {
	v.Edges[edge] = w
}

func (v *Vertex[V, W]) RemoveEdge(edge *Vertex[V, W]) {
	delete(v.Edges, edge)
}

func (v *Vertex[V, W]) HasEdge(data V) bool {
	for e := range v.Edges {
		if e.State == data {
			return true
		}
	}
	return false
}
