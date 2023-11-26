package vertex

import "fmt"

type Vertex[V any] struct {
	VertexData V
	Neighbours []*Vertex[V]
}

func NewVertex[V any](data V) *Vertex[V] {
	return &Vertex[V]{
		VertexData: data,
		Neighbours: make([]*Vertex[V], 0),
	}
}

func (v *Vertex[V]) GetVertexData() V {
	return v.VertexData
}

func (v *Vertex[V]) AddNeighbour(neighbour *Vertex[V]) {
	v.Neighbours = append(v.Neighbours, neighbour)
}

func (v *Vertex[V]) RemoveNeighbour(neighbour *Vertex[V]) error {
	var index int32 = 0
	found := false
	for i, n := range v.Neighbours {
		if n == neighbour {
			index = int32(i)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("vertex %v not found in tree", neighbour)
	}

	neighbours := v.Neighbours[0:index]
	neighbours = append(neighbours, v.Neighbours[index+1:]...)
	v.Neighbours = neighbours
	return nil
}
