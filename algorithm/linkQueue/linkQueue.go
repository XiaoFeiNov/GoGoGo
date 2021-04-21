package linkQueue

import (
	"errors"
)

// 单链表实现栈
type node struct {
	value interface{}
	next  *node
}

type Queue struct {}

var(
	head *node
	tail *node
	size int
)

func New() *Queue {
	head = new(node)
	tail = new(node)
	return new(Queue)
}

// 从链表尾节点入队
func (q *Queue) Enqueue(value interface{}) {
	oldTail := tail
	tail = new(node)
	tail.value = value
	if size > 0 { // 至少有一个
		oldTail.next = tail
		if size == 1 { // 有一个节点时设置头结点
			head = oldTail
		}
	}
	size++
}

// 出队,需要注意队列的size，size为1的时候只有尾结点
func (q *Queue) Dequeue() (value interface{}, err error) {
	if size <= 0 {
		return nil, errors.New("queue Dequeue error: queue is empty")
	} else if size == 1 { // 如果size为1说明只有尾结点，从链表尾节点出队
		value = tail.value
		tail = new(node)
	} else { // size大于1从链表头结点出队
		value = head.value
		head = head.next
	}
	size--
	switch value.(type) { // 这边只处理了int类型
	case int:
		return value.(int), nil
	default:
		return value, errors.New("queue Dequeue value type error")
	}
}

func (q *Queue) Size() interface{} {
	return size
}

func (q *Queue) IsEmpty() bool {
	return size == 0
}
