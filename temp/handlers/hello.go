package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHellow(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops Error", http.StatusBadRequest)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Opps"))
		return
	}
	h.l.Printf("Data %s\n", d)

	fmt.Fprintf(rw, "Hello %s ", d)
}
