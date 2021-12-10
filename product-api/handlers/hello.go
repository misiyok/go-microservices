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

// In Go, a name is exported if it begins with a capital letter
// So if we name this as newHello, then wont access this function from another package
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
