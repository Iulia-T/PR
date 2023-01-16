package main

// Queue represents a queue that holds a slice
var orderList2 Queue2

type Queue2 struct {
	elements []Order
}

func (q *Queue2) Enqueue2(order Order) {
	q.elements = append(q.elements, order)
}

func (q *Queue2) isEmpty2() bool {
	return len(q.elements) == 0
}

func (ol *Queue2) getSize2() int {
	return len(ol.elements)
}

func (q *Queue2) Dequeue2() *Order {
	if q.isEmpty2() {
		return nil
	}
	order := q.elements[0]
	if q.getSize2() == 1 {
		q.elements = nil
		return &order
	}

	// discard top element
	q.elements = q.elements[1:]
	return &order
}
