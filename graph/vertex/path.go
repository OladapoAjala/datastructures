package vertex

import (
	"golang.org/x/exp/constraints"
)

type Path[V comparable, W constraints.Ordered] []*Vertex[V, W]
