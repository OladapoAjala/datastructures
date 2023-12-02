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
	g := NewVertex[string, int]("G")

	a.AddEdge(b, 1)
	a.AddEdge(c, 1)
	b.AddEdge(d, 1)
	b.AddEdge(e, 1)
	c.AddEdge(e, 1)
	c.AddEdge(f, 1)
	e.AddEdge(a, 1)
	e.AddEdge(g, 1)

	is.True(a.HasEdge(b.GetState()))
	is.True(a.HasEdge(c.GetState()))
	is.True(b.HasEdge(d.GetState()))
	is.True(b.HasEdge(e.GetState()))
	is.True(c.HasEdge(e.GetState()))
	is.True(c.HasEdge(f.GetState()))
	is.True(e.HasEdge(a.GetState()))
	is.True(e.HasEdge(g.GetState()))

	is.False(a.HasEdge(d.GetState()))
	is.False(a.HasEdge(e.GetState()))
	is.False(a.HasEdge(f.GetState()))
	is.False(a.HasEdge(g.GetState()))

	is.False(b.HasEdge(a.GetState()))
	is.False(b.HasEdge(c.GetState()))
	is.False(b.HasEdge(f.GetState()))
	is.False(b.HasEdge(g.GetState()))

	is.False(c.HasEdge(a.GetState()))
	is.False(c.HasEdge(b.GetState()))
	is.False(c.HasEdge(d.GetState()))
	is.False(c.HasEdge(g.GetState()))

	is.False(d.HasEdge(a.GetState()))
	is.False(d.HasEdge(b.GetState()))
	is.False(d.HasEdge(c.GetState()))
	is.False(d.HasEdge(e.GetState()))
	is.False(d.HasEdge(f.GetState()))
	is.False(d.HasEdge(g.GetState()))

	is.False(e.HasEdge(b.GetState()))
	is.False(e.HasEdge(c.GetState()))
	is.False(e.HasEdge(d.GetState()))
	is.False(e.HasEdge(f.GetState()))

	is.False(f.HasEdge(a.GetState()))
	is.False(f.HasEdge(b.GetState()))
	is.False(f.HasEdge(c.GetState()))
	is.False(f.HasEdge(d.GetState()))
	is.False(f.HasEdge(e.GetState()))
	is.False(f.HasEdge(g.GetState()))

	is.False(g.HasEdge(a.GetState()))
	is.False(g.HasEdge(b.GetState()))
	is.False(g.HasEdge(c.GetState()))
	is.False(g.HasEdge(d.GetState()))
	is.False(g.HasEdge(e.GetState()))
	is.False(g.HasEdge(f.GetState()))
}
