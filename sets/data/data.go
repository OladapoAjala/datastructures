package data

import (
	"hash/fnv"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Data[K constraints.Ordered, V any] struct {
	Key   K
	Value V
	hash  uint32
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

func NewDataWithHash[K constraints.Ordered, V any](key K, val V) *Data[K, V] {
	hasher := fnv.New32()
	hasher.Write([]byte(ToString(key)))

	return &Data[K, V]{
		Key:   key,
		Value: val,
		hash:  hasher.Sum32(),
	}
}

func (d *Data[K, V]) Equal(val *Data[K, V]) bool {
	if d.hash != val.hash {
		return false
	}
	return d.GetKey() == val.GetKey()
}

func (d *Data[K, V]) GetKey() K {
	return d.Key
}

func (d *Data[K, V]) GetValue() V {
	return d.Value
}

func (d *Data[K, V]) GetHash() uint32 {
	return d.hash
}

func (d *Data[K, V]) Tomb() *Data[K, V] {
	return &Data[K, V]{
		Key:   *new(K),
		Value: *new(V),
		hash:  0,
	}
}

func ToString(val any) string {
	switch v := val.(type) {
	case int:
		return strconv.Itoa(val.(int))
	case int8:
		return strconv.Itoa(int(val.(int8)))
	case int16:
		return strconv.Itoa(int(val.(int16)))
	case int32:
		return strconv.Itoa(int(val.(int32)))
	case int64:
		return strconv.Itoa(int(val.(int64)))
	case uint:
		return strconv.FormatUint(uint64(val.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(val.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(val.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(val.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(val.(uint64)), 10)
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case string:
		return v
	default:
		return "unknown type"
	}
}
