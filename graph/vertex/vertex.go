package vertex

import (
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
)

type Vertex[V any] struct {
	VertexData V
	Neighbours *linkedlist.LinkedList[*Vertex[V]]
}

func NewVertex[V any](data V) *Vertex[V] {
	return &Vertex[V]{
		VertexData: data,
		Neighbours: linkedlist.NewList[*Vertex[V]](),
	}
}

func (v *Vertex[V]) GetVertexData() V {
	return v.VertexData
}

func (v *Vertex[V]) AddNeighbour(neighbour *Vertex[V]) error {
	return v.Neighbours.InsertLast(neighbour)
}

func (v *Vertex[V]) RemoveNeighbour(neighbour *Vertex[V]) error {
	var index int32 = 0
	for it := v.Neighbours.Head; it != nil; it = it.Next {
		if it.Data == neighbour {
			break
		}
		index++
	}
	return v.Neighbours.Delete(index)
}
