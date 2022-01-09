package collections

// Queue Implementation

// API
// - Offer
// - Poll
// - Size
// - Empty

const (
	QueueArrayType = 1
	QueueListType = 2
)

type Queue interface {
	Offer(entry interface{}) bool
	Poll() interface{}
	Size() int
	Empty() bool
}

type arrayQueue struct {

}

func (q *arrayQueue) Offer(entry interface{}) bool {
	return true
}

func (q *arrayQueue) Poll() interface{} {
	return true
}

func (q *arrayQueue) Size() int {
	return 0
}

func (q *arrayQueue) Empty() bool {
	return true
}

type linkedListQueue struct {

}

func (q *linkedListQueue) Offer(entry interface{}) bool {
	return true
}

func (q *linkedListQueue) Poll() interface{} {
	return true
}

func (q *linkedListQueue) Size() int {
	return 0
}

func (q *linkedListQueue) Empty() bool {
	return true
}

func NewQueue(t int) Queue {
	switch t {
	case QueueArrayType:
		return &arrayQueue{}
	case QueueListType:
		return &linkedListQueue{}
	}

	// Default
	return &arrayQueue{}
}
