package graph

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/queues/queue"
	"github.com/OladapoAjala/datastructures/sequences/node"
)

type Graph[V comparable] struct {
	Vertecies []*vertex.Vertex[V]
}

func (g *Graph[V]) BreadthFirstSearch(val *vertex.Vertex[V]) {
	visitedNodes := make(map[*vertex.Vertex[V]]bool)
	vertices := queue.NewQueue[*vertex.Vertex[V]]()
	vertices.Enqueue(val)

	for v, err := vertices.Dequeue(); err == nil; {
		v.Neighbours.ForEach(func(n *node.Node[*vertex.Vertex[V]]) {
			if _, visited := visitedNodes[n.Data]; !visited {
				visitedNodes[n.Data] = true
				if err := vertices.Enqueue(n.Data); err != nil {
					return
				}
			}
			vertices.Enqueue(n.Data)
		})
		fmt.Println(v)
	}
}
