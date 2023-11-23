package vertex

import "github.com/OladapoAjala/datastructures/sequences/linkedlist"

type Vertex[V any] struct {
	Data       V
	Neighbours *linkedlist.LinkedList[*Vertex[V]]
}
