package data

import (
	"golang.org/x/exp/constraints"
)

type Data[K constraints.Ordered, V comparable] struct {
	Key    K
	Value  V
	Size   int32
	Height int32
	Parent *Data[K, V]
	Left   *Data[K, V]
	Right  *Data[K, V]
}

type IData[K constraints.Ordered, V comparable] interface {
	GetKey() K
	GetValue() V
	GetParent() *Data[K, V]
	GetLeft() *Data[K, V]
	GetRight() *Data[K, V]
	GetHeight() int32
	GetSize() int32
	IsLeaf() bool
	Skew() int32
}

var _ IData[int, string] = new(Data[int, string])

func NewData[K constraints.Ordered, V comparable](key K, val V) *Data[K, V] {
	return &Data[K, V]{
		Key:   key,
		Value: val,
		Size:  1,
	}
}

func (d *Data[K, V]) GetKey() K {
	return d.Key
}

func (d *Data[K, V]) GetValue() V {
	return d.Value
}

func (n *Data[K, V]) GetParent() *Data[K, V] {
	return n.Parent
}

func (n *Data[K, V]) GetLeft() *Data[K, V] {
	return n.Left
}

func (n *Data[K, V]) GetRight() *Data[K, V] {
	return n.Right
}

func (n *Data[K, V]) GetHeight() int32 {
	return n.Height
}

func (n *Data[K, V]) GetSize() int32 {
	return n.Size
}

func (n *Data[K, V]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Data[K, V]) Skew() int32 {
	var hr, hl int32 = -1, -1
	if n.Right != nil {
		hr = n.Right.Height
	}
	if n.Left != nil {
		hl = n.Left.Height
	}
	return hr - hl
}

func (n *Data[K, V]) IsEqual(v *Data[K, V]) bool {
	return n.Value == v.Value
}

func (n *Data[K, V]) IsEmpty() bool {
	return n.Value == *new(V) || n.Key == *new(K)
}
