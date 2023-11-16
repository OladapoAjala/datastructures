package data

import (
	"golang.org/x/exp/constraints"
)

type Data[K constraints.Ordered, V comparable] struct {
	Key   K
	Value V
	Index int32
}

type IData[K constraints.Ordered, V comparable] interface {
	GetKey() K
	GetValue() V
	GetIndex() int32
}

var _ IData[int, string] = new(Data[int, string])

func NewData[K constraints.Ordered, V comparable](key K, val V, index int32) *Data[K, V] {
	return &Data[K, V]{
		Key:   key,
		Value: val,
		Index: index,
	}
}

func (d *Data[K, V]) GetKey() K {
	return d.Key
}

func (d *Data[K, V]) GetValue() V {
	return d.Value
}

func (n *Data[K, V]) GetIndex() int32 {
	return n.Index
}

func (n *Data[K, V]) GetParentIndex() int32 {
	return (n.Index - 1) / 2
}
func (n *Data[K, V]) GetLeftIndex() int32 {
	return (2*n.Index + 1)
}
func (n *Data[K, V]) GetRightIndex() int32 {
	return 2 * (n.Index + 1)
}
