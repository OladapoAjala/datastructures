package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Unify(t *testing.T) {
	is := assert.New(t)
	uf, err := NewUnionFind("A", "B", "C", "D", "E")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		a, b string
	}

	tests := []struct {
		name      string
		args      args
		unionfind *UnionFind[string]
		want      func(int32, error)
	}{
		{
			name: "simple unification",
			args: args{
				a: "B",
				b: "D",
			},
			unionfind: uf,
			want: func(i int32, err error) {
				is.Nil(err)
				is.EqualValues(i, 1)
			},
		},
		{
			name: "second unification",
			args: args{
				a: "A",
				b: "C",
			},
			unionfind: uf,
			want: func(i int32, err error) {
				is.Nil(err)
				is.EqualValues(i, 0)
			},
		},
		{
			name: "third unification",
			args: args{
				a: "D",
				b: "E",
			},
			unionfind: uf,
			want: func(i int32, err error) {
				is.Nil(err)
				is.EqualValues(i, 1)
			},
		},
		{
			name: "fourth unification",
			args: args{
				a: "C",
				b: "E",
			},
			unionfind: uf,
			want: func(i int32, err error) {
				is.Nil(err)
				is.EqualValues(i, 1)
			},
		},
		{
			name: "unify elements in the same group",
			args: args{
				a: "A",
				b: "D",
			},
			unionfind: uf,
			want: func(i int32, err error) {
				is.Error(err, "A and B belong to the same group")
				is.EqualValues(i, -1)
			},
		},
		{
			name: "unify elements that doesn't exist",
			args: args{
				a: "F",
				b: "B",
			},
			unionfind: uf,
			want: func(i int32, err error) {
				is.Error(err, "F doesn't exist in unionfind")
				is.EqualValues(i, -1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, err := tt.unionfind.Unify(tt.args.a, tt.args.b)
			tt.want(root, err)
		})
	}
}

func Test_Conected(t *testing.T) {
	is := assert.New(t)
	uf, err := NewUnionFind("A", "B", "C", "D", "E")
	if err != nil {
		t.Error(err)
	}
	uf.Unify("C", "E")

	type args struct {
		a, b string
	}

	tests := []struct {
		name      string
		args      args
		unionfind *UnionFind[string]
		want      func(bool, error)
	}{
		{
			name: "check unconnected elements",
			args: args{
				a: "A",
				b: "D",
			},
			unionfind: uf,
			want: func(b bool, err error) {
				is.False(b)
				is.Nil(err)
			},
		},
		{
			name: "check connected elements",
			args: args{
				a: "C",
				b: "E",
			},
			unionfind: uf,
			want: func(b bool, err error) {
				is.True(b)
				is.Nil(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isConnected, err := tt.unionfind.Connected(tt.args.a, tt.args.b)
			tt.want(isConnected, err)
		})
	}
}
