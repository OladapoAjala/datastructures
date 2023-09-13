package data

import "golang.org/x/exp/constraints"

type Data[K constraints.Ordered, V any] struct {
	Key   K
	Value V
}

type IData[K constraints.Ordered, V any] interface {
	GetKey() K
	GetValue() V
}

var _ IData[int, string] = new(Data[int, string])

func NewData[K constraints.Ordered, V any](key K, val V) *Data[K, V] {
	return &Data[K, V]{
		Key:   key,
		Value: val,
	}
}

func (d *Data[K, V]) GetKey() K {
	return d.Key
}

func (d *Data[K, V]) GetValue() V {
	return d.Value
}
