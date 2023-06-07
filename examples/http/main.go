package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pchchv/json_errors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path[1:] == "" {
		s := json_errors.New("Failed")
		err := json_errors.Wrap(s, "You need to enter a name as URL params")
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "Hello,", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
