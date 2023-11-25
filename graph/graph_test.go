package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	is := assert.New(t)
	graph := NewGraph[int]()

	tests := []struct {
		name   string
		data   int
		parent int
		want   func(error)
	}{
		{
			name:   "Add to an empty graph",
			data:   1,
			parent: 0,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(graph.Vertices.GetSize(), 1)
				data, err := graph.Vertices.GetData(0)
				is.Nil(err)
				is.Equal(data.GetVertexData(), 1)
			},
		},
		{
			name:   "Add to an existing graph",
			data:   2,
			parent: 1,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(graph.Vertices.GetSize(), 2)
				data, err := graph.Vertices.GetData(1)
				is.Nil(err)
				is.Equal(data.GetVertexData(), 2)
				is.EqualValues(data.Neighbours.GetSize(), 0)

				parentVertex, err := graph.Vertices.GetData(0)
				is.Nil(err)
				is.EqualValues(parentVertex.Neighbours.GetSize(), 1)
				data, err = parentVertex.Neighbours.GetData(0)
				is.Nil(err)
				is.Equal(data.GetVertexData(), 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := graph.Add(tt.data, tt.parent)
			tt.want(err)
		})
	}
}
