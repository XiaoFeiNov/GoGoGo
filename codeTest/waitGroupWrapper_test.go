package codeTest

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

type service struct {

	exitChan chan struct{}
}

func TestWrap(t *testing.T)  {
	w := new(WaitGroupWrapper)
	w.WaitGroup = sync.WaitGroup{}
	s := new(service)
	s.exitChan = make(chan struct{})
	errCh := make(chan error)
	w.Wrap(func() {
		s.Start(errCh)
	})
}

func (s *service) Start(errCh chan error) {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}

func (s *service) Stop() {

}

func handleFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("test", w)
}


