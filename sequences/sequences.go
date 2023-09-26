package sequences

// NOTE: insert_*/delete_* operations change the rank of all items after the modified item.
type Sequencer[T any] interface {
	// New[S types.Sequences[T]]() (S, error)
	GetData(int32) (T, error)
	Contains(T) bool
	Insert(int32, T) error // put value at index
	InsertFirst(T) error
	InsertLast(T) error
	Delete(int32) error
	DeleteFirst() error
	DeleteLast() error
	// Set(int32, T) error // replace value at index
	Size() int32
	IsEmpty() bool
}
