package vertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Vertex(t *testing.T) {
	is := assert.New(t)

	a := NewVertex[string, int]("A")
	b := NewVertex[string, int]("B")
	c := NewVertex[string, int]("C")
	d := NewVertex[string, int]("D")
	e := NewVertex[string, int]("E")
	f := NewVertex[string, int]("F")
	h := NewVertex[string, int]("H")

	a.AddEdge(b, 1)
	a.AddEdge(c, 1)
	b.AddEdge(d, 1)
	b.AddEdge(e, 1)
	c.AddEdge(e, 1)
	c.AddEdge(f, 1)
	e.AddEdge(a, 1)
	e.AddEdge(h, 1)

	is.Equal(a.Edges[b], 1)
	is.Equal(a.Edges[c], 1)
	is.Equal(b.Edges[d], 1)
	is.Equal(b.Edges[e], 1)
	is.Equal(c.Edges[e], 1)
	is.Equal(c.Edges[f], 1)
	is.Equal(e.Edges[a], 1)
	is.Equal(e.Edges[h], 1)
}
