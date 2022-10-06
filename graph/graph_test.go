package graph

import "testing"

func TestGraph_AddVertex(t *testing.T) {
	type fields struct {
		vertices []*Vertex
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
			}
			g.AddVertex(tt.args.key)
		})
	}
}
