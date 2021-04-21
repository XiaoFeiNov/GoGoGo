package linkStack

import "errors"

// 单链表实现栈
type node struct {
	value interface{}
	next  *node
}

// TODO：head和size放在Stack结构体内比较好，还是外面设置包内变量比较好？（如果不需要访问head，个人认为设置包内变量比较好）
type Stack struct{}

var (
	head *node
	size int
)

func NewStack() *Stack {
	head = new(node)
	return new(Stack)
}

func (s *Stack) Push(value interface{}) {
	oldHead := head
	head = new(node)
	head.value = value
	head.next = oldHead
	size++
}

func (s *Stack) Pop() (interface{}, error) {
	if size <= 0 {
		return -1, errors.New("stack Pop error: stack is empty")
	}
	value := head.value
	head = head.next
	size--
	switch value.(type) { // 这边只处理了int类型
	case int:
		return value.(int), nil
	default:
		return value, errors.New("stack Pop value type error")
	}
}

func (s *Stack) Size() int {
	return size
}

func (s *Stack) IsEmpty() bool {
	return size == 0
}
