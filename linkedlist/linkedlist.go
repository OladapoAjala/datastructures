package linkedlist

type node struct {
	index int32
	value string
	prev  *node
	next  *node
}

type LinkedList struct {
	length int
	head   *node
	tail   *node
}

func (L *LinkedList) Append(value string) {
	if L.head == nil {
		newTail := &node{
			value: value,
			index: 0,
			prev:  nil,
			next:  nil,
		}

		L.head, L.tail = newTail, newTail
		L.length++

		return
	}

	newTail := &node{
		value: value,
		index: L.tail.index + 1,
		prev:  L.tail,
		next:  nil,
	}

	L.tail.next = newTail
	L.tail = newTail
	L.length++

	return
}

func (L *LinkedList) prepend(value string) {
	if L.head == nil {
		head := &node{
			value: value,
			index: 0,
			prev:  nil,
			next:  nil,
		}

		L.head, L.tail = head, head
		L.length++

		return
	}

	newHead := &node{
		value: value,
		index: 0,
		prev:  nil,
		next:  L.head,
	}

	// Shift the indices of the remaining elements
	for it := L.head; it != nil; it = it.next {
		it.index = it.index + 1
	}

	L.head.prev = newHead
	L.head = newHead
	L.length++
}

func (L *LinkedList) insert(index int32, value string) {
	if L.head == nil {
		newTail := &node{
			value: value,
			index: 0,
			prev:  nil,
			next:  nil,
		}

		L.head, L.tail = newTail, newTail
		L.length++

		return
	}

	// Get the current node at the desired index.
	var oldNode *node
	for it := L.head; it != nil; it = it.next {
		if it.index == index {
			oldNode = it
			break
		}
	}

	// Shift the indices of all the elements from the desired index.
	for it := oldNode; it != nil; it = it.next {
		it.index = it.index + 1
	}

	// Create the new node
	newNode := &node{
		value: value,
		index: index,
		prev:  oldNode.prev,
		next:  oldNode,
	}

	// Insert the new node
	oldNode.prev.next = newNode
	oldNode.prev = newNode
	L.length++
}

func (L *LinkedList) remove(index int32) {
	// Loop to the required point
	var oldNode *node
	for it := L.head; it != nil; it = it.next {
		if it.index == index {
			oldNode = it
			break
		}
	}

	// Re-arrange the indices
	for it := oldNode.next; it != nil; it = it.next {
		it.index--
	}
	// set the indexNode.next.prev == indexNode.prev
	oldNode.next.prev = oldNode.prev
	// set the indexNode.prev.next == indexNode.next
	oldNode.prev.next = oldNode.next

	L.length--
}
