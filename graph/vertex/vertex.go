package vertex

import "golang.org/x/exp/constraints"

type Vertex[V comparable, W constraints.Ordered] struct {
	State V
	Edges map[*Vertex[V, W]]W
}

func NewVertex[V comparable, W constraints.Ordered](state V) *Vertex[V, W] {
	return &Vertex[V, W]{
		State: state,
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

func (v *Vertex[V, W]) HasEdge(state V) bool {
	for edge := range v.Edges {
		if edge.GetState() == state {
			return true
		}
	}
	return false
}

func (v *Vertex[V, W]) HasEmptyEdges() bool {
	return len(v.Edges) == 0
}

func (v *Vertex[V, W]) GetEdge(state V) *Vertex[V, W] {
	for edge := range v.Edges {
		if edge.GetState() == state {
			return edge
		}
	}
	return nil
}
