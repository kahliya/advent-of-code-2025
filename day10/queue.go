package main

type State struct {
	counter         []int
	excludedButtons []int
	presses         int
}

type Queue struct {
	queue []State
}

func (q *Queue) push(item State) {
	q.queue = append(q.queue, item)
}

func (q *Queue) pop() State {
	val := q.queue[0]
	q.queue = q.queue[1:]

	return val
}
