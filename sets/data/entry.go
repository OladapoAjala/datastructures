package data

import (
	"hash/fnv"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Entry[K constraints.Ordered] struct {
	*Data[K, any]
	hash uint32
}

func NewEntry[K constraints.Ordered](key K, value any) *Entry[K] {
	hasher := fnv.New32()
	hasher.Write([]byte(ToString(key)))

	return &Entry[K]{
		Data: NewData[K, any](key, value),
		hash: hasher.Sum32(),
	}
}

func (e *Entry[K]) Equal(val *Entry[K]) bool {
	if e.hash != val.hash {
		return false
	}
	return e.GetKey() == val.GetKey()
}

func (e *Entry[K]) GetHash() uint32 {
	return e.hash
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
