package graph

import (
	"errors"
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/sequences/node"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Add(t *testing.T) {
	is := assert.New(t)
	graph := NewGraph[string, int]()

	tests := []struct {
		name   string
		data   string
		parent string
		want   func(error)
	}{
		{
			name:   "Add to an empty graph",
			data:   "A",
			parent: "",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(len(graph.Vertices), 1)
				a := graph.Vertices[0]
				is.Equal(a.GetState(), "A")
			},
		},
		{
			name:   "Add to an existing graph",
			data:   "B",
			parent: "A",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(len(graph.Vertices), 2)
				b := graph.Vertices[1]
				is.Equal(b.GetState(), "B")
				is.EqualValues(len(b.Edges), 0)

				a := graph.Vertices[0]
				is.EqualValues(len(a.Edges), 1)
				is.True(a.HasEdge(b.GetState()))
			},
		},
		{
			name:   "Add a back edge",
			data:   "A",
			parent: "B",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(len(graph.Vertices), 2)

				b := graph.Vertices[1]
				is.Equal(b.GetState(), "B")
				is.EqualValues(len(b.Edges), 1)

				a := graph.Vertices[0]
				is.EqualValues(len(a.Edges), 1)
				is.True(a.HasEdge(b.GetState()))
				is.True(b.HasEdge(a.GetState()))
			},
		},
		{
			name:   "Add the same edge",
			data:   "A",
			parent: "B",
			want: func(err error) {
				is.Error(err, fmt.Errorf("edge B -> A already present in graph"))
				is.EqualValues(len(graph.Vertices), 2)
			},
		},
		{
			name:   "Add with non-existent parent",
			data:   "C",
			parent: "Z",
			want: func(err error) {
				is.Error(err, fmt.Errorf("data Z not found in graph"))
				is.EqualValues(len(graph.Vertices), 2)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := graph.Add(0, tt.parent, tt.data)
			tt.want(err)
		})
	}
}

func Test_DepthFirstSearch(t *testing.T) {
	graph := NewGraph[string, int]()
	graph.Add(1, "", "A")
	graph.Add(1, "A", "B")
	graph.Add(1, "A", "C")
	graph.Add(1, "B", "D")
	graph.Add(1, "B", "E")
	graph.Add(1, "C", "F")
	graph.Add(1, "C", "E")
	graph.Add(1, "E", "A")
	// graph.DepthFirstSearch(graph.Vertices[0])
	graph.DepthFirstSearchAll()
}

func Test_BreadthFirstSearch(t *testing.T) {
	graph := NewGraph[string, int]()
	graph.Add(1, "", "A")
	graph.Add(1, "A", "B")
	graph.Add(1, "A", "C")
	graph.Add(1, "B", "D")
	graph.Add(1, "B", "E")
	graph.Add(1, "C", "F")
	graph.Add(1, "C", "E")
	// graph.Add(1, "E", "A")
	graph.BreadthFirstSearch(graph.Vertices[0])

	s, _ := graph.TopologicalSort()
	s.ForEach(func(n *node.Node[*vertex.Vertex[string, int]]) error {
		fmt.Printf("%s -> ", n.Data.GetState())
		return nil
	})
}

func Test_ShortestPath(t *testing.T) {
	graph := NewGraph[string, int]()
	graph.Add(1, "", "A")
	graph.Add(2, "A", "B")
	graph.Add(1, "A", "C")
	graph.Add(1, "B", "D")
	graph.Add(2, "B", "E")
	graph.Add(1, "C", "F")
	graph.Add(1, "C", "B")

	testCases := []struct {
		Start         string
		Stop          string
		ExpectedPath  []string
		ExpectedError error
	}{
		{"A", "B", []string{"A", "B"}, nil},
		{"A", "F", []string{"A", "C", "F"}, nil},
		{"E", "A", nil, errors.New("no path from E -> A")},
		{"X", "Y", nil, errors.New("data X not found in graph")},
		{"A", "X", nil, errors.New("data X not found in graph")},
		{"A", "A", []string{"A"}, nil},
		{"A", "C", []string{"A", "C"}, nil},
		{"D", "D", []string{"D"}, nil},
		{"F", "B", nil, errors.New("no path from F -> B")},
	}

	for _, tc := range testCases {
		path, err := graph.ShortestPath(tc.Start, tc.Stop)

		if tc.ExpectedError != nil {
			require.Error(t, err)
			require.EqualError(t, err, tc.ExpectedError.Error())
		} else {
			require.NoError(t, err)

			var actualPath []string
			path.ForEach(func(n *node.Node[*vertex.Vertex[string, int]]) error {
				actualPath = append(actualPath, n.Data.GetState())
				return nil
			})
			require.Equal(t, tc.ExpectedPath, actualPath)
		}
	}
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Graph[string, int])
		hasCycle bool
	}{
		{
			name:     "Empty Graph",
			setup:    func(g *Graph[string, int]) {},
			hasCycle: false,
		},
		{
			name: "Single Vertex No Cycle",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
			},
			hasCycle: false,
		},
		{
			name: "Single Vertex With Self-Loop",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
				g.Add(0, "A", "A")
			},
			hasCycle: true,
		},
		{
			name: "Acyclic Graph",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				g.Add(0, "B", "C")
			},
			hasCycle: false,
		},
		{
			name: "Simple Cycle",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				g.Add(0, "B", "C")
				g.Add(0, "C", "A")
			},
			hasCycle: true,
		},
		{
			name: "Complex Graph With Cycle",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				g.Add(0, "B", "C")
				g.Add(0, "C", "D")
				g.Add(0, "D", "B")
			},
			hasCycle: true,
		},
		{
			name: "Complex Graph Without Cycle",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				g.Add(0, "B", "C")
				g.Add(0, "C", "D")
				g.Add(0, "D", "E")
			},
			hasCycle: false,
		},
		{
			name: "Disconnected Graph With Cycle",
			setup: func(g *Graph[string, int]) {
				// Component 1 (Acyclic)
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				// Component 2 (Cyclic)
				g.Add(0, "", "X")
				g.Add(0, "X", "Y")
				g.Add(0, "Y", "X")
			},
			hasCycle: true,
		},
		{
			name: "Disconnected Graph Without Cycle",
			setup: func(g *Graph[string, int]) {
				// Component 1
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				// Component 2
				g.Add(0, "", "X")
				g.Add(0, "X", "Y")
			},
			hasCycle: false,
		},
		{
			name: "Back Edge Test",
			setup: func(g *Graph[string, int]) {
				g.Add(0, "", "A")
				g.Add(0, "A", "B")
				g.Add(0, "B", "C")
				g.Add(0, "C", "A")
			},
			hasCycle: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph[string, int]()
			tt.setup(graph)
			assert.Equal(t, tt.hasCycle, graph.HasCycle())
		})
	}
}
