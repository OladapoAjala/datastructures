package vertex

import "golang.org/x/exp/constraints"

type destination[V any, W constraints.Ordered] *Vertex[V, W]

type Vertex[V any, W constraints.Ordered] struct {
	VertexData V
	Edges      map[destination[V, W]]W
}

func NewVertex[V any, W constraints.Ordered](data V) *Vertex[V, W] {
	return &Vertex[V, W]{
		VertexData: data,
		Edges:      make(map[destination[V, W]]W),
	}
}

func (v *Vertex[V, W]) GetVertexData() V {
	return v.VertexData
}

func (v *Vertex[V, W]) AddEdge(edge *Vertex[V, W], w W) {
	v.Edges[edge] = w
}

func (v *Vertex[V, W]) RemoveEdge(edge *Vertex[V, W]) {
	delete(v.Edges, edge)
}
