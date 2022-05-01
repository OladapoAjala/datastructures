package unionfind

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type UnionFind[T constraints.Ordered] struct {
	hash          map[T]int32
	indices       []int32
	size          int32
	numComponents int32
	compSize      []int32
}

type UnionFinder[T constraints.Ordered] interface {
	Unify(T, T) (int32, error)
	Find(T) (int32, error)
	Size() int32
	Connected(T, T) (bool, error)
}

var _ UnionFinder[int] = new(UnionFind[int])

func NewUnionFind[T constraints.Ordered](elems ...T) (*UnionFind[T], error) {
	if len(elems) == 0 {
		return nil, fmt.Errorf("can't create unionfind with empty elements")
	}

	unionFind := &UnionFind[T]{
		hash:    make(map[T]int32),
		indices: make([]int32, 0),
	}

	for i, elem := range elems {
		unionFind.hash[elem] = int32(i)
		unionFind.indices = append(unionFind.indices, int32(i))
		unionFind.size++
		unionFind.numComponents++
		unionFind.compSize[i] = 1
	}

	return unionFind, nil
}

func (uf *UnionFind[T]) Unify(a, b T) (int32, error) {
	rootA, err := uf.Find(a)
	if err != nil {
		return -1, err
	}

	rootB, err := uf.Find(b)
	if err != nil {
		return -1, err
	}

	if rootA == rootB {
		return -1, fmt.Errorf("%v and %v belong to the same group", a, b)
	}

	if uf.compSize[rootA] >= uf.compSize[rootB] {
		uf.indices[rootB] = rootA
		uf.compSize[rootA] += uf.compSize[rootB]
		uf.numComponents--
		return rootA, nil
	}

	uf.indices[rootA] = rootB
	uf.compSize[rootB] += uf.compSize[rootA]
	uf.numComponents--
	return rootB, nil
}

func (uf *UnionFind[T]) Find(a T) (int32, error) {
	root, ok := uf.hash[a]
	if !ok {
		return -1, fmt.Errorf("%v doesn't exist in unionfind", a)
	}

	for root != uf.indices[root] {
		root = uf.indices[root]
	}

	p := uf.hash[a]
	for p != root {
		p = uf.indices[p]
		uf.indices[p] = root
	}

	return root, nil
}

func (uf UnionFind[T]) Size() int32 {
	return uf.size
}

func (uf *UnionFind[T]) Connected(a, b T) (bool, error) {
	rootA, err := uf.Find(a)
	if err != nil {
		return false, err
	}

	rootB, err := uf.Find(b)
	if err != nil {
		return false, err
	}

	return (rootA == rootB), nil
}
