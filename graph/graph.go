package graph

import "fmt"

type Vertex struct {
	key       int
	adjacents []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		panic("Key already exists in graph")
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		panic(fmt.Sprintf("Invalid edge (%v --> %v)", from, to))
	}

	if contains(fromVertex.adjacents, to) {
		panic(fmt.Sprintf("Existing edge (%v --> %v)", from, to))
	}
	fromVertex.adjacents = append(fromVertex.adjacents, toVertex)
}

func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

func contains(vertices []*Vertex, k int) bool {
	for _, v := range vertices {
		if v.key == k {
			return true
		}
	}

	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex: %v", v)
		for _, a := range v.adjacents {
			fmt.Printf(" %v", a.key)
		}
		fmt.Println()
	}
}
