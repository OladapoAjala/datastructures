package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	is := assert.New(t)
	graph := NewGraph[int, int]()

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
				is.EqualValues(len(graph.Vertices), 1)
				data := graph.Vertices[0]
				is.Equal(data.GetVertexData(), 1)
			},
		},
		{
			name:   "Add to an existing graph",
			data:   2,
			parent: 1,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(len(graph.Vertices), 2)
				data := graph.Vertices[1]
				is.Nil(err)
				is.Equal(data.GetVertexData(), 2)
				is.EqualValues(len(data.Edges), 0)

				parentVertex := graph.Vertices[0]
				is.Nil(err)
				is.EqualValues(len(parentVertex.Edges), 1)
				_, present := parentVertex.Edges[data]
				is.True(present)
			},
		},
		{
			name:   "Add with existing data",
			data:   1,
			parent: 2,
			want: func(err error) {
				is.Error(err, fmt.Errorf("data 1 already present in graph"))
				is.EqualValues(len(graph.Vertices), 2)
			},
		},
		{
			name:   "Add with non-existent parent",
			data:   3,
			parent: 999,
			want: func(err error) {
				is.Error(err, fmt.Errorf("data 999 not found in graph"))
				is.EqualValues(len(graph.Vertices), 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := graph.Add(0, tt.data, tt.parent)
			tt.want(err)
		})
	}
}
