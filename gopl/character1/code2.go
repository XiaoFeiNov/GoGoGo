package main

import (
	"fmt"
	"log"
	"net/http"
)

type hand struct {
}

func (h *hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main()  {
	h := new(hand)
	http.Handle("", h)
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":10086", nil))
}

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("w: ", w)
	fmt.Println("Method: ", req.Method)
	fmt.Println("URL: ", req.URL)
	fmt.Println("Proto: ", req.Proto)
	fmt.Println("Body: ", req.Body)
}