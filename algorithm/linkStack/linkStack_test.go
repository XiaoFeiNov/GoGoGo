package linkStack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	fmt.Println("stack is empty: ", stack.IsEmpty())
	fmt.Println("stack size: ", stack.Size())
	stack.Push(1)
	fmt.Println("stack size: ", stack.Size())
	value, err := stack.Pop()
	fmt.Println("stack size: ", stack.Size())
	if err != nil {
		fmt.Println("stack pop error: ", err)
	}
	fmt.Println("stack pop value: ", value)
}
