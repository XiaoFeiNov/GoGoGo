package codeTest

import "sync"

type WaitGroupWrapper struct {
	sync.WaitGroup
}

// add并起一个goroutine执行回调，执行完回调后done
func (w *WaitGroupWrapper) Wrap (callBack func()) {
	w.Add(1)
	go func() {
		callBack()
		w.Done()
	}()
}
