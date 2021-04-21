package linkQueue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	queue := New()
	fmt.Println("queue is empty: ", queue.IsEmpty())
	fmt.Println("queue size: ", queue.Size())
	queue.Enqueue(1)
	fmt.Println("queue size: ", queue.Size())
	value, err := queue.Dequeue()
	fmt.Println("queue size: ", queue.Size())
	if err != nil {
		fmt.Println("queue Dequeue error: ", err)
		return
	}
	fmt.Printf("queue Dequeue value: %d\n", value)
}
