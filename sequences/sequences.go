package sequences

type Sequencer[T any] interface {
	GetData(int32) (T, error)
	Contains(T) bool
	Insert(int32, T) error // put new value at index
	Set(int32, T) error    // replace value at index
	InsertFirst(T) error
	InsertLast(T) error
	Delete(int32) error
	DeleteFirst() error
	DeleteLast() error
	GetSize() int32
	IsEmpty() bool
}
