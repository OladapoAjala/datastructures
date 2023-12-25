package queues

type IQueuer[K any] interface {
	Insert(K)
	DeleteMax()
	FindMax()
}
