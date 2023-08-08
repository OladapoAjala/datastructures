package sequences

type Sequencer[T any] interface {
	GetData(int32) (T, error)
	Contains(T) bool
	Insert(int32, T) error
	InsertFirst(T) error
	InsertLast(T) error
	Delete(int32) error
	DeleteFirst() error
	DeleteLast() error
	Size() int32
	Sort() error
	IsEmpty() bool
}

type Sequence Sequencer[any]
